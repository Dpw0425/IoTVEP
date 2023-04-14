package schema

type ItemInfo struct {
	IID      uint   `json:"iid"`
	ItemIMG  string `json:"item_img"`
	ItemName string `json:"item_name"`
}

type AddItem struct {
	UID        uint   `json:"uid"`
	ItemIMG    string `json:"item_img"`
	ItemName   string `json:"item_name"`
	ItemIntro  string `json:"item_intro"`
	Equipments []int  `json:"equipments"`
}
