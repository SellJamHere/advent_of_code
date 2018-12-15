package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(walkGraph(createGraph(p1)))
	fmt.Println(walkGraph(createGraph(input)))

	fmt.Println(timeExecution(createGraph(p1), 2, 0))
	fmt.Println(timeExecution(createGraph(input), 5, 60))
}

func createGraph(input string) map[string][]string {
	inverseGraph := map[string][]string{}

	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")
		dependency := words[1]
		dependant := words[7]

		inverseGraph[dependant] = append(inverseGraph[dependant], dependency)
		if _, ok := inverseGraph[dependency]; !ok {
			inverseGraph[dependency] = []string{}
		}
	}

	for _, dependancyList := range inverseGraph {
		sort.Strings(dependancyList)
	}

	return inverseGraph
}

func walkGraph(inverseGraph map[string][]string) string {
	order := ""

	for {
		availableNodes := findAvailableNodes(inverseGraph)
		if len(availableNodes) == 0 {
			break
		}

		currentNode := availableNodes[0]
		order += currentNode

		// remove node
		deleteNode(inverseGraph, currentNode)
	}

	return order
}

func findAvailableNodes(inverseGraph map[string][]string) []string {
	// find nodes without dependancies
	freeNodes := []string{}
	for node, dependencies := range inverseGraph {
		if len(dependencies) == 0 {
			freeNodes = append(freeNodes, node)
		}
	}

	if len(freeNodes) == 0 {
		return nil
	}

	sort.Strings(freeNodes)
	return freeNodes
}

func deleteNode(inverseGraph map[string][]string, currentNode string) {
	delete(inverseGraph, currentNode)
	for node, dependencies := range inverseGraph {
		// create a zero-length slice with the same underlying array
		tmp := []string{}
		for _, dependency := range dependencies {
			if dependency != currentNode {
				tmp = append(tmp, dependency)
			}
		}

		inverseGraph[node] = tmp
	}
}

func timeExecution(inverseGraph map[string][]string, parallelization int, baseTime int) int {
	type work struct {
		node string
		time int
	}
	// map of worker to time left
	workers := map[int]work{}
	for i := 0; i < parallelization; i++ {
		workers[i] = work{}
	}

	nodesInProgress := map[string]struct{}{}

	totalTime := 0
	for {
		fmt.Println(totalTime, "initial workers:", workers)

		// reduce worker times
		availableWorkers := []int{}
		for i, w := range workers {
			if w.time > 0 {
				w.time--
				workers[i] = w
			} else if w.time == 0 {
				availableWorkers = append(availableWorkers, i)
				workers[i] = work{}
				deleteNode(inverseGraph, w.node)
				delete(nodesInProgress, w.node)
			}
		}

		availableNodes := findAvailableNodes(inverseGraph)
		if len(availableNodes) == 0 {
			zeroCount := 0
			for _, work := range workers {
				if work.time > 0 {
					work.time--
				} else {
					zeroCount++
				}
			}

			if zeroCount == len(workers) {
				break
			}

			totalTime++
			continue
		}

		// assign work
		workerIndex := 0
		for _, currentNode := range availableNodes {
			fmt.Println("available nodes:", availableNodes)
			if workerIndex >= len(availableWorkers) {
				break
			}

			if _, ok := nodesInProgress[currentNode]; ok {
				continue
			}

			nodeTime := int([]rune(currentNode)[0]) - 65 + baseTime
			workers[availableWorkers[workerIndex]] = work{currentNode, nodeTime}

			nodesInProgress[currentNode] = struct{}{}
			workerIndex++
		}

		fmt.Println(totalTime, "final workers:", workers)

		totalTime++
	}

	return totalTime
}

var p1 = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

var input = `Step Z must be finished before step V can begin.
Step V must be finished before step K can begin.
Step M must be finished before step Q can begin.
Step E must be finished before step X can begin.
Step J must be finished before step W can begin.
Step L must be finished before step O can begin.
Step Q must be finished before step T can begin.
Step Y must be finished before step P can begin.
Step X must be finished before step R can begin.
Step T must be finished before step U can begin.
Step I must be finished before step O can begin.
Step P must be finished before step H can begin.
Step G must be finished before step A can begin.
Step N must be finished before step A can begin.
Step H must be finished before step B can begin.
Step F must be finished before step D can begin.
Step S must be finished before step O can begin.
Step O must be finished before step W can begin.
Step D must be finished before step U can begin.
Step W must be finished before step B can begin.
Step A must be finished before step K can begin.
Step B must be finished before step R can begin.
Step K must be finished before step C can begin.
Step R must be finished before step C can begin.
Step U must be finished before step C can begin.
Step A must be finished before step U can begin.
Step J must be finished before step I can begin.
Step D must be finished before step K can begin.
Step V must be finished before step S can begin.
Step H must be finished before step C can begin.
Step R must be finished before step U can begin.
Step I must be finished before step G can begin.
Step D must be finished before step R can begin.
Step M must be finished before step B can begin.
Step G must be finished before step R can begin.
Step M must be finished before step I can begin.
Step G must be finished before step N can begin.
Step M must be finished before step N can begin.
Step Q must be finished before step S can begin.
Step I must be finished before step S can begin.
Step J must be finished before step R can begin.
Step O must be finished before step B can begin.
Step G must be finished before step S can begin.
Step J must be finished before step C can begin.
Step M must be finished before step D can begin.
Step T must be finished before step H can begin.
Step P must be finished before step N can begin.
Step S must be finished before step K can begin.
Step T must be finished before step C can begin.
Step J must be finished before step A can begin.
Step G must be finished before step F can begin.
Step N must be finished before step R can begin.
Step N must be finished before step W can begin.
Step T must be finished before step I can begin.
Step S must be finished before step B can begin.
Step H must be finished before step F can begin.
Step B must be finished before step C can begin.
Step L must be finished before step W can begin.
Step N must be finished before step O can begin.
Step O must be finished before step A can begin.
Step H must be finished before step S can begin.
Step F must be finished before step A can begin.
Step F must be finished before step C can begin.
Step M must be finished before step A can begin.
Step Z must be finished before step H can begin.
Step Z must be finished before step L can begin.
Step E must be finished before step H can begin.
Step X must be finished before step T can begin.
Step Y must be finished before step X can begin.
Step E must be finished before step W can begin.
Step P must be finished before step R can begin.
Step Z must be finished before step E can begin.
Step W must be finished before step C can begin.
Step I must be finished before step P can begin.
Step X must be finished before step A can begin.
Step Y must be finished before step C can begin.
Step I must be finished before step F can begin.
Step L must be finished before step T can begin.
Step A must be finished before step B can begin.
Step F must be finished before step W can begin.
Step T must be finished before step R can begin.
Step X must be finished before step F can begin.
Step M must be finished before step O can begin.
Step N must be finished before step K can begin.
Step T must be finished before step S can begin.
Step J must be finished before step N can begin.
Step J must be finished before step S can begin.
Step O must be finished before step D can begin.
Step T must be finished before step P can begin.
Step Z must be finished before step D can begin.
Step L must be finished before step X can begin.
Step Q must be finished before step G can begin.
Step M must be finished before step G can begin.
Step P must be finished before step W can begin.
Step V must be finished before step P can begin.
Step D must be finished before step B can begin.
Step Y must be finished before step D can begin.
Step X must be finished before step S can begin.
Step K must be finished before step U can begin.
Step Z must be finished before step Y can begin.
Step D must be finished before step W can begin.`
