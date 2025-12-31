package model

type LoginRequest struct {
	Username string `json:"username" name:"username" example:"admin"`
	Password string `json:"password" name:"password" example:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
