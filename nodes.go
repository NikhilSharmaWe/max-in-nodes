package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	ID         int
	NodeNumber uint64
	MaxNumber  uint64
	Peers      []int
	MsgCh      chan uint64

	sync.Mutex
	*sync.WaitGroup
}

func NewNode(id int, wg *sync.WaitGroup) *Node {
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := seedRand.Uint64()

	return &Node{
		ID:         id,
		NodeNumber: n,
		MaxNumber:  n,
		MsgCh:      make(chan uint64, NumOfNodes-1),
		WaitGroup:  wg,
	}
}

func (node *Node) AddPeers(nodes []*Node) {
	numOfPeers := int(PeersLimitRatio * float64(NumOfNodes))
	peersIDs := make([]int, 0, numOfPeers)

	for i := 0; i < numOfPeers; i++ {
		peerID := rand.Intn(NumOfNodes)
		if peerID != node.ID && !alreadyPresent(peersIDs, peerID) {
			peersIDs = append(peersIDs, peerID)
		}
	}

	node.Peers = peersIDs
}

func (node *Node) Start(nodes []*Node) {
	defer node.Done()

	for _, peerID := range node.Peers {
		nodes[peerID].MsgCh <- node.NodeNumber
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case number := <-node.MsgCh:
			node.Lock()
			if number > node.MaxNumber {
				node.MaxNumber = number

				for _, peerID := range node.Peers {
					nodes[peerID].MsgCh <- number
				}
				ticker.Reset(5 * time.Second)
			}
			node.Unlock()

		case <-ticker.C:
			fmt.Printf("Node: %d | MaxNumber: %d\n", node.ID, node.MaxNumber)
			return
		}
	}
}

func CreateNodes(wg *sync.WaitGroup, numOfNodes int) []*Node {
	nodes := make([]*Node, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		nodes[i] = NewNode(i, wg)
	}

	return nodes
}
