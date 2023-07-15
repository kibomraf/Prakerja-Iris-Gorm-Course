package students

// input user for create account
type Input struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	NoHandphone string `json:"no_handphone" validate:"required"`
}
