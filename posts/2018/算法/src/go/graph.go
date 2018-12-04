package main

import (
	"container/list"
	"fmt"
)

// 创建无向图邻接表
func createAdjlist(size int) []list.List {
	adj := make([]list.List, size)
	return adj
}

// 无向图添加元素
func addUndirectEdge(adj []list.List, s, t int) {
	adj[s].PushBack(t)
	adj[t].PushBack(s)
}

// 有向图添加元素
func addEdge(adj []list.List, s, t int) {
	adj[s].PushBack(t)
}

func bfs(adj []list.List, s, t int) {
	if s == t {
		return
	}

	size := len(adj)
	// 访问过的顶点
	visited := make([]bool, size)
	visited[s] = true

	// 存储已经被访问，但相连的顶点还没有被访问的顶点
	queue := list.New()
	queue.PushBack(s)

	// 访问路径
	prev := make([]int, size)

	for i := 0; i < size; i++ {
		prev[i] = -1
	}

	fmt.Println("bfs: ")
	for queue.Len() != 0 {
		// 队首出队
		w := queue.Remove(queue.Front()).(int)

		// 遍历相关联的顶点
		for q := adj[w].Front(); q != nil; q = q.Next() {
			val := q.Value.(int)
			// 该路径没有访问过
			if !visited[val] {
				prev[val] = w
				if val == t {
					print(prev, s, t)
					return
				}

				// 设置给顶点已经访问过
				visited[val] = true
				// 入队
				queue.PushBack(val)
			}
		}
	}
}

func print(prev []int, s, t int) {
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Print(t, " ")
}

var found = false

func dfs(adj []list.List, s, t int) {
	found = false
	size := len(adj)

	visited := make([]bool, size)

	// 访问路径
	prev := make([]int, size)

	for i := 0; i < size; i++ {
		prev[i] = -1
	}

	recurDfs(adj, s, t, visited, prev)
	fmt.Println("\ndfs:")
	print(prev, s, t)
}

func recurDfs(adj []list.List, w, t int, visited []bool, prev []int) {
	if found {
		return
	}

	visited[w] = true
	if w == t {
		found = true
		return
	}

	for q := adj[w].Front(); q != nil; q = q.Next() {
		val := q.Value.(int)
		if !visited[val] {
			prev[val] = w
			recurDfs(adj, val, t, visited, prev)
		}
	}
}

func main() {
	adj := createAdjlist(8)
	addUndirectEdge(adj, 0, 3)
	addUndirectEdge(adj, 0, 1)
	addUndirectEdge(adj, 0, 2)

	addUndirectEdge(adj, 1, 4)
	addUndirectEdge(adj, 1, 5)

	addUndirectEdge(adj, 3, 7)
	addUndirectEdge(adj, 3, 6)

	addUndirectEdge(adj, 6, 4)

	bfs(adj, 0, 4)
	dfs(adj, 0, 4)
}
