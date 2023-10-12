package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	node "github.com/Grumlebob/AuctionSystemReplication/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type peer struct {
	node.UnimplementedNodeServer
	id             int32
	privateValue   int32
	receivedShares []int32
	summedValues   int32
	clients        map[int32]node.NodeClient
	ctx            context.Context
}

var uniqueIdentifier = int32(0)
var randomValueCap = int32(1000)

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	log.SetFlags(log.Ltime)

	ctx, cancel := context.WithCancel(context.Background())

	randomPrivateValue := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(randomValueCap)

	defer cancel()
	p := &peer{
		id:             ownPort,
		privateValue:   randomPrivateValue,
		receivedShares: make([]int32, 3),
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
	//3 Peers connected on port 5000 (Alice), 5001(Bob), 5002(Charlie), 5003(Hospital)
	for i := 0; i < 4; i++ {
		port := int32(5000) + int32(i)
		if port == ownPort {
			continue
		}
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := node.NewNodeClient(conn)
		p.clients[port] = c
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())

		//FORCE CRASH COMMAND
		if strings.Contains(text, "crash") {
			log.Printf("Crashing node id %v ", p.id)
			os.Exit(1)
		}
	}
}

func (p *peer) HandlePeerRequest(ctx context.Context, req *node.Request) (*node.Reply, error) {
	//p er den client der svarer på requesten.
	//req kommer fra anden peer.
	//Reply er det svar peer får.
	if p.state == WANTED {
		if req.State == RELEASED {
			p.responseNeeded--
		}
		if req.State == WANTED {
			if req.LamportTime > p.lamportTime {
				p.responseNeeded--
			} else if req.LamportTime == p.lamportTime && req.Id < p.id {
				p.responseNeeded--
			}
		}
	}
	p.lamportTime = int32(math.Max(float64(p.lamportTime), float64(req.LamportTime))) + 1
	reply := &node.Reply{Id: p.id, State: p.state, LamportTime: p.lamportTime}
	return reply, nil
}

func (p *peer) sendMessageToAllPeers() {
	request := &node.Request{Id: p.id, State: p.state, LamportTime: p.lamportTime}
	for _, client := range p.clients {
		reply, err := client.HandlePeerRequest(p.ctx, request)
		if err != nil {
			log.Println("something went wrong")
		}

	}
}

func (p *peer) sendHeartbeatPingToPeerId(peerId int32) {
	_, found := p.clients[peerId]
	if !found {
		log.Printf("Client node %v is dead", peerId)
		//Remove dead node from list of clients
		delete(p.clients, peerId)
		//If leader crashed, elect new leader
		if peerId == p.idOfPrimaryReplicationManager {
		}
		return
	}
	_, err := p.clients[peerId].PingLeader(p.ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("Client node %v is dead", peerId)
		//Remove dead node from list of clients
		delete(p.clients, peerId)
		//If leader crashed, elect new leader
		if peerId == p.idOfPrimaryReplicationManager {
			p.LookAtMeLookAtMeIAmTheCaptainNow()
		}
	}
}
