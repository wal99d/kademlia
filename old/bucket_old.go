package kademlia

import (
	"errors"
	"fmt"
	"sort"
)

const BucketSize = 20

type RoutingTable struct{
	buckets [IDLen * 8]*PeersTree
}



func NewRoutingTable() *RoutingTable {
	return &RoutingTable{
		buckets: make([]*PeersTree, IDLen*8),
	}
}

func (t PeersTree) binarySearch(address string) (int, *Peer, error) {
	pos := sort.Search(len(t.peers), func(i int) bool {
		return t.peers[i].address >= address
	})

	if pos < len(t.peers) && t.peers[pos].address == address {
		return pos, t.peers[pos], nil
	}

	return pos, nil, errors.New("peer with this id not found")
}

// Insert adds a new peer at the correct place in the tree.
func (pt *PeersTree) InsertPeer(peer *Peer) error {
	pos, _, err := t.binarySearch(peer.address)
	bucketIndex := peer.id.Xor(pt.peers[pos].id).Prefixlen()

	if err == nil {
		return errors.New("peer already present")
	}

	// https://code.google.com/p/go-wiki/wiki/SliceTricks
	pt.peers = append(pt.peers, &Peer{})
	copy(pt.peers[pos+1:], pt.peers[pos:])
	pt.peers[pos] = peer
	insertToBucket(pt, )


	return nil
}

func insertToBucket(pt *PeersTree) error{

}

func (t RoutingTable) String() string {
	return fmt.Sprintf("RoutingTableTree<%d peers>", len(t.peers))
}

// func (cr *ContactRecord) Less(other interface{}) bool {
// 	return cr.sortkey.Less(other.(*ContactRecord).sortkey)
// }

// // assumes bucket is already locked, slice has proper capacity
// func AddBucketContentsToSlice(bucket *list.List, requester NodeID, s []Contact) {
// 	var maxToAdd, count int = cap(s) - len(s), 0
// 	for con := bucket.Front(); con != nil && count < maxToAdd; con = con.Next() {
// 		//if false == con.Value.(Contact).id.Equals(requester) {
// 		s = append(s, con.Value.(Contact))
// 		count += 1
// 		//}
// 	}
// }

// func AddBucketToSlice(rt *RoutingTable, requester NodeID, bucketNum int, s []Contact) {
// 	//k.contactsMutex[bucketNum].Lock()
// 	AddBucketContentsToSlice(rt.buckets[bucketNum], requester, s)
// 	//k.contactsMutex[bucketNum].Unlock()
// }

// func (rt *RoutingTable) FindClosest(tagetNode NodeID, requester NodeID, count int) []Contact {
// 	bucketIndex := tagetNode.Xor(rt.node.id).Prefixlen()
// 	contacts := make([]Contact, 0, count)

// 	if bucketIndex < BucketSize {
// 		AddBucketToSlice(rt, requester, bucketIndex, contacts)
// 	}

// 	for i := 1; (i <= bucketIndex || i+bucketIndex < BucketSize) && len(contacts) < count; i++ {
// 		if i <= bucketIndex {
// 			AddBucketToSlice(rt, requester, bucketIndex-i, contacts)
// 		}

// 		if i+bucketIndex < BucketSize {
// 			AddBucketToSlice(rt, requester, i+bucketIndex, contacts)
// 		}
// 	}

// 	// this shouldn't happen
// 	if len(contacts) > count {
// 		contacts = contacts[0:count]
// 	}

// 	return contacts
// }
