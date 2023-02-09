package api

import "github.com/gofiber/fiber/v2"

type HostDetail struct {
	Name            string `json:"name"`
	Ip              string `json:"ip"`
	Active_user     string `json:"active_user"`
	First_heartbeat string `json:"first_heartbeat"`
	Last_hartbeat   string `json:"last_hartbeat"`
}

type Host struct {
	Name HostDetail `json:"name"`
}

hostlist := []Host{}

func GetHost(c *fiber.Ctx) error {
	
	name := c.Params("name")
	if name != "" {
		return c.SendString("Hello, " + name)
	}

	print("hostlist")
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    hostlist,
	})
}

func CreateHost(c *fiber.Ctx) error {
	payload := new(HostDetail)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"data":    "",
			"message": "Cannot parse input data",
		})
	}

	return c.Status(201).JSON(payload)
}
