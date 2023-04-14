package entity

// 项目与传感器设备关联实体
type ItemEquipment struct {
	ID  uint `json:"id" gorm:"primaryKey" binding:"omitempty"`
	IID uint `json:"iid" binding:"omitempty"` // 项目 id
	EID uint `json:"eid" binding:"omitempty"` // 设备 id
}
