# Graph
Implementation of graph with insertion and depth search using map in golang 


package main

import "fmt"

var graph = make(map[int][]int)

func addVertex(key int) {
	if contains(key) {
		fmt.Println("Vertex already exsists")
		return
	}
	graph[key] = []int{}
}
func addEdge(from, to int) {
	if !contains(from) || !contains(to) {
		fmt.Println("Vertex doesn't exsists to add edge")
	} else if exsistedge(from, to) {
		fmt.Println("Edge already exsists")
	} else {
		graph[from] = append(graph[from], to)
	}
}
func exsistedge(from, to int) bool {
	for _, v := range graph[from] {
		if v == to {
			return true
		}
	}
	return false
}
func contains(key int) bool {
	for v := range graph {
		if v == key {
			return true
		}
	}
	return false
}
func depthSearch(key int, visited *[]int) {
	*visited = append(*visited, key)
	for _, v := range graph[key] {
		if !check(v, *visited) {
			depthSearch(v, visited)
		}
	}
}
func check(key int, visited []int) bool {
	for _, v := range visited {
		if v == key {
			return true
		}
	}
	return false
}
func main() {
	for i := 0; i < 5; i++ {
		addVertex(i)
	}
	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(1, 2)
	addEdge(2, 0)
	addEdge(2, 3)
	addEdge(3, 3)
	addEdge(3, 4)
	var visited []int
	depthSearch(2, &visited)
	fmt.Println(visited)
}
