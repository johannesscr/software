package challenges

import (
	"fmt"
	"testing"
)

/*
Given an m x n 2D binary grid `grid` which represents a map of '1's (land) and
'0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands
horizontally or vertically. You may assume all four edges of the grid are all
surrounded by water.

Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

/*
NOTES:

1. We see from example two that an island's land is only connected horizontally
   or vertically and NOT diagonally.

*/

/*
SOLUTIONS:

1. loop through the entire map, check each index.
   1. If the index has no adjacent island give it a new sequential name.
   2. If it has an island next to it (vertically or horizontally), give it the
      same name.
   3. While naming keep a list/array/slice of all the island names given
2. return the list of island names.
*/

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

// NumIslands steps through the grid and numbers each piece of land according
// to which island is associated with, and returns the total number of islands
// in the grid.
//
// Note that the inputs are strings, which means that when they are converted
// bytes they are converted using ASCII, which means that "0" -> 48 and
// "1" -> 49.
func NumIslands(grid [][]byte) int {
	//printGrid(grid)
	islands := make([]int, 0)
	// no rows
	if len(grid) == 0 {
		return 0
	}
	nRows := len(grid)
	// no columns
	if len(grid[0]) == 0 {
		return 0
	}
	nCols := len(grid[0])
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			// the grid index is a piece of land
			isLand := string(grid[i][j]) == "1"
			// find the north and west neighbours
			north, west := findNW(i, j, grid)
			neighbour := "-1"
			if north > "0" {
				neighbour = north
			}
			if west > "0" {
				neighbour = west
			}
			if isLand && neighbour > "-1" {
				grid[i][j] = []byte(neighbour)[0]
			}
			if isLand && neighbour == "-1" {
				newIsland := len(islands) + 49 // offset of 48+1
				// we ensure to skip 48 and 49 which are the ascii character
				// values for 0 and 1.
				islands = append(islands, newIsland)
				grid[i][j] = byte(newIsland)
			}
		}
	}
	//printGrid(grid)
	return len(islands)
}

// findNW returns the possible North and West neighbour islands for a given
// coordinate, where neighbours are denoted by their integer index. If no
// neighbour exists -1 is returned.
//
// We only need North West, since we are stepping through the grid in a South
// East manner, meaning for a coordinate i, j South and East will always be
// unknown.
func findNW(i, j int, islandGrid [][]byte) (string, string) {
	// 0 denotes no neighbour
	north := "-1"
	west := "-1"
	// we check the index is in the grid
	if i-1 >= 0 {
		north = string(islandGrid[i-1][j])
	}
	// we check the index is in the grid
	if j-1 >= 0 {
		west = string(islandGrid[i][j-1])
	}
	return north, west
}

//func TestFindNW(t *testing.T) {
//	grid := [][]byte{
//		{1, 2},
//		{3, 0},
//	}
//	tests := []struct {
//		name string
//		grid [][]byte
//		i    int
//		j    int
//		out1 string
//		out2 string
//	}{
//		{
//			name: "bottom right",
//			grid: grid,
//			i:    1,
//			j:    1,
//			out1: "2",
//			out2: "3",
//		},
//		{
//			name: "top right",
//			grid: grid,
//			i:    0,
//			j:    1,
//			out1: "-1",
//			out2: "1",
//		},
//		{
//			name: "top left",
//			grid: grid,
//			i:    0,
//			j:    0,
//			out1: "-1",
//			out2: "-1",
//		},
//		{
//			name: "bottom left",
//			grid: grid,
//			i:    1,
//			j:    0,
//			out1: "1",
//			out2: "-1",
//		},
//	}
//
//	for i, tc := range tests {
//		name := fmt.Sprintf("%d %s", i, tc.name)
//		t.Run(name, func(t *testing.T) {
//			o1, o2 := findNW(tc.i, tc.j, grid)
//			if tc.out1 != o1 {
//				t.Errorf("expected north '%s' got '%s'", tc.out1, o1)
//			}
//			if tc.out2 != o2 {
//				t.Errorf("expected west '%s' got '%s'", tc.out2, o2)
//			}
//		})
//	}
//}

