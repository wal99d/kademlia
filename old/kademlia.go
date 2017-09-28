package kademlia

import (
	"errors"
	"flag"
	"fmt"
	pb "kademlia/pb"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Kademlia struct {
	routes    *RoutingTable
	networkId string
}

func NewKademlia(c *Contact, nid string) (k *Kademlia) {
	k = new(Kademlia)
	k.networkId = nid
	k.routes = NewRoutingTable(c)
	return
}

func (k *Kademlia) Run(peerId NodeID) {
	port := flag.Int("p", 22221, "Port to listen to..")
	flag.Parse()
	log.Printf("Peer id:%s is Listening to port: %d", peerId, *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterKademliaServer(s, *k)
	s.Serve(lis)
}

func (k *Kademlia) Connect(peerId NodeID) (pb.KademliaClient, *grpc.ClientConn) {
	address := flag.String("a", "127.0.0.1:22221", "Address of Peer..")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to grpc server %v", err)
	}

	log.Printf("===> Connected to peer id:%s through address:%s\n", peerId, *address)
	client := pb.NewKademliaClient(conn)

	return client, conn
}

func (k Kademlia) Ping(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	if in.NetworkId != k.networkId {
		return nil, errors.New("Cannot connect to differenct network ID!!")
	}
	c := NewCotact(in.Contact.Id, in.Contact.Address)
	k.routes.Update(c)
	resp := &pb.Contact{
		Id:      c.id,
		Address: c.address,
	}
	return &pb.Response{resp}, nil
}

func (k Kademlia) FindNode(ctx context.Context, t *pb.Target) (*pb.Nodes, error) {
	contacts := k.routes.FindClosest(t.Id, k.routes.node.id, BucketSize)
	resp := &pb.Nodes{}
	//c := make([]*pb.Contact, len(contacts))
	//copy(resp.Contact, contacts)
	fmt.Println("contacts", contacts)
	return resp, nil
}
