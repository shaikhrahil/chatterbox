package main

import (
	"chatterbox/accounts"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func main() {

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis-13055.c244.us-east-1-2.ec2.cloud.redislabs.com:13055"
	}
	redisDB := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "qHPPABRBDxqKS7JoF46scyM7pSKBiNbG",
	})

	log.Println("Connecting to redis at ", redisURL)

	connected := redisDB.Set("Connected", true, 0)

	if connected.Err() != nil {
		log.Fatal("Failed to connect to redis due to following error : \n", connected.Err())
	}

	log.Println("Starting server ..")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/ws/:id", accounts.Register(redisDB))

	app.Listen(":3000")
}
