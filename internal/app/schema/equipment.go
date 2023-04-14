package schema

// 本项目此阶段暂时无需使用此结构体传输数据
//type Equipment struct {
//	EID        uint        `json:"eid"`
//	EquipIMG   string      `json:"equip_img"`
//	EquipName  string      `json:"equip_name"`
//	EquipClass string      `json:"equip_class"`
//	EquipType  string      `json:"equip_type"`
//	EquipIntro string      `json:"equip_intro"`
//	EquipData  []EquipData `json:"equip_data"`
//}
//
//type EquipData struct {
//	AID   uint   `json:"aid"`
//	AName string `json:"a_name"`
//	AData string `json:"a_data"`
//}

type EquipmentInfo struct {
	EID       uint   `json:"eid" gorm:"primaryKey" binding:"omitempty"`
	EquipIMG  string `json:"equip_img" binding:"omitempty"`
	EquipName string `json:"equip_name" binding:"omitempty"`
}
