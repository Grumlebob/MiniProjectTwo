package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	node "github.com/Grumlebob/MiniProjectTwo/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type peer struct {
	node.UnimplementedNodeServer
	id             int32
	name           string
	privateValue   int32
	receivedShares []int32
	summedValues   int32
	clients        map[int32]node.NodeClient
	ctx            context.Context
}

var randomValueCap = int32(1000)

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	log.SetFlags(log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())

	randomPrivateValue := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(randomValueCap)

	//if ownport is port 5000 (Alice), 5001(Bob), 5002(Charlie), 5003(Hospital)
	nameOfPeer := ""
	if ownPort == 5000 {
		nameOfPeer = "Alice"
	} else if ownPort == 5001 {
		nameOfPeer = "Bob"
	} else if ownPort == 5002 {
		nameOfPeer = "Charlie"
	} else if ownPort == 5003 {
		nameOfPeer = "Hospital"
	} else {
		nameOfPeer = "Unknown"
	}

	defer cancel()
	p := &peer{
		id:             ownPort,
		name:           nameOfPeer,
		privateValue:   randomPrivateValue,
		receivedShares: make([]int32, 0),
		summedValues:   0,
		clients:        make(map[int32]node.NodeClient),
		ctx:            ctx,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	node.RegisterNodeServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()
	//4 Peers connected on port 5000 (Alice), 5001(Bob), 5002(Charlie), 5003(Hospital)
	for i := 0; i < 4; i++ {
		port := int32(5000) + int32(i)
		if port == ownPort {
			continue
		}
		//tlsConfig := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false}) //accept all certificates

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := node.NewNodeClient(conn)
		p.clients[port] = c
	}

	if ownPort != 5003 {
		log.Printf("Peer %v port(%v) started, with private value %v", p.name, p.id, p.privateValue)
		p.sendSharesToAllPeers()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())

		//FORCE print
		if strings.Contains(text, "print") {
			log.Printf("Peer %v has value %v", p.name, p.privateValue)
			log.Printf("has received %v shares", len(p.receivedShares))
			for i := range p.receivedShares {
				log.Printf("share %v is %v", i, p.receivedShares[i])
			}
		}
	}
}

func (p *peer) HandlePeerRequest(ctx context.Context, req *node.Request) (*emptypb.Empty, error) {
	//p er den client der svarer på requesten.
	//req kommer fra anden peer.
	//Reply er det svar den anden peer får.

	log.Printf("Peer %v received data %v", p.name, req.Share)
	//add share to receivedShares
	p.receivedShares = append(p.receivedShares, req.Share)

	//print how many shares were received so far
	log.Printf("Peer %v received a total of %v shares", p.name, len(p.receivedShares))

	//if all shares are received, sum them up and send to hospital
	if len(p.receivedShares) == 3 {
		log.Printf("Peer %v received all shares", p.name)
		for _, share := range p.receivedShares {
			log.Printf("Peer %v has share %v", p.name, share)
			p.summedValues += share
		}
		if p.id == 5003 { //If we are hospital, instead of sending shares again, just sum up the values, as that is final output
			log.Printf("Hospital has final output aggregated sum of: %v", p.summedValues)

		} else {
			log.Printf("Peer %v is sending to hospital", p.name)
			p.sendMessageToHospital()
		}
	}

	return &emptypb.Empty{}, nil
}

func (p *peer) sendSharesToAllPeers() {

	//Share One and Share Two should be between -RandomValueCap and RandomValueCap
	shareForAlice := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(randomValueCap)
	shareForBob := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(randomValueCap)
	//Share three should have difference between share one and share two and P's private value
	shareForCharlie := p.privateValue - (shareForAlice + shareForBob)

	if shareForAlice+shareForBob+shareForCharlie != p.privateValue {
		log.Printf("Shares do not add up to private value")
	}

	log.Printf("Peer %v is sending shares to all peers", p.name)

	for portCounter := 5000; portCounter <= 5003; portCounter++ {

		currentClient := p.clients[int32(portCounter)]
		if portCounter == 5003 { //Hospital does not make shares and send them
			continue
		} else if int32(portCounter) == p.id { //If it is the peer itself, just add the their own share to receivedShares, without communicating over network.

			if portCounter == 5000 { //Alice
				log.Printf("Alice is sending share %v to herself", shareForAlice)
				p.receivedShares = append(p.receivedShares, shareForAlice)
			}
			if portCounter == 5001 { //Bob
				log.Printf("Bob is sending share %v to himself", shareForBob)
				p.receivedShares = append(p.receivedShares, shareForBob)
			}
			if portCounter == 5002 { //Charlie
				log.Printf("Charlie is sending share %v to himself", shareForCharlie)
				p.receivedShares = append(p.receivedShares, shareForCharlie)
			}

			//if all shares are received, sum them up and send to hospital
			if len(p.receivedShares) == 3 {
				log.Printf("Peer %v received all shares", p.name)
				for _, share := range p.receivedShares {
					log.Printf("Peer %v has share %v", p.name, share)
					p.summedValues += share
				}
				//peer is sending to hospital
				p.sendMessageToHospital()
			}

		} else if portCounter == 5000 { //Alice
			aliceRequest := &node.Request{Share: shareForAlice}
			log.Printf("Port %v is sending to %v, with share %v", p.id, portCounter, shareForAlice)
			_, err := currentClient.HandlePeerRequest(p.ctx, aliceRequest)
			if err != nil {
				log.Println("something went wrong")
			}
		} else if portCounter == 5001 { //Bob
			bobRequest := &node.Request{Share: shareForBob}
			log.Printf("Port %v is sending to %v, with share %v", p.id, portCounter, shareForBob)
			_, err := currentClient.HandlePeerRequest(p.ctx, bobRequest)
			if err != nil {
				log.Println("something went wrong")
			}
		} else if portCounter == 5002 { //Charlie
			charlieRequest := &node.Request{Share: shareForCharlie}
			log.Printf("Port %v is sending to %v, with share %v", p.id, portCounter, shareForCharlie)
			_, err := currentClient.HandlePeerRequest(p.ctx, charlieRequest)
			if err != nil {
				log.Println("something went wrong")
			}
		}
	}

}

func (p *peer) sendMessageToHospital() {

	request := &node.Request{Share: p.summedValues}
	log.Printf("Peer %v is sending public aggregated shares: %v to hospital", p.name, p.summedValues)

	_, err := p.clients[5003].HandlePeerRequest(p.ctx, request)
	if err != nil {
		log.Println("something went wrong")
	}

}
