package user

type UserFormatter struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatterUserResponse(user User, token string) UserFormatter {
	formtter := UserFormatter{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formtter
}
