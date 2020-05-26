package memtable

type BstNode struct {
        key string
        value uint32

        left *BstNode
        right *BstNode
}

type BinarySearchTree struct {
        root *BstNode
        numElements int
}

func BstConstructor() BinarySearchTree {
        return BinarySearchTree{nil, 0}
}

func (this *BinarySearchTree) Insert(key string, value uint32) {
        node := BstNode{key, value, nil, nil}
        if this.root == nil {
                this.root = &node
                this.numElements = 1
        } else {
                insert(this.root, &node)
                this.numElements = this.numElements + 1
        }
}

func insert(root *BstNode, node *BstNode) {
        if root.key < node.key {
                if root.right == nil {
                        root.right = node
                } else {
                        insert(root.right, node)
                }
        } else {
                if root.left == nil {
                        root.left = node
                } else {
                        insert(root.left, node)
                }
        }
}

func (this *BinarySearchTree) Inorder() []KeyValue {
        return inorder(this.root)
}

func inorder(node *BstNode) []KeyValue {
        if node == nil {
                return []KeyValue{}
        }

        left := inorder(node.left)
        right := inorder(node.right)
        tmp := append(left, KeyValue{node.key, node.value})
        elements := append(tmp, right...)
        return elements
}
