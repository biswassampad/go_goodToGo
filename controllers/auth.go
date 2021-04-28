package controller

import (
	"gofiber_backend/models"
	"gofiber_backend/repository"
	util "gofiber_backend/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface{
	
}

type authController struct{
	userRepo repository.UsersRepository
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController{
	return &authController{usersRepo}
}


func (c *authController) SignUp(ctx *fiber.Ctx) error {
	var input models.User
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(util.NewJError(err))
	}
	input.Email = util.NormalizeEmail(input.Email)
	if !govalidator.IsEmail(input.Email)
}