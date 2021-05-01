package segment

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	start1 := types.IDX_0
	count := types.IDX_T(1024)
	start2 := start1 + count
	seg1 := NewSegment(start1, count)
	seg2 := NewSegment(start2, count)
	tree := NewSegmentTree()
	assert.Equal(t, tree.Depth(), types.IDX_0)
	t.Log(tree.String())
	tree.Append(seg1)
	assert.Equal(t, tree.Depth(), types.IDX_1)
	t.Log(tree.String())
	tree.Append(seg2)
	assert.Equal(t, tree.Depth(), types.IDX_2)
	t.Log(tree.String())
	t.Log(tree.ToString(types.IDX_1))

	assert.Equal(t, tree.GetRoot(), seg1)
	assert.Equal(t, tree.GetTail(), seg2)

	seg3 := tree.WhichSeg(count)
	assert.Equal(t, seg3, seg2)
	seg4 := tree.WhichSeg(start1)
	assert.Equal(t, seg4, seg1)
	seg5 := tree.WhichSeg(start1 + 2*count)
	assert.Equal(t, seg5, nil)
}
