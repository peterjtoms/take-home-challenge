## Algorithm

### Search 

I chose to use Breadth First Search to find the shallowest duplicate. I selected this algorithm because Breadth First Search will go through one level completely before starting the next level. In this scenario, it's never a good idea to advance to the next level early, because we may find a duplicate that is not the most shallow. Something like Depth First Search would not be able to terminate early like a Breadth First Search. 

I implemented this with an array-based queue to add subsequent children to the end of the array.

### Managing Duplicates

To manage the checking of duplicates I used a basic Go map. My understanding is this is the Go equivalent of a hash map or hash table. I chose this because it has the fastest time complexity for adding and retrieving values. Admittedly, Go allows for an "exists" return value when retrieving from a map, so the boolean stored as a value is never actually used.

### Retrieving Current Level

For retrieving the current level of an item in the tree, I used a really simple form of indexing. Since it's Breadth First Search, all items on the queue will either be from the current level (x) or from the next level (x+1). When I reach the end of a level, I know how many children were added to the next level, so I use that as the new number of children for the current level.

## Big O Time Complexity (runtime)

The worst-case and average time complexity of the search portion should be O(n). In some cases the Big O for a Breadth First Search is considered O(V + E) where V is the number of nodes and E is the number of edges between them. I believe the worst case for this tree should simplify to O(n) since we can only iterate through each node a single time, and each node only has one parent.

We check whether a node is a duplicate once per iteration. Retrieving from a map or inserting to a map should be an O(1) operation on average. In a worst case scenario where re-hashing happens very often, hash map operations can be O(n). Re-hashing should be fairly rare with integer-based keys, but as I'm not fully confident with the behind-the-scenes of Go maps, I believe O(n) is the true worst case.

Combining these values, the average time complexity would be O(n), while the worst-case time complexity would likely be O(n^2).

## Big O Space Complexity (space)

The space complexity of this implementation should be O(n). The space complexity for the map that I used to determine whether or not there are duplicates is at worst case O(n). I also use a queue for storing the possible next items in the tree, which can at worst case also be O(n). These are not multiplied together, so in the worst case, as well as the average case, this should still take up O(n) space.

TreeNode may be a bad data structure, as the root node contains references to the entirety of the tree. It's possible that using pointers to the children would make this more effective, though I'm fairly new to Go so I did not go in that direction. O(n) should still be the worst case.

## Assumptions

### Shallowest duplicate node

I made an assumption that the goal was to find the shallowest node that duplicates a previous node in the tree. After completing this implementation, I realized this might have been a bad assumption on my part.

For an example of the difference, if there is a tree with a root value of 20, two children both with a value of 3, and a grandchild with a value of 20, my algorithm will return "3,1". However, it's possible the intent of this assignment was to return "20, 0" because the first node in the tree that has a duplicate anywhere in the tree is the root.

I included an alternate solution (see alternate-tree-search/search2.go) for this possible case. It's not as well tested, but it performs a similar algorithm. Instead of terminating early, it waits until all nodes have been explored. It then iterates through an array of duplicates, accessing the hash map, which now stores the level in which that value was first found. The item with the minimal level will be returned with its value and level. The time and space complexities should be pretty similar.

If I made the right assumption, feel free to ignore the "alternate-tree-search" folder altogether.

### Other assumptions

I made the assumption that a child node in the tree could not have multiple parents, and that the root node has no parent node.

I also assumed that the Value for each node in the tree would be an integer. This is not necessarily true based on the README, where no value type was mentioned, although the word "ID" implies an integer or a string.

I also return a pointer as the first return value instead of returning the actual value. As mentioned in the requirements, the first element has to return as nil if no duplicates are found. Go would not allow me to return both an integer and nil for the same return value, so I return a pointer to an integer or nil.

Finally, I made an assumption that maps in Go have an average time complexity of O(1) for insertion and O(1) for retrieval.