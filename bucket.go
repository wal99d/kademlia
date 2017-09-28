package kademlia

import "container/list"

const BucketSize = 20

type Bucket struct {
	peerList [IDLen * 8]*list.List
}

func NewBucket() (b Bucket) {
	for i := 0; i < IDLen*8; i++ {
		b.peerList[i] = list.New()
	}
	return
}
