package main

import "fmt"

const MAXNODES = 1024

// oriented graph, adjacency list
// nodes are from 0 to numNodes-1

var AdjacencyList [MAXNODES][] int
var numNodes, numEdges int
var Capacity, CurrentFlow[MAXNODES][MAXNODES] int
var SourceNode, TargetNode int

var usedNodes [MAXNODES]bool
var isThereAPath bool

var augmentingPath []int

func makeEdge (node1, node2, capacity int){ 
    //makes the edges from node1 to node2 and its reverse one
    AdjacencyList[node1] = append (AdjacencyList[node1], node2)
    AdjacencyList[node2] = append (AdjacencyList[node2], node1)

    //sets the capacities
    Capacity[node1][node2] = capacity
    Capacity[node2][node1] = 0
}

func diffCapacityCurrentFlow (node1, node2 int) int {
    return Capacity[node1][node2] - CurrentFlow[node1][node2]
}

func makeFlow (){
    egdesFlow := diffCapacityCurrentFlow (SourceNode, augmentingPath[0])

    for i := 1; i < len (augmentingPath); i++ {
        current := diffCapacityCurrentFlow(augmentingPath[i-1], augmentingPath[i])

        if current < egdesFlow{
            egdesFlow = current
        }
    }

    CurrentFlow[SourceNode][augmentingPath[0]] += egdesFlow
    CurrentFlow[augmentingPath[0]][SourceNode] -= egdesFlow

    for i := 1; i < len (augmentingPath); i++ {
        CurrentFlow[augmentingPath[i-1]][augmentingPath[i]] += egdesFlow
        CurrentFlow[augmentingPath[i]][augmentingPath[i-1]] -= egdesFlow
    }
}

func dfs (currNode int) {
    if currNode == TargetNode{
        isThereAPath = true
        return
    }

    for i := 0; i < len(AdjacencyList[currNode]); i++ {
        neighbourNode := AdjacencyList[currNode][i]

        if usedNodes[neighbourNode] == false && CurrentFlow[currNode][neighbourNode] < Capacity[currNode][neighbourNode]{
            usedNodes[neighbourNode] = true
            augmentingPath = append (augmentingPath, neighbourNode)

            dfs (neighbourNode)
            if isThereAPath == true{
                return
            }

            augmentingPath = augmentingPath[:len(augmentingPath) - 1] 
        }
    }
}

func readGraph (){
    var node1, node2, weight int

    fmt.Scanf ("%d %d",&numNodes, &numEdges)

    for i := 0; i < numEdges; i++ {
        fmt.Scanf ("%d %d %d", &node1, &node2, &weight)

        makeEdge (node1, node2, weight)
    }

    fmt.Printf ("Which vertex is the Source?\n")
    fmt.Scanf ("%d", &SourceNode)

    fmt.Printf ("Which vertex is the Target?\n")
    fmt.Scanf ("%d", &TargetNode)
}

func main (){
    readGraph ()

    //as long as there is a path from the source (start node) to the target (end node)
    // with available capacity on all edges in the path, we send flow along one of the paths.

    for {
        //Clearing the previous dfs data
        isThereAPath = false
        augmentingPath = nil

        for i := 0; i < MAXNODES; i++{
            usedNodes[i] = false
        }

        //New path search
        usedNodes[SourceNode] = true
        dfs (SourceNode)

        if isThereAPath == false {
            break
        }

        makeFlow ()
    }

    var answer int

    for i := 0; i < len (AdjacencyList[TargetNode]); i++ {
        neighbour := AdjacencyList[TargetNode][i]
        answer += CurrentFlow[neighbour][TargetNode]
    }

    fmt.Printf ("The maximum flow from Source to Target is %d\n", answer)
}
