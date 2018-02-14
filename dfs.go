package main

import "fmt"

const MAXNODES = 1024

// unoriented graph, adjacency list
// nodes are from 0 to numNodes-1
type pair struct{
    neighbour int
    weight int
}

var inputGraph [MAXNODES][]pair
var numNodes, numEdges int
var usedNodes [MAXNODES]bool

func dfs (currNode int) {
//    fmt.Printf ("Jumped in %d\n", currNode)

    for i := 0; i<len(inputGraph[currNode]); i++{
        var neighbourNode = inputGraph[currNode][i].neighbour;

        if usedNodes[neighbourNode] == false{
            usedNodes[neighbourNode] = true;

//            fmt.Printf ("Using the edge to %d with weight %d\n", neighbourNode, inputGraph[currNode][i].weight)
            dfs (neighbourNode);

//            fmt.Printf ("Back in %d\n", currNode)
        }
    }
}

func main (){
    var node1, node2, weight int

    fmt.Scanf ("%d %d",&numNodes, &numEdges)

    for i := 0; i<numEdges; i++ {
        fmt.Scanf ("%d %d %d", &node1, &node2, &weight)

        var edge pair
        edge.neighbour = node2
        edge.weight = weight

        inputGraph[node1] = append (inputGraph[node1], edge)

        edge.neighbour = node1
        inputGraph[node2] = append (inputGraph[node2], edge)
    }

    var startNode int

    fmt.Scanf ("%d", &startNode)

    usedNodes[startNode] = true
    dfs (startNode)
}
