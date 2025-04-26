package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khemingkapat/been_chillin/auth"
	"github.com/khemingkapat/been_chillin/handlers"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5430
	user     = "admin"
	password = "admin"
	dbname   = "been_chillin"
)

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("Hello from march")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("fiber example for dorm management")
	})

	app.Post("/register", handlers.CreateUserHandler(db))
	app.Post("/login", handlers.LoginUserHandler(db))
	app.Get("/contents", handlers.GetContentsHandler(db))
	app.Post("/profile", handlers.CreateProfileHandler(db))
	app.Use(auth.JWTMiddleware)
	app.Post("/reviews", handlers.CreateReviewHandler(db))
	app.Get("/reviews", handlers.GetAllReviewsHandler(db))
	app.Get("/reviews/:id", handlers.GetReviewByIDHandler(db))
	app.Post("/favorites", handlers.CreateFavoriteHandler(db))
	app.Get("/favorites", handlers.GetAllFavoritesHandler(db))
	app.Delete("/favorites/:id", handlers.DeleteFavoriteHandler(db))

	log.Fatal(app.Listen(":8080"))
}
