package filetree

import "fmt"

// 文件树(二叉树形式)
type FileTreeNode struct {
	// 当是二叉树类型时，非左即右
	IsFile     bool
	LeftChild  *FileTreeNode
	RightChild *FileTreeNode
	Name       string
}

func (f *FileTreeNode)CreateNode(name string, isFile bool) *FileTreeNode {
	f.Name = name
	f.IsFile = isFile
	return f
}

func (f FileTreeNode)BiTraverse(node *FileTreeNode) {
	if node != nil {
		fmt.Println(node.Name)
		f.BiTraverse(node.LeftChild)
		f.BiTraverse(node.RightChild)
	}
}
