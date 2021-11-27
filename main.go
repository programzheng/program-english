package main

import (
	"os"

	"github.com/programzheng/program-english/dictionary"
	"github.com/programzheng/program-english/orm"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	orm.InitDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Index!")
	})

	apiGroup := app.Group("/api")
	v1Group := apiGroup.Group("/v1")
	dictionaryGroup := v1Group.Group("/dictionary")
	dictionaryGroup.Get("", dictionary.GetDictionaries)
	dictionaryGroup.Get(":id", dictionary.GetDictionary)
	dictionaryGroup.Post("", dictionary.NewDictionary)
	dictionaryGroup.Put(":id", dictionary.UpdateDictionary)
	dictionaryGroup.Delete(":id", dictionary.DeleteDictionary)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
