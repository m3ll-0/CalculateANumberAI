package main

import "sync"

type SolverMethod func(node *Node, goal float32) *Node
type PriorityQueue []*Node

var operators = []string{"+", "-", "*", "/"}
var wg sync.WaitGroup

// Colors to print
const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const ColorWhite = "\033[97m"