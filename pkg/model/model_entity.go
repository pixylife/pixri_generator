package model


type Entity struct {
	Model
	Name          string      `json:"name"`
	Description   string      ` json:"description"`
	Application   Application `json:"-"`
	ApplicationID int         `json:"application_id"`
}



type EntityData struct {
	ID        int `json:"id"`
	FieldCount int `json:"field_count"`
}
