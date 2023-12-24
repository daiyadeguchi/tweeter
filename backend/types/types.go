package types

type Account struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"password"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Body   string `json:"postBody"`
}
