package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/ronakk4/url-shortner/database"
)

func ResolveURL(c *fiber.Ctx)error{

	url:=c.Params("url")

	r:=database.CreateClient(0)
	defer r.Close()

	value,err:=r.Get(database.Ctx,url).Result()
	if err ==redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error":"short not found"})
	}else if err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":"Cannot connect db"})
	}
	rInr:=database.CreateClient(1)
	defer rInr.Close()
	_=rInr.Incr(database.Ctx,"counter")
	return c.Redirect(value,301)
}