package kademlia

import (
	"fmt"
)

type Peer struct {
	id      ID
	address string
}

func NewPeer(i ID, a string) *Peer {
	return &Peer{
		id:      i,
		address: a,
	}
}

// Peers is a slice of peers that can be sorted.
type Peers []*Peer

func (p Peers) Len() int           { return len(p) }
func (p Peers) Less(i, j int) bool { return p[i].id.Less(p[j].id) }
func (p Peers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *Peer) String() string {
	return fmt.Sprintf("Peer ID: %s, Address=%s\n", p.id, p.address)
}
