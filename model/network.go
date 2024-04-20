package model

import (
	"CienciasII-RED/gui"
	"CienciasII-RED/utils"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup

type Network struct {
	Nodes            []*Node
	MaxPackageLength int
}

func (g *Network) ConnectRouters(id1 string, id2 string, weight int) {
	if weight == 0 {
		weight = rand.Intn(15) + 1
	}
	var node1 *Node
	var node2 *Node
	for _, i := range g.Nodes {
		if id1 == i.Id {
			node1 = i
		}
		if id2 == i.Id {
			node2 = i
		}
	}
	if node1 != nil && node2 != nil {
		vertex := NewVertex(node1, node2, weight)
		node1.AddVertex(vertex)
		node2.AddVertex(vertex)
	} else {
		fmt.Println("Nodo no encontrado")
	}
}

func (g *Network) AddRouter(id string) *Node {
	var newNode Node
	node := g.GetNode(id)
	if node != nil {
		fmt.Println("El nodo ya existe")
	} else {
		newNode = Node{Id: id}
		g.Nodes = append(g.Nodes, &newNode)
	}
	return &newNode
}

func (g *Network) CalculateShortestPath(idSender string, idReceiver string) utils.PathData {
	pathsTable := utils.PathsTable{}
	for _, node := range g.Nodes {
		pathsTable[node.Id] = &utils.PathData{Locked: false, Path: "", Shortest: -1}
	}
	pathsTable[idSender].Locked = true
	pathsTable[idSender].Shortest = 0
	iterable := g.GetNode(idSender)
	for pathsTable[idReceiver].Locked == false {
		iterableData := pathsTable[iterable.Id]
		for _, vertex := range iterable.Vertices {
			var objectiveNode *Node
			if vertex.Node1.Id == iterable.Id {
				objectiveNode = vertex.Node2
			} else {
				objectiveNode = vertex.Node1
			}
			objectiveData := pathsTable[objectiveNode.Id]
			if objectiveData.Locked == false {
				if objectiveData.Shortest > iterableData.Shortest+vertex.Weight || objectiveData.Shortest == -1 {
					objectiveData.Shortest = iterableData.Shortest + vertex.Weight
					objectiveData.Path = fmt.Sprintf("%s->%s", iterableData.Path, iterable.Id)
				}
			}
		}
		minShortest := -1
		var minId string
		for id, data := range pathsTable {
			if data.Locked == false && data.Shortest != -1 {
				if minShortest == -1 || data.Shortest < minShortest {
					minShortest = data.Shortest
					minId = id
				}
			}
		}
		pathsTable[minId].Locked = true
		iterable = g.GetNode(minId)
	}
	receiverData := pathsTable[idReceiver]
	receiverData.Path = fmt.Sprintf("%s->%s", receiverData.Path, idReceiver)
	return *receiverData
}

func (g *Network) GetNode(id string) *Node {
	for _, node := range g.Nodes {
		if node.Id == id {
			return node
		}
	}
	return nil
}

func (g *Network) SendMessage(id1 string, id2 string, message string) []*utils.PackageResultInfo {
	packages := utils.SplitStringByLength(message, g.MaxPackageLength)
	var progress []*utils.PackageResultInfo
	genId := 0
	for _, packag := range packages {
		wg.Add(1)
		genId = genId + 1
		progressData := utils.PackageResultInfo{Message: packag, Status: "En progreso", Id: genId, Estimated: 0}
		progress = append(progress, &progressData)
		go g.simulateSend(id1, id2, &progressData)
	}
	gui.PrintSendResults(&progress)
	wg.Wait()
	return progress
}

func (g *Network) simulateSend(id1 string, id2 string, progressData *utils.PackageResultInfo) {
	defer wg.Done()
	mutex.Lock()
	resultData := g.CalculateShortestPath(id1, id2)
	mutex.Unlock()
	progressData.Estimated = resultData.Shortest
	progressData.PathData = resultData
	progressData.Shortest = 0
	g.recalculateVerticesWeight()
	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				progressData.Shortest++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	time.Sleep(time.Duration(progressData.Estimated) * time.Second)
	close(quit)
	progressData.Status = "Finalizado"
}

func (g *Network) recalculateVerticesWeight() {
	for _, node := range g.Nodes {
		for _, vertex := range node.Vertices {
			vertex.Weight = rand.Intn(15)
		}
	}
}