func convertGrid(grid [][]string) [][]byte {
	var rowLength = len(grid)
	var columnLength = len(grid[0])
	var xGrid = make([][]byte, rowLength)
	for i := range grid {
		var row = make([]byte, columnLength)
		for j, v := range grid[i] {
			// this will only ever be a single character
			row[j] = []byte(v)[0]
		}
		xGrid[i] = row
	}
	return xGrid
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name       string
		grid       [][]byte
		numIslands int
	}{
		{
			name: "example one",
			grid: convertGrid([][]string{
				{"1", "1", "1", "1", "0"},
				{"1", "1", "0", "1", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "0", "0", "0"},
			}),
			numIslands: 1,
		},
		{
			name: "example two",
			grid: convertGrid([][]string{
				{"1", "1", "0", "0", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "1", "0", "0"},
				{"0", "0", "0", "1", "1"},
			}),
			numIslands: 3,
		},
		//{
		//	name: "example three",
		//	grid: convertGrid([][]string{
		//		{"1", "1", "1"},
		//		{"0", "1", "0"},
		//		{"1", "1", "1"},
		//	}),
		//	numIslands: 1,
		//},
	}

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			n := NumIslands(tc.grid)
			if n != tc.numIslands {
				t.Errorf("expected number islands %d got %d", tc.numIslands, n)
			}
		})
	}
}

/* Solution Depth First Search */

/*
Linear scan the 2d grid map, if a node contains a '1', then it is a root node
that triggers a Depth First Search. During DFS, every visited node should be set
as '0' to mark as visited node. Count the number of root nodes tha trigger DFS,
this number would be the number of islands since each DFS starting at some root
identifies an island.
*/

// dfs does a depth first search, finds all children nodes that are a "1" and
// turns them into a "0".
func dfs(grid *[][]byte, r, c int) {
	var g = *grid
	var nRow = len(g)
	var nCol = len(g[0])

	g[r][c] = 48 // 48 => "0"
	// if the node above this node is land then it also need to be removed.
	if r-1 >= 0 && g[r-1][c] == 49 {
		dfs(grid, r-1, c)
	}
	// if the node below this node is land then it also need to be removed.
	if r+1 < nRow && g[r+1][c] == 49 {
		dfs(grid, r+1, c)
	}
	// if the node to the left is land then it also needs to be removed.
	if c-1 >= 0 && g[r][c-1] == 49 {
		dfs(grid, r, c-1)
	}
	// if the node to the right is land then it also needs to be removed.
	if c+1 < nCol && g[r][c+1] == 49 {
		dfs(grid, r, c+1)
	}
}

// numIslandsDFS returns the number of islands in a grid.
func numIslandsDFS(grid [][]byte) int {
	nRow := len(grid)
	if nRow == 0 {
		return 0
	}
	nCol := len(grid[0])

	numIslands := 0
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			if grid[i][j] == 49 { // same as grid[i][j] == "1"
				numIslands++
				dfs(&grid, i, j)
			}
		}
	}
	return numIslands
}

func TestNumIslandsDFS(t *testing.T) {
	tests := []struct {
		name       string
		grid       [][]byte
		numIslands int
	}{
		{
			name: "example one",
			grid: convertGrid([][]string{
				{"1", "1", "1", "1", "0"},
				{"1", "1", "0", "1", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "0", "0", "0"},
			}),
			numIslands: 1,
		},
		{
			name: "example two",
			grid: convertGrid([][]string{
				{"1", "1", "0", "0", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "1", "0", "0"},
				{"0", "0", "0", "1", "1"},
			}),
			numIslands: 3,
		},
		{
			name: "example three",
			grid: convertGrid([][]string{
				{"1", "1", "1"},
				{"0", "1", "0"},
				{"1", "1", "1"},
			}),
			numIslands: 1,
		},
	}

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			n := numIslandsDFS(tc.grid)
			if n != tc.numIslands {
				t.Errorf("expected number islands %d got %d", tc.numIslands, n)
			}
		})
	}
}
