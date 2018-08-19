package entity

type Fileroot struct {
	Id uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	FileId uint64 `json:"file_id"`
	CreateTime uint64 `json:"create_time"`
}
