package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DepartmentModel struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	DepartmentName *string            `bson:"department_name" json:"department_name"`
	CreatedAT      time.Time          `bson:"created_at" json:"created_at"`
	LastUpdate     time.Time          `bson:"last_update" json:"last_update"`
}

type UpdateDepartmentModel struct {
	DepartmentName *string   `bson:"department_name" json:"department_name"`
	LastUpdate     time.Time `bson:"last_update" json:"last_update"`
}

type DepartmentCollection struct {
	Departments []DepartmentModel `json:"departments"`
}
