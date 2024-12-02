package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/ronakk4/url-shortner/helpers"
)





type request struct{
URL string `json:"url`
CustomShort string  `json:"short"`
Expiry time.Duration `json:"expiry`
}

type response struct{
	URL string `json:"url`
CustomShort string  `json:"short"`
Expiry time.Duration `json:"expiry"`
XRateRemaining int `json:"rate_limit"`
XRateLimitReset time.Duration `json:"rate_limit_reset"`

}

func ShortenURL( c * fiber.Ctx)error{
	body:=new(request)
	if err:=c.BodyParser(&body);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse request"})
	}
	if ! govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid url"})

	}
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"InvalidUrl"})
	}
	
body.URL=helpers.EnforceHTTP(body.URL)
	return nil
}

func ResolveURL( c * fiber.Ctx)error{


	return nil
}
