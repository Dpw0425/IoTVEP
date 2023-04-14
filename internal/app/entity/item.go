package entity

// 项目实体
type Item struct {
	IID       uint   `json:"i_id" gorm:"primaryKey" binding:"omitempty"`
	UID       uint   `json:"uid" binding:"omitempty"`
	ItemIMG   string `json:"item_img" binding:"omitempty"`
	ItemName  string `json:"item_name" binding:"omitempty"`
	ItemIntro string `json:"item_intro" binding:"omitempty"`
}
