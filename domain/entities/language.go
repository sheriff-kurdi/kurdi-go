package entities

type Language struct {
	Code string `json:"code" gorm:"primaryKey"`
	Name string `json:"name"`
	Entity
}
