package main

import (
	"flag"
	"log"
	"sync"
)

var (
	NumOfNodes      int
	PeersLimitRatio = 0.1
)

func main() {
	var wg sync.WaitGroup

	flag.IntVar(&NumOfNodes, "numofnodes", 0, "number of nodes")
	flag.Parse()

	if NumOfNodes <= 0 {
		log.Fatal("number of nodes assigned should be more than 0")
	}

	app := NewApplication(&wg, CreateNodes(&wg, NumOfNodes))

	app.ManagePeerDistribution()
	app.StartNodes()
}
