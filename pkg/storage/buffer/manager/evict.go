package manager

import (
	"fmt"
	"tae/pkg/common/types"
	blkif "tae/pkg/storage/buffer/block/iface"

	log "github.com/sirupsen/logrus"
)

func NewSimpleEvictHolder(capacity types.IDX_T) IEvictHolder {
	holder := &SimpleEvictHolder{
		Queue: make(chan *EvictNode, capacity),
	}
	return holder
}

func (holder *SimpleEvictHolder) Enqueue(node *EvictNode) {
	log.Infof("Equeue evict blk %v", node.Block.GetID())
	holder.Queue <- node
}

func (holder *SimpleEvictHolder) Dequeue() *EvictNode {
	select {
	case node := <-holder.Queue:
		log.Infof("Dequeue evict blk %v", node.Block.GetID())
		return node
	default:
		log.Info("Dequeue empty evict blk")
		return nil
	}
}

func (node *EvictNode) String() string {
	return fmt.Sprintf("EvictNode(%v, %d)", node.Block, node.Iter)
}

func (node *EvictNode) Unloadable(blk blkif.IBlockHandle) bool {
	if node.Block != blk {
		panic("Logic error")
	}
	return blk.Iteration() == node.Iter
}
