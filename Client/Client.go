package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"sort"
)

type uints []uint32

func (x uints) Len() int { return len(x) }
func (x uints) Less(i, j int) bool { return x[i] < x[j] }
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type ConsistentHashingMgr struct {
	circle           map[uint32]string
	members          map[string]bool
	sortedHashes     uints
}

func hashKey(elt string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(elt))
	return h.Sum32() % 1000
}

// need c.Lock() before calling
func (c *ConsistentHashingMgr) init(elt []string) {
	c.circle = make(map[uint32]string)
	c.members = make(map[string]bool)
	for i := 0; i < len(elt); i++ {
		hash := hashKey(elt[i])
		c.sortedHashes = append(c.sortedHashes, hash)
		c.circle[hash] = elt[i]
		c.members[elt[i]] = true
		fmt.Println("hash added is : ", hash)
	}
	sort.Sort(c.sortedHashes)
}

// Get returns an element close to where name hashes to in the circle.
func (c *ConsistentHashingMgr) GetShardForKey(name string) (string, string) {
	if len(c.circle) == 0 {
		return "", "Empty list of shards"
	}
	key := hashKey(name)
	i := c.searchShardForKey(key)
	return c.circle[c.sortedHashes[i]], ""
}

func (c *ConsistentHashingMgr) searchShardForKey(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func main() {
	router := NewRouter()
	cons.init([]string{"localhost:3000", "localhost:3001", "localhost:3002"})
	http.ListenAndServe("localhost:8080", router)
}
