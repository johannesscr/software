package challenges

import (
	"fmt"
	"testing"
)

/*
You are given an m x n grid rooms initialized with these three possible values.

- -1 A wall or an obstacle.
- 0 A gate.
- INF Infinity means an empty room. We use the value 2^31 - 1 = 2147483647 to
  represent INF as you may assume that the distance to a gate is less than
  2147483647.

Fill each empty room with the distance to its nearest gate. If it is impossible
to reach a gate, it should be filled with INF.

Constraints:

m == rooms.length
n == rooms[i].length
1 <= m, n <= 250
rooms[i][j] is -1, 0, or 231 - 1.
*/

// mental notes
// i = rows (north/south)
// j = columns (east/west)

//func traverseBFS(i, j int, rooms [][]int, level int) int {
//	// look up
//	if i-1 >= 0 {
//
//	}
//}

const (
	GATE  = 0
	EMPTY = 2147483647
)

const (
	North = iota
	South
	East
	West
)

func getDirectionCoords(i, j, rows, columns int, direction int) ([2]int, bool) {
	switch direction {
	case North:
		return [2]int{i - 1, j}, i-1 >= 0
	case South:
		return [2]int{i + 1, j}, i+1 < rows
	case East:
		return [2]int{i, j + 1}, j+1 < columns
	case West:
		return [2]int{i, j - 1}, j-1 >= 0
	default:
		return [2]int{0, 0}, false
	}
}

func WallsAndGates(rooms [][]int) [][]int {
	// Instead of searching from an empty room to the gates, how about searching
	// the other way round? In other words, we initiate breadth-first search
	// (BFS) from all gates at the same time. Since BFS guarantees that we
	// search all rooms of distance d before searching rooms of distance d + 1,
	// the distance to an empty room must be the shortest.
	rows := len(rooms)
	columns := len(rooms[0])
	// make a queue for all the gates
	q := make([][2]int, 0)
	// find all gates
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if rooms[i][j] == GATE {
				q = append(q, [2]int{i, j})
			}
		}
	}
	for len(q) > 0 {
		point := q[0]
		q = q[1:]
		row := point[0]
		col := point[1]
		// BFS in all directions
		for d := 0; d < 4; d++ {
			coords, ok := getDirectionCoords(row, col, rows, columns, d)
			r, c := coords[0], coords[1]
			if ok && rooms[r][c] == EMPTY {
				rooms[r][c] = rooms[row][col] + 1
				q = append(q, coords)
			} else {
				continue
			}
		}
	}
	return rooms
}

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		rooms [][]int
		out   [][]int
	}{
		{
			rooms: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			},
			out: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			rooms: [][]int{
				{-1},
			},
			out: [][]int{
				{-1},
			},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			o := WallsAndGates(tc.rooms)
			s1 := fmt.Sprintf("%v", o)
			s2 := fmt.Sprintf("%v", tc.out)
			if s1 != s2 {
				t.Errorf("expected\n%v\ngot\n%v", tc.out, tc.rooms)
			}
		})
	}
}
