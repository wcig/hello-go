package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	Epoch = 1577808000000 // 2020-01-01 00:00:00

	SeqBits       = 12 // 序列号占用位数
	NodeBits      = 10 // 工作机器标识占用位数
	TimestampBits = 41 // 时间戳占用位数

	MaxSeqNum  = -1 ^ (-1 << SeqBits) // 序列号最大数
	MaxNodeNum = -1 ^ (-1 << SeqBits) // 工作机器最大数

	NodeLeaf      = SeqBits            // 工作机器偏移量
	TimestampLeaf = SeqBits + NodeBits // 时间戳偏移量
)

type Node struct {
	mu   sync.Mutex
	seq  int64
	node int64
	time int64
}

func NewNode(node int64) (*Node, error) {
	n := Node{
		seq:  0,
		node: node,
		time: -1,
	}

	if n.node < 0 || n.node > MaxNodeNum {
		return nil, errors.New("Node number must be between 0 and " + strconv.FormatInt(MaxNodeNum, 10))
	}

	return &n, nil
}

func (n *Node) GenerateId() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if now == n.time {
		n.seq = (n.seq + 1) & MaxSeqNum
		if n.seq == 0 {
			for now <= n.time {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		n.seq = 0
	}

	n.time = now
	id := ((now - Epoch) << TimestampLeaf) | (n.node << NodeLeaf) | n.seq
	return id
}
