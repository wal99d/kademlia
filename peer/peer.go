package main

import (
	"fmt"
	"kademlia"
)

func main() {
	p1 := kademlia.NewPeer(kademlia.NewRandomID(), "127.0.0.1:33331")
	p2 := kademlia.NewPeer(kademlia.NewRandomID(), "127.0.0.1:33332")
	p3 := kademlia.NewPeer(kademlia.NewRandomID(), "127.0.0.1:33333")
	p4 := kademlia.NewPeer(kademlia.NewRandomID(), "127.0.0.1:33334")
	fmt.Printf("p1:%v", p1)
	fmt.Printf("p2:%v", p2)
	fmt.Printf("p3:%v", p3)
	fmt.Printf("p4:%v", p4)

	t := kademlia.NewTree()
	t.Insert(p1)
	t.Insert(p2)
	t.Insert(p3)
	t.Insert(p4)
	t.ShowTree()

	fmt.Printf("Buckets: %s\n", t)
	/*
		ctx := context.Background()
		c := kademlia.NewCotact(kademlia.NewRandomNodeID(), "home")
		//log.Printf("Created peer with id:%s and address:%s\n", c.GetId(), c.GetAddress())
		k := kademlia.NewKademlia(c, "home")
		go k.Run(c.GetId())
		client, conn := k.Connect(c.GetId())
		defer conn.Close()
		req := &pb.Request{}
		req.Contact = &pb.Contact{
			Id:      c.GetId(),
			Address: c.GetAddress(),
		}
		req.NetworkId = "home"
		resp, err := client.Ping(ctx, req)
		if err != nil {
			log.Printf("%s\n", err)
		}
		log.Printf("**************")
		if resp == nil {
			log.Printf("There is not response from remote peer\n")
		} else {
			log.Printf("Got response from peer with peer id:%x and address:%s\n", resp.Contact.Id, resp.Contact.Address)
		}
		tar := &pb.Target{}
		tar.Id = c.GetId()
		nodes, err := client.FindNode(ctx, tar)
		if err != nil {
			log.Printf("%s\n", err)
		}
		log.Printf("&&&&&&&&&&&&")
		for i := 0; i < len(nodes.Contact); i++ {
			log.Printf("peer with id:%x and address:%s is Online\n", &nodes.Contact[i].Id, &nodes.Contact[i].Address)
		}
		log.Printf("**************")
		for {
		}
	*/
}
