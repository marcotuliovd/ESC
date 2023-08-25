package models

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Vaga struct {
	ID         	uuid.UUID 	`db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  	time.Time 	`db:"created_at" json:"created_at"`
	UpdatedAt  	time.Time 	`db:"updated_at" json:"updated_at"`
	Porte      	string 		`db:"porte" json:"porte"`
	Ocupacao   	time.Time 	`db:"ocupacao" json:"ocupacao"`
	Dono     	string    	`db:"dono" json:"dono" validate:"required,lte=255"`
	Telefone 	string      `db:"telefone" json:"telefone" validate:"required,lte=15"`
	TipoVeiculo string 		`db:"tipo_veiculo" json:"tipo_veiculo" validate:"required,lte=255"`
	Placa 		string 		`db:"placa" json:"placa" validate:"required,lte=255"`
	Modelo 		string 		`db:"modelo" json:"modelo" validate:"required,lte=255"`
}
