package entity

// 项目与传感器设备关联实体
type ItemEquipment struct {
	ID  uint `json:"id"`
	IID uint `json:"iid"` // 项目 id
	EID uint `json:"eid"` // 设备 id
}
