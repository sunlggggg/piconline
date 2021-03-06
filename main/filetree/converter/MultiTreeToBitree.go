package converter

import "github.com/sunlggggg/piconline/main/entity"

//树转二叉树
//
//       1
//    /  \  \
//    2   4   5
//  /
// 3
//     ||
//     ||
//    \  /
//     \/
//
//       1
//    /  \  \
//    2 - 4 - 5
//  /
// 3
//
//     ||
//     ||
//    \  /
//     \/
//
//       1
//    /
//    2 - 4 - 5
//  /
// 3
//
//     ||
//     ||
//    \  /
//     \/
//
//     1
//    /
//    2
//  /   \
// 3     4
//        \
//         5
func MutliTreeConvertBinaryTree(root *entity.MutliTree) *entity.FileTreeNode{
	mtNodes :=make([]*entity.MutliTree,1)
	mtNodes[0] = root
	newRoot := mutliTreeConvertBinaryTree(mtNodes)
	return newRoot
}

func mutliTreeConvertBinaryTree( nodes []*entity.MutliTree) *entity.FileTreeNode {
	var btRoot *entity.FileTreeNode = nil
	// 当前只有一个节点
	if len(nodes) > 0 {
		btRoot = new(entity.FileTreeNode)
		btRoot.IsFile = nodes[0].IsFile
		btRoot.Name = nodes[0].Name
		// 左子树继续递归
		btRoot.LeftChild = mutliTreeConvertBinaryTree(nodes[0].ChildNodes)
		if len(nodes) > 1 {
			// 删除第一个
			nodes = append(nodes[0:0], nodes[1:]...)
			btRoot.RightChild = mutliTreeConvertBinaryTree(nodes)
		}
	}
	return btRoot
}
