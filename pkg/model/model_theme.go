package model


type Theme struct {
	Model
	PrimaryColor     string `json:"primary_color"`
	SecondaryColor   string `json:"secondary_color"`
	PrimaryDarkColor string `json:"primary_dark_color"`
	BodyColor string `json:"body_color"`
	TextColorBody string `json:"text_color_body"`
	TextColorAppBar string `json:"text_color_appBar"`
	Application   Application `json:"-"`
	ApplicationID int `json:"application_id"`
}