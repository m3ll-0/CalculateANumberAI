package main

import "strconv"

type Node struct {
	level int
	parent *Node
	children []*Node
	operator string
	expressionValue float32
	expression string
	numbersLeft []float32
}

func (node *Node) generateRootChildren(){

	// Numbers that are processed
	numbersProcessed := []float32{}

	// Generate all combinations of expressions
	for _, numberLeft := range node.numbersLeft { // For each expression left

			// Create child node
			childNode := &Node{
				level: node.level + 1,
				parent: node,
				operator: "+",
				expression: node.expression + "+" + strconv.Itoa(int(numberLeft)),
				numbersLeft: node.removeNumberFromList(numberLeft),
			}

			// Use current value in combination with number and operator left to get a new value
			childNode.evaluate(numberLeft)

			// Add child node to current node
			node.children = append(node.children, childNode)

			// Add number to numberProcessed list
			numbersProcessed = append(numbersProcessed, numberLeft)
	}
}

func (node *Node) generateChildren(){

	// Generate all combinations of expressions
	for _, numberLeft := range node.numbersLeft { // For each expression left
		for _, operator := range operators { // For each operator

			// Skip /1 as /1 is the same as *1
			if operator == "/" && numberLeft == 1 {
				continue
			}

			// Create child node
			childNode := &Node{
				level: node.level + 1,
				parent: node,
				operator: operator,
				expression: node.expression + operator + strconv.Itoa(int(numberLeft)),
				numbersLeft: node.removeNumberFromList(numberLeft),
			}

			// Use current value in combination with number and operator left to get a new value
			childNode.evaluate(numberLeft)

			// Add child node to current node
			node.children = append(node.children, childNode)
		}
	}
}

func (node *Node) evaluate(numberLeft float32){
	if(node.operator) == "+" {
		node.expressionValue = node.parent.expressionValue + numberLeft
	} else if (node.operator) == "-"{
		node.expressionValue = node.parent.expressionValue - numberLeft
	} else if (node.operator) == "/" && node.parent.expressionValue != 0 {
		node.expressionValue = node.parent.expressionValue / numberLeft
	} else if (node.operator) == "*"{
		node.expressionValue = node.parent.expressionValue * numberLeft
	}
}

func (node *Node)removeNumberFromList(numberToBeRemoved float32) []float32 {

	newNumbersLeftList := []float32{}
	numberHasBeenRemoved := false

	for _, currentNumber := range node.numbersLeft {
		if currentNumber != numberToBeRemoved || numberHasBeenRemoved{
			newNumbersLeftList = append(newNumbersLeftList, currentNumber)
		} else {
			numberHasBeenRemoved = true
		}
	}

	return newNumbersLeftList

}

func (nodeList PriorityQueue)removeNodeFromPriorityQueue(nodeToBeRemoved *Node) []*Node {

	newList := []*Node{}

	for _, node := range nodeList{
		if node != nodeToBeRemoved {
			newList = append(newList, node)
		}
	}

	return newList
}