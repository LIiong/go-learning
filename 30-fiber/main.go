package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := fiber.New()

	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(&User{"John", 20})
		// => {"name":"John", "age":20}
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hi John!",
		})
	})
	//log.Fatal(app.Listen(":3000"))
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	fmt.Printf("Shutdown Server ... \r\n")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	return c.SendString(c.Path())
}
