package alternate

import (
	"testing"
)

func TestB1(t *testing.T) {
	// Test what happens if the TreeNode passed in has no children
	parent := TreeNode{100, []TreeNode{}}
	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be nil and expect level to be 0
	expectedLevel := 0

	if value != nil {
		t.Error("Invalid value expected:", nil, "actual: ", value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB2(t *testing.T) {
	// Test what happens if the TreeNode passed has children, but no duplicates
	child1 := TreeNode{1, []TreeNode{}}
	child2 := TreeNode{2, []TreeNode{}}
	parent := TreeNode{100, []TreeNode{child1, child2}}

	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be nil and expect level to be 0
	expectedLevel := 0

	if value != nil {
		t.Error("Invalid value expected:", nil, "actual: ", value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB3(t *testing.T) {
	// Test what happens if the TreeNode passed in has grandchildren, but no duplicates
	child1 := TreeNode{1, []TreeNode{}}

	grandChild1 := TreeNode{3, []TreeNode{}}
	grandChild2 := TreeNode{4, []TreeNode{}}
	child2 := TreeNode{2, []TreeNode{grandChild1, grandChild2}}

	parent := TreeNode{100, []TreeNode{child1, child2}}

	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be nil and expect level to be 0
	expectedLevel := 0

	if value != nil {
		t.Error("Invalid value expected:", nil, "actual: ", value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}
func TestB4(t *testing.T) {
	// Test what happens if the TreeNode passed has children, and an immediate duplicate
	child1 := TreeNode{3, []TreeNode{}}
	child2 := TreeNode{3, []TreeNode{}}
	parent := TreeNode{100, []TreeNode{child1, child2}}

	value, level := CheckDuplicateIDs2(parent)
	// Expect value to be 3 and expect level to be

	expectedValue := 3
	expectedLevel := 1

	if *value != expectedValue {
		t.Error("Invalid value expected:", expectedValue, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB5(t *testing.T) {
	// Test what happens if the TreeNode passed in has grandchildren, and grandchild duplicates
	child1 := TreeNode{1, []TreeNode{}}

	grandChild1 := TreeNode{2, []TreeNode{}}
	grandChild2 := TreeNode{4, []TreeNode{}}
	child2 := TreeNode{2, []TreeNode{grandChild1, grandChild2}}

	parent := TreeNode{100, []TreeNode{child1, child2}}

	value, level := CheckDuplicateIDs2(parent)
	// Expect value to be nil and expect level to be 0

	expectedValue := 2
	expectedLevel := 1

	if *value != expectedValue {
		t.Error("Invalid value expected:", expectedValue, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB6(t *testing.T) {
	// Test what happens if the TreeNode has sooner duplicates, but at a further level

	greatGrandChild1 := TreeNode{100, []TreeNode{}}
	leftGrandChild1 := TreeNode{5, []TreeNode{greatGrandChild1}}
	child1 := TreeNode{1, []TreeNode{leftGrandChild1}}

	grandChild1 := TreeNode{4, []TreeNode{}}
	grandChild2 := TreeNode{2, []TreeNode{}}
	child2 := TreeNode{2, []TreeNode{grandChild1, grandChild2}}

	parent := TreeNode{100, []TreeNode{child1, child2}}

	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be nil and expect level to be 0
	expectedValue := 100
	expectedLevel := 0

	if *value != expectedValue {
		t.Error("Invalid value expected:", expectedValue, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB7(t *testing.T) {
	// Test nothing goes wrong if the TreeNode sets children to nil
	badParent := TreeNode{100, nil}
	value, level := CheckDuplicateIDs2(badParent)

	// Expect value to be nil and expect level to be 0
	expectedLevel := 0

	if value != nil {
		t.Error("Invalid value expected:", nil, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB8(t *testing.T) {
	// Test what happens if the TreeNode has many singular children, and the last one is the duplicate
	childrenArr := []TreeNode{}
	for i := 0; i < 500; i++ {
		childrenArr = append(childrenArr, TreeNode{i, []TreeNode{}})
	}

	parent := TreeNode{499, childrenArr}

	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be 49 and expect level to be 1
	expectedValue := 499
	expectedLevel := 0

	if *value != expectedValue {
		t.Error("Invalid value expected:", expectedValue, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB9(t *testing.T) {
	// Test what happens if the TreeNode has many singular children, and no duplicates
	childrenArr := []TreeNode{}
	for i := 0; i < 500; i++ {
		grandChild1 := TreeNode{i + 500, []TreeNode{}}
		grandChild2 := TreeNode{i + 1000, []TreeNode{}}
		childrenArr = append(childrenArr, TreeNode{i, []TreeNode{grandChild1, grandChild2}})
	}

	parent := TreeNode{-1, childrenArr}

	value, level := CheckDuplicateIDs2(parent)

	expectedLevel := 0
	// Expect value to be nil and expect level to be 0

	if value != nil {
		t.Error("Invalid value expected:", nil, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}

func TestB10(t *testing.T) {
	// Test what happens if the TreeNode is one long, linear tree
	prevChild := TreeNode{50, []TreeNode{}}
	for i := 49; i > 0; i-- {
		prevChild = TreeNode{i, []TreeNode{prevChild}}
	}

	parent := TreeNode{40, []TreeNode{prevChild}}

	value, level := CheckDuplicateIDs2(parent)

	// Expect value to be nil and expect level to be 0
	expectedValue := 40
	expectedLevel := 0

	if *value != expectedValue {
		t.Error("Invalid value expected:", expectedValue, "actual: ", *value)
		t.Fail()
	}
	if level != expectedLevel {
		t.Error("Invalid level expected:", expectedLevel, "actual: ", level)
		t.Fail()
	}
}
