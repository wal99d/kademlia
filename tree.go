package kademlia

type hashTree struct {
	zero, one *hashTree
	peer      *Peer
}

func (h *hashTree) insert(p *Peer) error {
	return h.add(p, 0)
}

func (h *hashTree) add(p *Peer, i int) error {

	switch {
	case i >= len(p.id)*8:
		h.peer = p
		return nil
	case h.peer != nil:
		if h.peer.id == p.id {
			h.peer = p
			return nil
		}
		//start compressing hashTree
		old := h.peer
		h.peer = nil
		h.compressHashTree(p, old, i)
		return nil
	case h.peer == nil:
		postion := byte(i % 8)
		pivotValue := byte(p.id[i/8])
		if (pivotValue<<postion)&128 == 0 {
			if h.zero == nil {
				h.zero = &hashTree{peer: p}
				return nil
			}
			h.zero.add(p, i+1)
		} else {
			if h.one == nil {
				h.one = &hashTree{peer: p}
				return nil
			}
			h.one.add(p, i+1)
		}
	}
	return nil
}

func (h *hashTree) compressHashTree(p *Peer, old *Peer, i int) error {
	postion := byte(i % 8)
	pivotNewValue := byte(p.id[i/8])
	msbNew := (pivotNewValue << postion) & 128
	pivotOldValue := byte(old.id[i/8])
	msbOld := (pivotOldValue << postion) & 128

	switch {
	case msbNew == msbOld:
		if msbNew == 0 {
			h.zero = &hashTree{}
			h.zero.compressHashTree(p, old, i+1)
		} else {
			h.one = &hashTree{}
			h.one.compressHashTree(p, old, i+1)
		}
	case msbNew != msbOld:
		h.add(p, i)
		h.add(old, i)
	}
	return nil
}

func (h *hashTree) traverse(id ID, i int, peers []*Peer, filter bool) []*Peer {
	//TODO: Implement it
}
