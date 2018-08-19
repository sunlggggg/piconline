package filecache

import "github.com/sunlggggg/piconline/main/entity"


// cache
var BitTeeMap map[uint64]*entity.FileTreeNode
var MultiTeeMap map[uint64]*entity.MutliTree

func bitTreeIsExisted(userID uint64){
	if BitTeeMap[userID] == nil {
		// 从数据库中加载
	}
}

// 从数据库中加载二叉乎
func loadFormDb(userID uint64 ){
	// 找到用户根目录

}
