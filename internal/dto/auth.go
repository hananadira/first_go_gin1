package dto 

type RegisterRequest struct {
	Name 		string `json:"name" binding:"required,min=2,max=100"`
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email 		string `json:"email" binding:"required,email"`
	Password 	string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID 		int64  `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
}

type LoginResponse struct {
	Token  string 		`json:"token"`
	User   UserResponse `json:"user"`
}