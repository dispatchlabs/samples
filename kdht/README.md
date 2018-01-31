# DHT
- a peer-to-peer networking primitive that permits storage/lookup of key-value pairs
- it's a hash table that's distributed across many nodes
- newer DHTs permit many other operations besides data storage

# Kademlia DHT
- competition
	- http://en.wikipedia.org/wiki/Chord_(peer-to-peer)
	- http://en.wikipedia.org/wiki/Pastry_(DHT)
- are no explicit routing update messages
- the internal state it maintains is fairly straightforward and easy to understand
- lookups are accomplished in an obvious and very efficient manner
- Kademlia sacrifices a few of the features of competitors - it's not as practical to implement other primitives such as pubsub over it.
- Implementation Concepts
	- central concept is the Node ID, a 160 bit opaque identifier, used to identify 
		- nodes on the network
		- as well as content hashes


# Refs
- http://xlattice.sourceforge.net/components/protocol/kademlia/specs.html
- https://www.cs.rice.edu/Conferences/IPTPS02/109.pdf
- http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-1
- http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-2