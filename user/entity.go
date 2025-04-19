package user

type User struct {
	ID       uint
	Serial   string
	Username string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}
