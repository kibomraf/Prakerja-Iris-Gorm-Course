package students

type StudentsFormatter struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Addrees     string `json:"address"`
	NoHandphone string `json:"no_handphone"`
	Token       string `json:"token"`
}

func FormatStudents(f Student, token string) StudentsFormatter {
	formatter := StudentsFormatter{
		Name:        f.Name,
		Email:       f.Email,
		Password:    f.Password_hash,
		Addrees:     f.Address,
		NoHandphone: f.NoHandphone,
		Token:       token,
	}
	return formatter

}
