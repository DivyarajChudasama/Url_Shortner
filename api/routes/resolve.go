package routes

import (
	"github.com/DivyarajChudasama/Url_Shortner/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("urls")
	r := database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}else if err !=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot connect to DB",
		})
	}
	rInr:=database.CreateClient(1)
	defer rInr.Close()

	_=rInr.Incr(database.Ctx,"counter")
	return c.Redirect(value,301)
}
