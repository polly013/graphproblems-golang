package main

import "fmt"

const MAXNODES = 1024

// unoriented graph, adjacency matrix
// nodes are from 0 to numNodes-1
type graph struct{
    numNodes int
    edges [MAXNODES][MAXNODES]int
}

func main (){
    var inputGraph graph
    var node1, node2, weight, numNodes, numEdges int

    fmt.Scanf ("%d %d",&numNodes, &numEdges)
    inputGraph.numNodes = numNodes

    for i := 0; i<numEdges; i++ {
        fmt.Scanf ("%d %d %d", &node1, &node2, &weight)
        inputGraph.edges[node1][node2] = weight
        inputGraph.edges[node2][node1] = weight
    }

    for i:= 0; i<numNodes; i++{
        for j:= i; j<numNodes; j++{
            if inputGraph.edges[i][j] != 0 {
                fmt.Printf ("the edge between %d %d is %d\n", i, j, inputGraph.edges[i][j])
            }
         }
    }
}
