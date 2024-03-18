package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1773061",
		Key:     "a19dafb420f8035c9a03",
		Secret:  "3d22ddfb904fbf253677",
		Cluster: "us2",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)
		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
