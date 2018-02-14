package main

import "fmt"

const MAXNODES = 1024

// unoriented graph, adjacency list
// nodes are from 0 to numNodes-1
type pair struct{
    neighbour int
    weight int
}

func main (){
    var inputGraph [MAXNODES][]pair
    var node1, node2, weight, numNodes, numEdges int

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

    for i:= 0; i<numNodes; i++{
        for j:=0; j<len(inputGraph[i]); j++{
            var node int = inputGraph[i][j].neighbour

            if node < i {
                continue
            }

            fmt.Printf ("the edge between %d %d is %d\n", i, node, inputGraph[i][j].weight)
         }
    }
}
