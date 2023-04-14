package schema

type Statistics struct {
	UserCount      int64 `json:"user_count"`
	ItemCount      int64 `json:"item_count"`
	EquipmentCount int64 `json:"equipment_count"`
}
