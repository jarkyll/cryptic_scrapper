package models

import "github.com/kamva/mgm/v3"

type TokenDetails struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AccessTokenExpiration    int64
	RefreshTokenExpiration    int64
}

type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthModel struct{}

