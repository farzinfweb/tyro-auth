package handler

import (
	"authn/auth"
	"authn/repo"
	"authn/util/tokeng"
	"authn/util/vld"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,gte=3"`
	Password string `json:"password" validate:"required,gte=5"`
}

func Login(c echo.Context, repos repo.IUserRepo) error {
	log.Println("New login request")

	loginReq := new(LoginRequest)
	if err := c.Bind(loginReq); err != nil {
		return err
	}

	validate, trans := vld.NewWithFa()
	err := validate.Struct(loginReq)
	if err != nil {
		result := ""
		for _, err := range err.(validator.ValidationErrors) {
			result += err.Translate(trans) + " "
		}
		return c.String(422, result)
	}

	_, err = auth.AttemptLogin(loginReq.Username, loginReq.Password, repos)
	if err != nil {
		return c.String(400, err.Error())
	}

	tokenSecret := c.Get("TOKEN_SECRET").(string)
	token, err := tokeng.Generate(tokenSecret)
	if err != nil {
		return c.String(400, "failed to generate token")
	}

	return c.JSON(200, map[string]string{
		"token": token,
	})
}
