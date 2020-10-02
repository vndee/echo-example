package models

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}