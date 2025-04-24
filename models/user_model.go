package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id           bson.ObjectID `bson:"_id" json:"id"`
	Username     *string       `bson:"username" json:"username" validate:"required"`
	Password     *string       `bson:"password" json:"-" validate:"required"`
	Email        *string       `bson:"email" json:"email" validate:"required"`
	ProfileID    *string       `bson:"profile_id" json:"profile_id"`
	DepartmentID *string       `bson:"department_id" json:"department_id"`
	CreatedAT    time.Time     `bson:"created_at" json:"created_at"`
	LastUpdate   time.Time     `bson:"last_update" json:"last_update"`
}

// type UpdateUserModel struct {
// 	Username     *string   `bson:"username" json:"username"`
// 	Password     *string   `bson:"password" json:"-"`
// 	Email        *string   `bson:"email" json:"email"`
// 	ProfileID    *string   `bson:"profile_id" json:"profile_id"`
// 	DepartmentID *string   `bson:"department_id" json:"department_id"`
// 	LastUpdate   time.Time `bson:"last_update" json:"last_update"`
// }

// type UserCollection struct {
// 	Users []UserModel `json:"users"`
// }
