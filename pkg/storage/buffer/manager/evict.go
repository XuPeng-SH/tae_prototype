package manager

import (
	"tae/pkg/common/types"
	blkif "tae/pkg/storage/buffer/block/iface"
)

func NewSimpleEvictHolder(capacity types.IDX_T) IEvictHolder {
	holder := &SimpleEvictHolder{
		Queue: make(chan *EvictNode, capacity),
	}
	return holder
}

func (holder *SimpleEvictHolder) Enqueue(node *EvictNode) {
	holder.Queue <- node
}

func (holder *SimpleEvictHolder) Dequeue() *EvictNode {
	select {
	case node := <-holder.Queue:
		return node
	default:
		return nil
	}
}

func (node *EvictNode) Unloadable(blk blkif.IBlockHandle) bool {
	if node.Block != blk {
		panic("Logic error")
	}
	return blk.Iteration() == node.Iter
}
