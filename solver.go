package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func DFS(node *Node, goal float32) *Node {

	// Return node if goal is reached
	if node.expressionValue == goal {
		return node
	}

	node.generateChildren()

	var result *Node

	// For every child, call dfs
	for _, cn := range node.children{
		result = DFS(cn, goal)

		// When node is being return from head, result is set. Result first has to be checked whether it is nul, then check goal
		//and return node so that node can be returned from recursive function
		if result != nil {
			if result.expressionValue == goal {
				defer wg.Done()
				break
			}
		}
	}

	return result
}

func aStar(rootNode *Node, goal float32) *Node {

	// Keep track of a general list of nodes, and add list of children to priorityQueue
	var priorityQueue PriorityQueue
	var currentNode *Node
	var evaluatedExpressions []string
	var processedNodes []*Node

	visitedNodeCounter := 0 // Keep track of counter instead of nodeList and added property hasBeenVisisted to node itself to save memory

	// Generate all children of root node, and add children to priority queue
	rootNode.generateRootChildren()

	// Add children of root node to priority queue
	priorityQueue = rootNode.children

	mainflag:
	for len(priorityQueue) > 0 {

		// Sort priorityQueue
		sortByDistanceFrom(goal, priorityQueue)

		// Generate random number between 1 and 10 such that 10% of the time it will pick a random node
		rand.Seed(time.Now().UnixNano())
		randNode := rand.Intn(11)

		// Pick random node or first node in priorityQueue
		if randNode == 3{
			currentNode = priorityQueue[rand.Intn(len(priorityQueue))]
		} else {
			currentNode = priorityQueue[0]
		}

		// Check if value and numbers left of current node is equal to processedNode, if yes, skip node and remove from priorityqueue
		for _, processedNode := range processedNodes{
			if (floatSlicesEqual(processedNode.numbersLeft, currentNode.numbersLeft) && processedNode.expressionValue == currentNode.expressionValue) || currentNode.expression == processedNode.expression {
				priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
				continue mainflag
			}
		}

		// Check if current node is goal
		if currentNode.expressionValue == goal && len(currentNode.numbersLeft) == 0 {
			println(ColorWhite, fmt.Sprintf("\nProblem solved: %v => [%v]\nSolution: %v\nNumber of visited nodes: %v", rootNode.numbersLeft, goal, currentNode.expression, visitedNodeCounter), ColorGreen, "\n[*] Thread finished.")
			defer wg.Done()
			return currentNode
		}

		// Generate children
		currentNode.generateChildren()

		// Add expression
		evaluatedExpressions = append(evaluatedExpressions, currentNode.expression)

		// Add children of current node to priority queue
		for _, cn := range currentNode.children{
			priorityQueue = append(priorityQueue, cn)
		}

		visitedNodeCounter++

		priorityQueue = priorityQueue.removeNodeFromPriorityQueue(currentNode)
		processedNodes = append(processedNodes, currentNode)
	}

	defer wg.Done()
	return &Node{}
}

func sortByDistanceFrom(goal float32, nodes []*Node) {
	sort.Slice(nodes, func(i, j int) bool {
		di := math.Abs(float64(goal - nodes[i].expressionValue))
		dj := math.Abs(float64(goal - nodes[j].expressionValue))
		return di < dj
	})
}

func floatSlicesEqual(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

