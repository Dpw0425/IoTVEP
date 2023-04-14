package entity

// 项目实体
type Item struct {
	IID       uint   `json:"iid"`
	UID       uint   `json:"uid"`
	ItemIMG   string `json:"item_img"`
	ItemName  string `json:"item_name"`
	ItemIntro string `json:"item_intro"`
}
