# Arrays

#data-structure #array

> **Definition**
> An Array is a collection of items. The items could be integers, strings, other
> structs—anything really. The items are stored in neighbouring (contiguous)
> memory locations. Because they're stored together, checking through the entire
> collection of items is straightforward.

**Important to note** that an Array has a fixed length. Some languages such as Go
provide multiple types such as an array `[5]interface{}` and slices 
`[]interface{}` the array specifies the length.

### Attributes

**Capacity** is the number of elements an array can hold.

**Length** is the number of elements in the array at that time.

### Methods

**Insert**
- Insert at the end.
- Insert in the middle.
- Insert at the beginning.

> Examples
> [88 Merge Sorted Array](./88_merge_sorted_array_test.go) and
> [1089 Duplicate Zeros](./1089_duplicate_zeros_test.go)

**Delete**
- Delete at the end.

  When deleting at the end, we do not need to remove the actual value all we
  need to do is just decrement the length by one, then when a new value is
  added it will overwrite. 
  - Time complexity `O(1)`.

- Delete at the beginning.

  Is the most costly, as the length is decremented. However, from index 1 all
  the way to index `length -1` need to be shifted to the left by one. 
    - Time complexity of `O(N)`.

- Delete at in the middle.

  Similar to the previous, however, we start the shift from the index of the
  element that is to be removed. 
    - Time complexity of `O(K)` where `K` is the number of elements shifted and
      `K = N` is possible therefore `O(N)`.

> Examples 
> [27 Remove Element](./27_remove_element_test.go) and 
> [26 Remove Duplicates from Sorted Array](./26_remove_duplicates_from_sorted_array_test.go)

**Search**


**In-Place Operations**