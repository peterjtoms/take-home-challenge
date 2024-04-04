package alternate

type TreeNode struct {
	Value    int
	Children []TreeNode
}

func CheckDuplicateIDs2(tree TreeNode) (*int, int) {
	numChildren := len(tree.Children)

	// If there is only one element, there is no duplicate
	if numChildren == 0 {
		return nil, 0
	}

	existenceMap := map[int]int{}

	// The tree level is indexed at 0, so 0 is the first level
	currLevel := 0
	// Set this to 1, as it's only the root node
	numLevelRemaining := 1
	numNextLevel := 0

	treeQueue := []TreeNode{tree}
	dupeArray := []int{}

	for len(treeQueue) > 0 {
		currNode := treeQueue[0]

		_, exists := existenceMap[currNode.Value]
		if exists {
			dupeArray = append(dupeArray, currNode.Value)
		} else {
			existenceMap[currNode.Value] = currLevel
		}

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

	if len(dupeArray) == 0 {
		return nil, 0
	}

	// There are several dupes, find the one with the lowest level in the tree
	minLevel := currLevel
	var minItem int

	for i := range dupeArray {
		integer := dupeArray[i]

		val := existenceMap[integer]

		if val < minLevel {
			minItem = integer
			minLevel = val
		}
	}
	return &minItem, minLevel
}
