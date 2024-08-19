# Trees

Trees (a subset of graphs) can be implemented in different ways, these include:
- A 2D array.
- A map.
- Object/Reference/Pointers

Given a tree
```text
        A
    +---+---+
    |       |
    B       C
+---+
|
D
```

**2D Array**
A 2D array defines as "parent" as the Row and "children" as the Column. The
index of each row and column denotes a node, where `i = j = same node`.

```text
    A   B   C   D
A | 0 | 1 | 1 | 0 |
B | 0 | 0 | 0 | 1 |
C | 0 | 0 | 0 | 0 |
D | 0 | 0 | 0 | 0 |
```

Here we see that row `0`, denotes node `A` has children columns indices `1, 2`
which denotes children `B, C`. Row `1` denotes node `B` has children columns
indices `3` which denotes children `D`. Finally, row indices `2, 3` denotes
nodes `C, D`, however these don't have any children nodes.

If we however, look at the column and column indices we are able to find the
children's parents, using the exact same look-up process with the transposed
matrix.

**Map**
A map is used where the key denotes the parent and the value is an array of
keys which maps to the children nodes.

```text
    A: [B, C]
    B: [D]
```

Here we see that `A` has children `B, C` and `B` has children `D`.

**Object/Reference/Pointer**
An Object/Reference/Pointer refers to the language syntax used but breaking it
down to the foundational concept, there is a structure that has a pointer to
another structure. Which is best represented as a tree.

```text
        A
    +---+---+
    |       |
    B       C
+---+
|
D
```

And represented in Go, where a node has a value and has children which is a
slice/array of pointers to other node structures.

```go
package main

type Node struct {
	Value string
	Children []*Node
}
```
