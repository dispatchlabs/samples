package main

import (
	"container/list"
	"encoding/hex"
	"math/rand"
)

// Data Type
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
const IdLength = 20

type NodeID [IdLength]byte

// NodeID Convert-ors / Sort-ers / Diff-ers
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}
	return
}
func NodeIDFromString(data string) (ret NodeID) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}
	return
}

func (node NodeID) ToString() string {
	return hex.EncodeToString(node[0:IdLength])
}
func (node NodeID) String() string {
	// Implement the `fmt.Stringer` interface
	// Meaning this will be used to convert a `NodeID` for display in `Printf` or similar
	return node.ToString()
}

func (node NodeID) Equals(other NodeID) bool {

	// `Equals / Less` define a well-ordering for NodeIDs.
	// NodeIDs are big-endian - most significant byte being the low-order one.
	// This will allow us to sort NodeIDs, which will be important later.

	for i := 0; i < IdLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other.(NodeID)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

func (node NodeID) Xor(other NodeID) (ret NodeID) {
	// All DHTs rely on having a 'distance metric'
	// a way of comparing two NodeIDs or hashes to determine how far apart they are.
	//
	// Kademlia uses the 'XOR metric'.
	// The distance between two NodeIDs is the XOR of those two NodeIDs interpreted as a number.

	for i := 0; i < IdLength; i++ {
		ret[i] = node[i] ^ other[i]
	}

	return
}

// Routing Table
//		To enable efficient traversal through a DHT, a routing table needs to contain a
//		selection of nodes both close to and far away from our own node.
//
//		Kademlia accomplishes this by breaking up the routing table into 'buckets'.
//		Each 'bucket' corresponds to a range of distances between the nodes in that
//		bucket and current node.
//
//		Ex:	Bucket 0 contains nodes that differ in the first bit.
//			Bucket 1 contains nodes that differ in the second bit.
//			And so forth.
//
//		This exponential sizing of buckets has a couple of implications.
//		First:
//			- half the nodes in the DHT should be expected to end up in bucket 0
//			- half of the remainder in bucket 1
//			- and so forth
//			This means that we should have a complete set of all the nodes nearest to us
//			gradually getting sparser over increasing distance.
//			This is necessary in order to ensure we can always find data if it exists in the DHT.
//		Second:
//			The number of the bucket a given node should be placed in is determined by
//			the number of leading 0 bits in the XOR of our node ID with the target node ID,
//			which makes for easy implementation. `PrefixLen` facilitates this.

func (node NodeID) PrefixLen() (ret int) {
	for i := 0; i < IdLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return IdLength*8 - 1
}

const BucketSize = 20

type Contact struct {
	id NodeID
}

type RoutingTable struct {
	node    NodeID
	buckets [IdLength * 8]*list.List
}

func NewRoutingTable(node NodeID) (ret RoutingTable) {
	for i := 0; i < IdLength*8; i++ {
		ret.buckets[i] = list.New()
	}
	ret.node = node
	return
}

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

// Entry Point
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
func main() {

}
