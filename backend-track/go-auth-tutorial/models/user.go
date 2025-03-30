package models

import "encoding/json"

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // Allow deserialization
}

// Custom MarshalJSON to exclude the Password field in responses
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    u.ID,
		Email: u.Email,
	})
}

//jkdfbjkgkdfbjdf dfjkkjsdfjksjkdkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// dbjfsfdkjjk
