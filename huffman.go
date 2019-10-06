package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"unicode/utf8"
)

type Node struct {
	char rune
	freq int
	left *Node
	right *Node
}

type CodeTable = map[rune]string

type PQ []Node

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PQ) Push(x interface{}) {
	node := x.(Node)
	*pq = append(*pq, node)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func CovertStringToNodes(s string) *[]Node {
	freq := make(map[rune]int, utf8.RuneCountInString(s))
	for _, c := range s {
		freq[c] += 1
	}

	nodes := make([]Node, 0, utf8.RuneCountInString(s))
	for k, v := range freq {
		nodes = append(nodes, Node{
			char: k,
			freq: v,
			right: nil,
			left: nil,
		})
	}

	return &nodes
}

func PrintNode(node Node) {
	fmt.Printf("Name: %c Expiry: %d Left: %v Right: %v\n", node.char, node.freq, node.left, node.right)
}

func BuildTree(nodes *[]Node) *Node {
	priorityQueue := make(PQ, len(*nodes))
	for i, node := range *nodes {
		priorityQueue[i] = node
	}
	heap.Init(&priorityQueue)
	for priorityQueue.Len() > 1 {
		node1 := heap.Pop(&priorityQueue).(Node)
		node2 := heap.Pop(&priorityQueue).(Node)
		newNode := Node{
			freq: node1.freq + node2.freq,
			left: &node1,
			right: &node2,
		}
		heap.Push(&priorityQueue, newNode)
		//fmt.Printf("Name: %c Expiry: %d Left: %v Right: %v\n", node.char, node.freq, node.left, node.right)
	}
	root := heap.Pop(&priorityQueue).(Node)
	return &root
}

func CreateCodeTableRec(root *Node, code string, table *CodeTable) {
	var undefined rune
	if root != nil {
		if root.char != undefined {
			(*table)[root.char] = code
		} else {
			CreateCodeTableRec(root.left, code + "0", table)
			CreateCodeTableRec(root.right, code + "1", table)
		}
	}
}

func CreateCodeTable(root *Node) *CodeTable {
	table := make(CodeTable)
	var undefined rune

	if root.char != undefined {
		table[root.char] = "0"
	} else {
		CreateCodeTableRec(root.left, "0", &table)
		CreateCodeTableRec(root.right, "1", &table)
	}

	return &table
}

func EncodeString(s string, table *CodeTable) string {
	res := ""
	for _, r := range s {
		res += (*table)[r]
	}
	return res
}

func Hello() string {
	return "Hello"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	nodes := CovertStringToNodes(s)
	root := BuildTree(nodes)
	table := CreateCodeTable(root)
	res := EncodeString(s, table)

	fmt.Printf("%d %d\n", len(*table), len(res))
	for k, v := range *table {
		fmt.Printf("%c: %v\n", k, v)
	}
	fmt.Printf(res)
}
