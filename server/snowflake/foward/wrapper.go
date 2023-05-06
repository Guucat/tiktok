package foward

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

type Wrapper struct {
	lastTamp int64
	timeId   int64
	workId   int64
	node     *snowflake.Node
	mu       sync.Mutex
}

// 用两个位表示时间回退
func (w *Wrapper) getNodeId() int64 {
	w.timeId = (w.timeId + 1) % 4
	return w.workId<<2 + w.timeId
}

func (w *Wrapper) Id() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.lastTamp == 0 {
		w.lastTamp = w.node.Generate().Int64()
		return w.lastTamp
	}
	id := w.node.Generate().Int64()
	if id <= w.lastTamp {
		newNode, _ := snowflake.NewNode(w.getNodeId())
		w.node = newNode
	}
	w.lastTamp = w.node.Generate().Int64()
	return w.lastTamp
}

func NewWrapper(workId int64) *Wrapper {
	if workId >= 256 {
		panic("nodeId must be less than 512")
	}
	node, _ := snowflake.NewNode(workId << 2)
	return &Wrapper{
		workId: workId,
		node:   node,
	}
}
