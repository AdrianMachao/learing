package algorithm

func mouseCat(graph [][]int) int {
	visited := make(map[int]map[int]bool)
	return trace(graph, 2, 1, visited)
}

func trace(graph [][]int, catPos, mousePos int, visited map[int]map[int]bool) int {
	if catPos == mousePos {
		return 2
	}
	if mousePos == 0 {
		return 1
	}

	nextCatPos := -1
	for _, v := range graph[catPos] {
		if visited[catPos][v] || v == 0 {
			continue
		}
		nextCatPos = v
		visited[catPos][v] = true
		break
	}

	nextMousePos := -1
	for _, v := range graph[mousePos] {
		if visited[mousePos][v] {
			continue
		}
		nextMousePos = v
		visited[mousePos][v] = true
		break
	}
	if nextCatPos == -1 || nextMousePos == -1 {
		return 0
	}

	return trace(graph, nextCatPos, nextMousePos, visited)
}
