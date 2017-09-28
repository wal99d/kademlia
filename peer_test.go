package kademlia_test

import(
	"testing"
)

func TestPeerComparison(t *testing.T){
	c1:=NewPeer(NewRandomID(), "127.0.0.1:3333")
	c2:=NewPeer(NewRandomID(), "127.0.0.1:3333")
	Assert(t, c1.Less(0,1), "peers aren't less")
}