package handler

import (
	"crypto/md5"

	"github.com/gofiber/fiber/v2"
)

func UlrRoutes(app *fiber.App) {
	// app.Post()
	app.Get("/:short", func(c *fiber.Ctx) error {
		short := c.Params("short")
		hasher := md5.New()
		hasher.Write([]byte(short))

		// url := hex.EncodeToString(hasher.Sum(nil))

		// return c.JSON(fiber.Map{"shortUrl": url})

		return c.Redirect("https://www.google.com", 301)
	})
}
