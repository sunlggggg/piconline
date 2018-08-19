package converter

import (
	"github.com/sunlggggg/piconline/main/entity"
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

func BinaryTreeConvertMutliTree(root *entity.FileTreeNode) *entity.MutliTree {
	var mutliNode *entity.MutliTree
	// 非空节点才要处理
	if root != nil {
		mutliNode = new(entity.MutliTree)
		mutliNode.ChildNodes = make([]*entity.MutliTree, 0, entity.MAX)
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
