package main

import "sync"

type Application struct {
	WG    *sync.WaitGroup
	Nodes []*Node
}

func NewApplication(wg *sync.WaitGroup, nodes []*Node) *Application {
	return &Application{
		WG:    wg,
		Nodes: nodes,
	}
}

func (app *Application) ManagePeerDistribution() {
	for _, node := range app.Nodes {
		node.AddPeers(app.Nodes)
	}
}

func (app *Application) StartNodes() {
	app.WG.Add(NumOfNodes)

	for _, node := range app.Nodes {
		go node.Start(app.Nodes)
	}

	app.WG.Wait()
}
