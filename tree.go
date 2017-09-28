package kademlia

import (
	"errors"
	"fmt"
)

type Node struct {
	bucket Bucket
	prefix int
	peer   *Peer
	left   *Node
	right  *Node
}

func (n *Node) Insert(p *Peer) error {
	if n == nil {
		return errors.New("Cannot insert bucket into empty Tree!!")
	}
	//fmt.Printf("n id is:%v\n", n.peer.id)
	switch {
	// case p.id.Xor(n.peer.id).Prefixlen() == n.prefix:
	// 	return nil
	case p.id.Xor(n.peer.id).Prefixlen() >= n.prefix:
		if n.right == nil {
			n.right = &Node{
				bucket: NewBucket(),
				peer:   p,
				prefix: p.id.Xor(n.peer.id).Prefixlen(),
			}
			return nil
		}
		return n.right.Insert(p)
	case p.id.Xor(n.peer.id).Prefixlen() < n.prefix:
		if n.left == nil {
			n.left = &Node{
				bucket: NewBucket(),
				peer:   p,
				prefix: p.id.Xor(n.peer.id).Prefixlen(),
			}
			return nil
		}
		return n.left.Insert(p)
	}
	return nil
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(p *Peer) error {
	if t.Root == nil {
		t.Root = &Node{
			bucket: NewBucket(),
			peer:   p,
		}
		return nil
	}
	return t.Root.Insert(p)
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.left, f)
	f(n)
	t.Traverse(n.right, f)
}

// func (t Tree) String() string {
// 	return fmt.Sprintf("Tree root:%v\n", t.Root.peer.id)
// }

func (t Tree) ShowTree() {
	t.Traverse(t.Root, func(n *Node) { fmt.Print(n.peer.id, ": ", n.peer.address, ":", n.prefix, " | \n") })
}
