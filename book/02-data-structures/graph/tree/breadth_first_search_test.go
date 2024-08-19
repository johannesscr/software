package tree

import (
	"fmt"
	"testing"
)

var nodes = map[string]*Node{
	"PHX": {Val: "PHX"},
	"LAX": {Val: "LAX"},
	"JFK": {Val: "JFK"},
	"OKC": {Val: "OKC"},
	"HEL": {Val: "HEL"},
	"LOS": {Val: "LOS"},
	"MEX": {Val: "MEX"},
	"BKK": {Val: "BKK"},
	"LIM": {Val: "LIM"},
	"EZE": {Val: "EZE"},
}

func init() {
	fmt.Println("created tree")
	nodes["PHX"].Children = append(nodes["PHX"].Children, nodes["LAX"], nodes["JFK"])
	nodes["JFK"].Children = append(nodes["JFK"].Children, nodes["OKC"], nodes["HEL"], nodes["LOS"])
	nodes["OKC"].Children = append(nodes["OKC"].Children, nodes["BKK"])
	nodes["MEX"].Children = append(nodes["MEX"].Children, nodes["LAX"], nodes["BKK"], nodes["LIM"], nodes["EZE"])
	nodes["LIM"].Children = append(nodes["LIM"].Children, nodes["BKK"])
	fmt.Println("tree created")
}

func TestNode_BFS(t *testing.T) {
	tree := nodes["PHX"]
	o := tree.BFS("HEL")
	output := ""
	outStr := fmt.Sprintf("%v", o)
	if outStr != output {
		t.Errorf("expected output %s got %s", output, outStr)
	}
}

func TestNode_DFS(t *testing.T) {
	tree := nodes["PHX"]
	o := tree.DFS("HEL")
	output := ""
	outStr := fmt.Sprintf("%v", o)
	if outStr != output {
		t.Errorf("expected output %s got %s", output, outStr)
	}
}
