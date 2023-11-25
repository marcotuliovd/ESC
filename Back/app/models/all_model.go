package models

import (
	"time"

	"github.com/google/uuid"
)

type RequestServiceReport struct {
	Init time.Time `json:"init"`
	Finish time.Time `json:"finish"`
}

type Space struct {
	Id         	uuid.UUID	        `db:"id" json:"id"`
	VehicleId      	uuid.UUID 		`db:"vehicle_id" json:"vehicle_id"`
	Type   	string 	`db:"type" json:"type"`
	Created_at time.Time `db:"created_at" json:"created_at"`
}

type Clients struct {
	Id         	uuid.UUID	        `db:"id" json:"id"`
	Username 	string 		`db:"username" json:"username"`
	Name		string		`db:"name" json:"name"`
	Phone		string		`db:"phone" json:"phone"`
}

type Vehicle struct {
	Id          	uuid.UUID	        `db:"id" json:"id"`
    Owner string    	`db:"owner" json:"owner" validate:"required,lte=255"`
    Title string `db:"title" json:"title" validate:"required,lte=255"`
    Type   	string 	`db:"type" json:"type"`
    Plate string `db:"plate" json:"plate" validate:"required,lte=255"`
}


type History struct {
	Id          	uuid.UUID	        `db:"id" json:"id"`
    Amount float32 `db:"amount" json:"amount"`
    VehicleId uuid.UUID `db:"vehicle_id" json:"vehicle_id"`
    Entry time.Time `db:"entry" json:"entry"`
    Exit time.Time `db:"exit" json:"exit"`
    Type string `db:"type" json:"type"`
}

type OutVehicle struct {
	VehicleId uuid.UUID `db:"vehicle_id" json:"vehicle_id"`
}

type HistoryReturn struct {
	Id          	uuid.UUID	        `db:"id" json:"id"`
    Amount float32 `db:"amount" json:"amount"`
    VehicleId uuid.UUID `db:"vehicle_id" json:"vehicle_id"`
    Entry time.Time `db:"entry" json:"entry"`
    Exit time.Time `db:"exit" json:"exit"`
    Type string `db:"type" json:"type"`
	Created_at time.Time `db:"created_at" json:"created_at"`
}
