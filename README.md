# Go array slice memory sharing bug

This repository demonstrates a common bug in Go related to array slices and memory sharing.  When a slice is created from an array, the slice does not create a copy of the underlying data; instead, it shares the same memory.

This means that modifications made through the slice are directly reflected in the original array. This behavior can lead to unexpected results if not properly understood.

The `bug.go` file contains code that illustrates this issue. The `bugSolution.go` file offers a solution demonstrating how to avoid the problem by creating a copy of the array slice to modify independently.

## How to reproduce
1. Clone the repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go`. Observe the output and how modifications to the slice affect the original array.
4. Run `go run bugSolution.go`. Observe the output and how modifications to the copy of the slice do not affect the original array.

## Solution
The solution involves creating a copy of the array slice before making modifications. This ensures that modifications do not affect the original array.  Refer to the `bugSolution.go` for implementation details.