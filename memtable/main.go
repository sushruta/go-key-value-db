package memtable

import "unsafe"

type Node struct {
        key uint32
        value uint32

        left *Node
        right *Node

        leftChildren int
        rightChildren int
}

type AvlTree struct {
        root *Node
        numNodes int
        totalSize uint32
}

func CreateTree(node *Node) AvlTree {
        return AvlTree{
                root: Insert(nil, node),
                numNodes: 1,
                totalSize: unsafe.Sizeof(node)
        }
}

func Insert(iteratorNode *Node, node *Node) *Node {
        if iteratorNode == nil {
                return node
        }

        if iteratorNode.key < node.key {
                iteratorNode.right = Insert(iteratorNode.right, node)
                iteratorNode.rightChildren = iteratorNode.rightChildren + 1
        } else {
            iteratorNode.left = Insert(iteratorNode.left, node)
            iteratorNode.leftChildren = iteratorNode.leftChildren + 1
        }
        diff = iteratorNode.rightChildren - iteratorNode.leftChildren
        if (diff < -1 || diff > 1) {
                this.Balance(iteratorNode)
        }
        return iteratorNode
}

// TODO: implement this method
func (this *AvlTree) Balance(node *Node) {
        return
}
