package models

import (
	"time"
)

// Book struct to describe book object.
type Vaga struct {
	Id         	int	        `db:"id" json:"id"`
	Porte      	string 		`db:"porte" json:"porte"`
	Ocupacao   	time.Time 	`db:"ocupacao" json:"ocupacao"`
	Dono     	string    	`db:"dono" json:"dono" validate:"required,lte=255"`
	Telefone 	string      `db:"telefone" json:"telefone" validate:"required,lte=15"`
	TipoDeVeiculo string 	`db:"tipo_de_veiculo" json:"tipo_de_veiculo" validate:"required,lte=255"`
	Placa 		string 		`db:"placa" json:"placa" validate:"required,lte=255"`
	Modelo 		string 		`db:"modelo" json:"modelo" validate:"required,lte=255"`
}
