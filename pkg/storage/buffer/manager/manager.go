package manager

import (
	log "github.com/sirupsen/logrus"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	blk "tae/pkg/storage/buffer/block"
	blkif "tae/pkg/storage/buffer/block/iface"
	mgrif "tae/pkg/storage/buffer/manager/iface"
	"tae/pkg/storage/layout"
)

var (
	_ mgrif.IBufferManager = (*BufferManager)(nil)
)

func NewBufferManager(capacity types.IDX_T) mgrif.IBufferManager {

	mgr := &BufferManager{
		Pool:        buf.NewSimpleMemoryPool(capacity),
		TransientID: layout.MIN_TRANSIENT_BLOCK_ID,
		Blocks:      make(map[layout.BlockId]blkif.IBlockHandle),
		EvictHolder: NewSimpleEvictHolder(SIMPLE_EVICT_HOLDER_CAPACITY),
	}

	return mgr
}

func (mgr *BufferManager) GetPool() buf.IMemoryPool {
	return mgr.Pool
}

func (mgr *BufferManager) RegisterBlock(blk_id layout.BlockId) blkif.IBlockHandle {
	mgr.Lock()
	defer mgr.Unlock()

	handle, ok := mgr.Blocks[blk_id]
	if ok {
		if !handle.IsClosed() {
			return handle
		}
	}
	ctx := blk.BlockHandleCtx{
		ID:      blk_id,
		Manager: mgr,
	}
	handle = blk.NewBlockHandle(&ctx)
	mgr.Blocks[blk_id] = handle
	return handle
}

func (mgr *BufferManager) GetUsageSize() types.IDX_T {
	return mgr.Pool.GetUsageSize()
}

func (mgr *BufferManager) GetCapacity() types.IDX_T {
	return mgr.Pool.GetCapacity()
}

// Temp only can SetCapacity with larger size
func (mgr *BufferManager) SetCapacity(capacity types.IDX_T) error {
	mgr.Lock()
	defer mgr.Unlock()
	// if !mgr.makeSpace(0, capacity) {
	// 	panic(fmt.Sprintf("Cannot makeSpace(%d,%d)", 0, capacity))
	// }
	// types.AtomicStore(&(mgr.Capacity), capacity)
	return mgr.Pool.SetCapacity(capacity)
}

func (mgr *BufferManager) UnregisterBlock(blk_id layout.BlockId, can_destroy bool) {
	if blk_id.IsTransientBlock() {
		// PXU TODO
		return
	}
	mgr.Lock()
	defer mgr.Unlock()
	delete(mgr.Blocks, blk_id)
}

func (mgr *BufferManager) Unpin(handle blkif.IBlockHandle) {
	handle.Lock()
	defer handle.Unlock()
	if !handle.UnRef() {
		panic("logic error")
	}
	if !handle.HasRef() {
		evict_node := &EvictNode{Block: handle, Iter: handle.IncIteration()}
		mgr.EvictHolder.Enqueue(evict_node)
	}
}

func (mgr *BufferManager) makePoolNode(capacity types.IDX_T) *buf.PoolNode {
	node := mgr.Pool.Get(capacity, false)
	if node != nil {
		return node
	}
	for node == nil {
		// log.Printf("makePoolNode capacity %d now %d", capacity, mgr.GetUsageSize())
		evict_node := mgr.EvictHolder.Dequeue()
		// log.Infof("Evict blk %s", evict_node.String())
		if evict_node == nil {
			log.Printf("Cannot get node from queue")
			return nil
		}
		if evict_node.Block.IsClosed() {
			continue
		}

		if !evict_node.Unloadable(evict_node.Block) {
			continue
		}

		{
			evict_node.Block.Lock()
			defer evict_node.Block.Unload()
			if !evict_node.Unloadable(evict_node.Block) {
				continue
			}
			if !evict_node.Block.Unloadable() {
				continue
			}
			evict_node.Block.Unload()
		}
		node = mgr.Pool.Get(capacity, false)
	}
	return node
}

func (mgr *BufferManager) Pin(handle blkif.IBlockHandle) blkif.IBufferHandle {
	handle.Lock()
	defer handle.Unlock()
	if handle.PrepareLoad() {
		node := mgr.makePoolNode(handle.GetCapacity())
		if node == nil {
			handle.RollbackLoad()
			log.Warnf("Cannot makeSpace(%d,%d)", handle.GetCapacity(), mgr.GetCapacity())
			return nil
		}
		buf := blk.NewBlockBuffer(handle.GetID(), node)
		handle.SetBuffer(buf)
		if err := handle.CommitLoad(); err != nil {
			handle.RollbackLoad()
			panic(err.Error())
		}
	}
	handle.Ref()
	return handle.MakeHandle()
}
