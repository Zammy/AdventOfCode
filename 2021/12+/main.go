package main

import (
	"fmt"
	"strings"

	"evgeni.com/util"
)

type node struct {
	Name  string
	Paths []*node
}

func main() {
	input, err := util.ReadLines("input")
	if err != nil {
		panic(err)
	}

	nodesMap := buildGraph(input)

	start := nodesMap["start"]
	paths := start.FindAllPathsToEnd([]string{})

	fmt.Println("Paths to end: ", paths)
}

func buildGraph(input []string) map[string]*node {
	nodesMap := make(map[string]*node)
	for _, row := range input {
		split := strings.Split(row, "-")

		name1 := split[0]
		node1 := nodesMap[name1]
		if node1 == nil {
			node1 = &node{Name: name1}
			nodesMap[name1] = node1
		}
		name2 := split[1]
		node2 := nodesMap[name2]
		if node2 == nil {
			node2 = &node{Name: name2}
			nodesMap[name2] = node2
		}
		node1.AddPath(node2)
		node2.AddPath(node1)
	}
	return nodesMap
}

func (n *node) AddPath(another *node) {
	if n.Paths == nil {
		n.Paths = []*node{another}
	} else {
		n.Paths = append(n.Paths, another)
	}
}

func (n *node) String() string {
	var sb strings.Builder
	for i, connected := range n.Paths {
		sb.WriteString(connected.Name)
		if i < len(n.Paths)-1 {
			sb.WriteString(",")
		}
	}
	return fmt.Sprintf("[%v] -> (%v)", n.Name, sb.String())
}

func (n *node) CanBeRevisited(visited []string) bool {
	if isBigCave := n.Name == strings.ToUpper(n.Name); isBigCave {
		return true
	}

	visits := countVisits(visited, n.Name)
	if n.IsStartOrEndNode() {
		return visits == 0
	}

	if twiceVisited := hasVisitedSmallCaveTwice(visited); !twiceVisited {
		return visits <= 1
	}

	return visits == 0
}

func (n *node) FindAllPathsToEnd(visited []string) int {
	visited = append(visited, n.Name)
	if n.Name == "end" {
		// fmt.Println(visited)
		return 1
	}

	paths := 0
	for _, connected := range n.Paths {
		if !connected.CanBeRevisited(visited) {
			continue
		}
		visitedCopy := make([]string, len(visited))
		copy(visitedCopy, visited)
		paths += connected.FindAllPathsToEnd(visitedCopy)
	}
	return paths
}

func (n *node) IsStartOrEndNode() bool {
	return n.Name == "start" || n.Name == "end"
}

func countVisits(visited []string, nodeName string) int {
	cnt := 0
	for _, visitedNodeName := range visited {
		if visitedNodeName == nodeName {
			cnt++
		}
	}
	return cnt
}

func hasVisitedSmallCaveTwice(visited []string) bool {
	for i := 0; i < len(visited); i++ {
		if isBigCave := visited[i] == strings.ToUpper(visited[i]); isBigCave {
			continue
		}
		for y := i + 1; y < len(visited); y++ {
			if visited[i] == visited[y] {
				return true
			}
		}
	}
	return false
}
