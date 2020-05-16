package model


type Field struct {
	Model
	Name string ` json:"name"`
	Type string `json:"type"`
	UIName string `json:"ui_name"`
	Entity   Entity `json:"-"`
	EntityID int   `json:"entity_id"`
}




