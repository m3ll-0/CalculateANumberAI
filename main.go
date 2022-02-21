package main

import (
	"fmt"
	"time"
)

func main(){

	// List of problems that need to be solved
	problems := []Problem{
		//{[]float32{30, 10, 5, 2, 1, 3}, 17},
		//{[]float32{3, 3, 3, 3, 3, 3, 3, 2, 10}, 30},
		//{[]float32{3, 3, 3, 3, 3, 3, 3, 2, 2}, 5},
		//{[]float32{3, 3, 3, 3, 3, 3, 3, 2, 10}, 300},
		//{[]float32{3, 3, 3, 3, 3, 3, 3}, 5},
		//{[]float32{1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}, 1},
		{[]float32{2, 3, 599, 23, 29, 53, 71}, 593},
	}

	solverMethod := aStar
	runAI(problems, solverMethod)
}

func runAI(problems []Problem, solverMethod SolverMethod){

	// Keep track of time
	var timeStart = time.Now()

	// Solve each problem
	println(ColorGreen, "[*] Starting threads.")
	for _, problem := range problems {
		wg.Add(1)
		rootNode := &Node{expressionValue: 0, expression: "0"}
		rootNode.numbersLeft = problem.numberList
		go solverMethod(rootNode, problem.goal)
	}

	wg.Wait()

	println(ColorGreen, fmt.Sprintf("\n[*] Problems solved. Total time: %v", time.Since(timeStart)))
}
