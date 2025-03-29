package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

func (u *User) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		ID:    u.ID.Hex(),
		Email: u.Email,
	})
}

type LoginResponse struct {
	AccessToken string             `bson:"access_token"`
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role        string             `bson:"role"`
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdfdjbjsdjdsj
// jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjdfnsjkjskj
// sdjfjksdfdjhbdjshdjfhhj
//jkjhfjdskjksdfjk
// fbjkgkdfbjdf jkdfjgjksdjbfjksdsjdhjsdjhjksdjk
