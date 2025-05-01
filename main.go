package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("how did you even get here ?")
	})

	app.Post("/register", handlers.CreateUserHandler(db))
	app.Post("/login", handlers.LoginUserHandler(db))
	app.Get("/contents", handlers.GetContentsHandler(db))
	app.Use(auth.JWTMiddleware)
	app.Put("/users/:id", handlers.UpdateUserHandler(db))
	fmt.Println("✅ ROUTE PUT /users/:id registered")
	app.Post("/reviews", handlers.CreateReviewHandler(db))
	app.Get("/reviews/:id", handlers.GetReviewByIDHandler(db))
	app.Post("/favorites", handlers.CreateFavoriteHandler(db))
	app.Get("/favorites", handlers.GetFavoritesByUserHandler(db))
	app.Delete("/favorites/:id", handlers.DeleteFavoriteHandler(db))
	app.Post("/api/watch_history", handlers.CreateWatchHistoryHandler(db))
	app.Get("/api/watch_history/:user_id", handlers.GetWatchHistoryHandler(db))

	log.Fatal(app.Listen(":8080"))
}
