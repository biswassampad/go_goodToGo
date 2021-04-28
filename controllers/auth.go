package controller

import (
	"gofiber_backend/models"
	"gofiber_backend/repository"
	"gofiber_backend/security"
	util "gofiber_backend/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/asaskevich/govalidator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthController interface{
	SignUp(ctx *fiber.Ctx) error
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
	if !govalidator.IsEmail(input.Email){
		return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrInvalidEmail))
	}
	exists,err := c.userRepo.GetByEmail(input.Email)
	if err == mgo.ErrNotFound{
		if strings.TrimSpace(input.Password) =="" {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(util.ErrEmptyPassword))
		}

		input.Password,err = security.EncryptPassword(input.Password)
		if err != nil{
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}

		input.CreatedAt = time.Now()
		input.UpdatedAt = input.CreatedAt
		input.Id = bson.NewObjectId()
		err = c.userRepo.Save(&input)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
		}
		return ctx.Status(http.StatusCreated).JSON(input)
	}
	if exists != nil{
		err = util.ErrEmailAlreadyExists
	}

	return ctx.Status(http.StatusBadRequest).JSON(util.NewJError(err))
}