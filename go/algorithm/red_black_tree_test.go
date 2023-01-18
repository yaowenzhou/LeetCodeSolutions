package algorithm

import (
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func TestRedBlackTree(t *testing.T) {
	rbt := redblacktree.NewWithIntComparator()
	rbt.Put(1, 1)
}
