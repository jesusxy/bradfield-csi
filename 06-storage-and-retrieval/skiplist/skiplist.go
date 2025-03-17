package skiplist

import (
	"bytes"
	"fmt"
	"math/rand"
)

type MemDb struct {
	sl *SkipList
}

const (
	maxLevel = 9
)

type SkipList struct {
	head  *Node
	level int // current max level of skiplist
}

type Node struct {
	key   []byte
	value []byte

	next  [maxLevel]*Node // array of pointers to the next nodes on each level
	level int             // highest level of this node
}

func NewMemDb() *MemDb {
	return &MemDb{
		sl: &SkipList{
			head: &Node{
				key:   nil,
				value: nil,
				next:  [maxLevel]*Node{},
				level: 0,
			},
		},
	}
}

func (memdb *MemDb) Put(newKey, value []byte) error {
	currNode, levelsToUpdate := memdb.search(newKey)

	// update value if the key already exists
	if currNode != nil && bytes.Equal(currNode.key, newKey) {
		currNode.value = value
		return nil
	}

	// instert new node if the key does not exist

	lvl := randomLevel()

	if lvl > memdb.sl.level {
		for i := memdb.sl.level; i < lvl; i++ {
			levelsToUpdate[i] = memdb.sl.head
		}

		memdb.sl.level = lvl
	}

	// create new node and set pointers at each level
	currNode = &Node{
		key:   newKey,
		value: value,
		next:  [maxLevel]*Node{},
		level: lvl,
	}

	for i := 0; i < lvl; i++ {
		currNode.next[i] = levelsToUpdate[i].next[i]
		levelsToUpdate[i].next[i] = currNode
	}

	return nil
}

func (memdb *MemDb) Get(targetKey []byte) (value []byte, err error) {
	currNode, _ := memdb.search(targetKey)

	if currNode != nil && bytes.Equal(currNode.key, targetKey) {
		return currNode.value, nil
	}

	return nil, fmt.Errorf("key not found in db, %v", targetKey)
}

func (memdb *MemDb) Has(targetKey []byte) (bool, error) {
	currNode, _ := memdb.search(targetKey)

	if currNode != nil && bytes.Equal(currNode.key, targetKey) {
		return true, nil
	}

	return false, nil
}

func (memdb *MemDb) Delete(targetKey []byte) error {
	currNode, levelsToUpdate := memdb.search(targetKey)

	// if node to delete is found
	if currNode != nil && bytes.Equal(currNode.key, targetKey) {
		// update pointers for all levels where node is present
		for i := 0; i < len(levelsToUpdate); i++ {
			if levelsToUpdate[i].next[i] != currNode {
				break
			}

			levelsToUpdate[i].next[i] = currNode.next[i]
		}

		// adjust skiplist level if highest level is empty
		for memdb.sl.level > 0 && memdb.sl.head.next[memdb.sl.level-1] == nil {
			memdb.sl.level--
		}

		return nil
	}

	return fmt.Errorf("given key not found in db")
}

func (memdb *MemDb) search(key []byte) (*Node, []*Node) {
	levelsToUpdate := make([]*Node, maxLevel)
	currentNode := memdb.sl.head

	// find position for the new node, start at the top level and move down
	for i := memdb.sl.level; i >= 0; i-- {
		for currentNode.next[i] != nil && bytes.Compare(currentNode.next[i].key, key) < 0 {
			currentNode = currentNode.next[i]
		}

		levelsToUpdate[i] = currentNode
	}

	return currentNode.next[0], levelsToUpdate

}

func randomLevel() int {
	lvl := 1

	for rand.Float64() < 0.5 && lvl < maxLevel-1 {
		lvl++
	}

	return lvl
}
