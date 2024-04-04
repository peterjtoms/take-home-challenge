package main

type TreeNode struct {
	Value    int
	Children []TreeNode
}

func CheckDuplicateIDs(tree TreeNode) (*int, int) {
	numChildren := len(tree.Children)

	// If there is only one element, there is no duplicate
	if numChildren == 0 {
		return nil, 0
	}

	existenceMap := map[int]bool{}

	// The tree level is indexed at 0, so 0 is the first level
	currLevel := 0
	// Set this to 1, as it's only the root node
	numLevelRemaining := 1
	numNextLevel := 0

	treeQueue := []TreeNode{tree}

	for len(treeQueue) > 0 {
		currNode := treeQueue[0]

		_, exists := existenceMap[currNode.Value]
		if exists {
			return &currNode.Value, currLevel
		}

		existenceMap[currNode.Value] = true

		// Perform bread-first-search while attempting to find a duplicate
		if len(currNode.Children) > 0 {
			for j := 0; j < len(currNode.Children); j++ {
				treeQueue = append(treeQueue, currNode.Children[j])
				numNextLevel++
			}
		}

		// If the are no remaining nodes from this level in the queue, increase the level
		numLevelRemaining--
		if numLevelRemaining == 0 {
			currLevel++

			// Set the remaining nodes in the new level to how many nodes were added
			numLevelRemaining = numNextLevel
			numNextLevel = 0
		}

		treeQueue = treeQueue[1:]
	}

	// No duplicates are found
	return nil, 0
}
