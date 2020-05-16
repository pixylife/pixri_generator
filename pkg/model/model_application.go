package model


type Application struct {
	Model
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	AgeGroup AgeGroup  `json:"age-group"`
	Purpose string  `json:"purpose"`
	BaseURL string`json:"baseURL"`
	Company string `json:"company"`
	ThemeID int `json:"selected_theme"`
}

type AgeGroup struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type ApplicationData struct {
	ID        int `json:"id"`
	ThemeCount int `json:"theme_count"`
	EntityCount int `json:"entity_count"`
	PageCount int `json:"page_count"`

}
