package models

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel 		`bson:",inline"`
	Email 					string `json:"email" bson:"email"`
	Password 				string `json:"password" bson:"password"`
}

func (model *Book) Creating() error {
	// Call to DefaultModel Creating hook check if the email is used
	if err:=model.DefaultModel.Creating();err!=nil{
		return err
	}

	// We can check if model fields is not valid, return error to
	// cancel document insertion .
	if model.Pages < 1 {
		return errors.New("book must have at least one page")
	}

	return nil
}