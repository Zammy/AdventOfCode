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

func (n *node) CanBeRevisited() bool {
	return n.Name == strings.ToUpper(n.Name)
}

func (n *node) FindAllPathsToEnd(visted []string) int {
	visted = append(visted, n.Name)

	if n.Name == "end" {
		return 1
	}

	paths := 0
	for _, connected := range n.Paths {
		if !connected.CanBeRevisited() && contains(visted, connected.Name) {
			continue
		}
		visitedCopy := make([]string, len(visted))
		copy(visitedCopy, visted)
		paths += connected.FindAllPathsToEnd(visitedCopy)
	}
	return paths
}

func contains(input []string, has string) bool {
	for _, v := range input {
		if v == has {
			return true
		}
	}
	return false
}
