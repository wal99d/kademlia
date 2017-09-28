package kademlia

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const IDLen = 20

type ID [IDLen]byte

func HandleErr(caller string, e error) {
	if e != nil {
		fmt.Printf("Error from: %s==> %v", caller, e)
	}
}

func NewID(s string) (n ID) {
	decoded, err := hex.DecodeString(s)
	HandleErr("NewID()", err)
	copy(n[:], decoded)
	return
}

func NewRandomID() (n ID) {
	b := make([]byte, IDLen)
	rand.Read(b)
	copy(n[:], b)

	return n
}

func (node ID) String() string {
	return hex.EncodeToString(node[:])
}

func (node ID) Equals(other ID) bool {
	return node == other
}

func (node ID) Less(other interface{}) bool {
	for i, v := range other.(ID) {
		if node[i] != v {
			return node[i] < v
		}
	}
	return false
}

func (node ID) Xor(other interface{}) (n ID) {
	for i, v := range other.(ID) {
		n[i] = node[i] ^ v
	}
	return
}

func (node ID) Prefixlen() int {
	for i, _ := range node {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(j))&0x01 != 0 {
				return i*8 + j
			}
		}
	}
	return IDLen * 8
}

// func main() {
// 	n1 := NewID("12345678900987654321")
// 	n2 := NewRandomID()
// 	n3 := n2
// 	n4 := NewRandomID()
// 	fmt.Println("n1=", n1.String())
// 	fmt.Println("n2=", n2.String())
// 	fmt.Println("n2.Equals(n3) = ", n2.Equals(n3))
// 	fmt.Println("n2.Less(n4) = ", n2.Less(n4))
// 	fmt.Printf("Metric distance of n2 to n4:%b\n", n2.Xor(n4))
// 	fmt.Printf("Bucket: %d\n", n2.Xor(n4).Prefixlen())
// }
