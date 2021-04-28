package main

import (
	"gofiber_backend/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()

	if err != nil{
		log.Panicln(err)
	}
}


func main(){
	conn := db.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello " + ctx.Params("name"))
	})
	log.Fatal(app.Listen(":8080"))
}