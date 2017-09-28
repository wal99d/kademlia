package kademlia

import (
	"errors"
	"sort"
)

type PeersTree struct {
	peers []*Peer
}

func (pt PeersTree) binarySearch(address string) (int, *Peer, error) {
	pos := sort.Search(len(pt.peers), func(i int) bool {
		return pt.peers[i].address >= address
	})

	if pos < len(pt.peers) && pt.peers[pos].address == address {
		return pos, pt.peers[pos], nil
	}

	return pos, nil, errors.New("peer with this id not found")
}

func (pt *PeersTree) FindPeer(peer *Peer) (int, error) {
	pos := sort.Search(len(peer), func(i int) bool {
		return pt.peers[i] == peer
	})

	return pos, nil
}

// Insert adds a new peer at the correct place in the tree.
func (pt *PeersTree) InsertPeer(peer *Peer) error {
	pos, _, err := t.binarySearch(peer.address)

	if err == nil {
		return errors.New("peer already present")
	}

	// https://code.google.com/p/go-wiki/wiki/SliceTricks
	pt.peers = append(pt.peers, &Peer{})
	copy(pt.peers[pos+1:], pt.peers[pos:])
	pt.peers[pos] = peer

	return nil
}
