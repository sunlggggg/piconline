package converter

import (
	"github.com/sunlggggg/piconline/main/models"
	"testing"
)

// 二叉树转多叉树
//     1
//    /
//    2
//  /   \
// 3     4
//        \
//         5
//    ||
//    ||
//   \  /
//    \/
//
//      1
//    /  \ \
//    2  | |
//  /   \| |
// 3     4 |
//        \|
//         5
//    ||
//    ||
//   \  /
//    \/
//
//      1
//    /  \ \
//    2  | |
//  /    | |
// 3     4 |
//         |
//         5
//    ||
//    ||
//   \  /
//    \/
//
//       1
//    /  \  \
//    2   4   5
//  /
// 3
func Test(t *testing.T) {
	root := new(models.FileTreeNode).CreateNode("1", false)
	root.LeftChild = new(models.FileTreeNode).CreateNode("2", false)
	root.LeftChild.LeftChild = new(models.FileTreeNode).CreateNode("3", true)
	root.LeftChild.RightChild = new(models.FileTreeNode).CreateNode("4", true)
	root.LeftChild.RightChild.RightChild = new(models.FileTreeNode).CreateNode("5", false)
	//root.LeftChild.RightChild.RightChild.RightChild = new(models.FileTreeNode).CreateNode("6", true)
	//root.LeftChild.RightChild.RightChild.LeftChild = new(models.FileTreeNode).CreateNode("7", true)
	mt := BinaryTreeConvertMutliTree(root)
	root = new(models.FileTreeNode)
	root.BiTraverse(root)
	root = MutliTreeConvertBinaryTree(mt)
	root.BiTraverse(root)
}
