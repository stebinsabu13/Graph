package main

import "fmt"

var graph = make(map[int][]int)

//adding a vertex

func addVertex(key int) {
	if contains(key) {
		fmt.Println("Vertex already exsists")
		return
	}
	graph[key] = []int{}
}

//adding an edge

func addEdge(from, to int) {
	if !contains(from) || !contains(to) {
		fmt.Println("Vertex doesn't exsists to add edge")
	} else if exsistedge(from, to) {
		fmt.Println("Edge already exsists")
	} else {
		graph[from] = append(graph[from], to)
	}
}

//checking if the edge exsists

func exsistedge(from, to int) bool {
	for _, v := range graph[from] {
		if v == to {
			return true
		}
	}
	return false
}

//checking if the vertex exsists

func contains(key int) bool {
	for v := range graph {
		if v == key {
			return true
		}
	}
	return false
}

//traversing with dfs

func depthSearch(key int, visited *[]int) {
	*visited = append(*visited, key)
	for _, v := range graph[key] {
		if !check(v, *visited) {
			depthSearch(v, visited)
		}
	}
}

//checking for if the vertex is already visited or not

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
