package user

import "MovieAPI/internal/core"

type (
	Login struct {
		core.BaseModel `bson:",inline"`
		UserName       string `json:"userName" bson:"UserName"`
		Password       string `json:"password" bson:"Password"`
	}

	User struct {
		core.BaseModel `bson:",inline"`
		UserName       string `json:"userName" bson:"UserName"`
		Password       string `json:"password" bson:"Password"`
		Name           string `json:"name" bson:"Name"`
		LastName       string `json:"lastName" bson:"LastName"`
	}
)
