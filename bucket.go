package kademlia

import (
	"container/list"
	"fmt"
)

const BucketSize = 20

type Bucket struct {
	peerList [IDLen * 8]*list.List
}

//NewBucket will create brand new Bucket of type *list.List to be used for any
//for each client once started
func NewBucket() (b Bucket) {
	for i := 0; i < IDLen*8; i++ {
		b.peerList[i] = list.New()
	}
	return
}

func (b *Bucket) updateMostRecentlySeen(elem *list.Element, i int) error {
	b.peerList[i].MoveToBack(elem)
	return nil
}

//InsertInBucket will check if Peer is added to list's back/Tail indexed by bucketIndex
func (b *Bucket) InsertInBucket(p *Peer, bucketIndex int) error {
	v := b.peerList[bucketIndex]
	for e := v.Front(); e != nil; e = e.Next() {
		if e.Value.(*Peer).id.Equals(p.id) {
			//Already existed peer
			err := b.updateMostRecentlySeen(e, bucketIndex)
			if err != nil {
				return err
			}
			return nil
		} else {
			//Add peer to list's back/tail
			if v.Len() < BucketSize {
				v.PushBack(p)
				return nil
			} else {
				//TODO: Handle if there is no space in Bucket list
				return nil
			}
		}
	}

	return nil
}

func (b Bucket) String() string {
	var str string
	for k, v := range b.peerList {
		for e := v.Front(); e != nil; e = e.Next() {
			str = str + fmt.Sprintf("At Bucket[%d] =  %s\n", k, e.Value.(*Peer).id)
		}
	}
	return str
}
