package model



type GenRequest struct {
	Application Application
	Entity []GenEntity
	Theme Theme
}


type GenEntity struct {
	Entity Entity
	Fields []*Field
}

