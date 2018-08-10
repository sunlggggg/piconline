package converter

import (
	"github.com/sunlggggg/piconline/main/models"
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

func BinaryTreeConvertMutliTree(root *models.FileTreeNode) *models.MutliTree {
	var mutliNode *models.MutliTree
	// 非空节点才要处理
	if root != nil {
		mutliNode = new(models.MutliTree)
		mutliNode.ChildNodes = make([]*models.MutliTree, 0, models.MAX)
		mutliNode.Name= root.Name
		mutliNode.IsFile = root.IsFile
		cur := root.LeftChild
		for i := 0; cur != nil; i++ {
			mutliNode.ChildNodes = append(mutliNode.ChildNodes, BinaryTreeConvertMutliTree(cur))
			cur = cur.RightChild
		}
	}
	return mutliNode
}
