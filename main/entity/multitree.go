package entity

import "fmt"

// 一个文件夹里最多100个文件和文件夹
const MAX = 100

// 多叉树结构
type MutliTree struct {
	Name  string
	IsFile bool
	ChildNodes []*MutliTree
}

func (m MutliTree)Traverse(node *MutliTree) {
	if node != nil {
		fmt.Println(node.Name)
		for i := 0; i < len(node.ChildNodes); i++ {
			m.Traverse(node.ChildNodes[i])
		}
	}
}
