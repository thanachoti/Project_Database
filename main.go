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
	fmt.Println("hello from profile branch")
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
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("how did you even get here ?")
	})
	app.Post("/register", handlers.CreateUserHandler(db))
	app.Post("/login", handlers.LoginUserHandler(db))
	app.Get("/contents", handlers.GetContentsHandler(db))

	app.Use(auth.JWTMiddleware)
	app.Put("/users/profile_picture", handlers.UploadProfilePictureHandler(db))
	app.Get("/users/:user_id", handlers.GetCurrentUserHandler(db))
	app.Put("/users/:user_id", handlers.UpdateUserProfileHandler(db))
	app.Post("/users/change-password", handlers.ChangePasswordHandler(db))
	app.Post("/reviews", handlers.CreateReviewHandler(db))
	app.Get("/reviews/:content_id", handlers.GetReviewByContentIDHandler(db))
	app.Delete("/reviews/:review_id", handlers.DeleteReviewHandler(db))
	app.Post("/favorites", handlers.CreateFavoriteHandler(db))
	app.Get("/favorites/:user_id", handlers.GetFavoritesByUserHandler(db))
	app.Delete("/favorites/:content_id", handlers.DeleteFavoriteHandler(db))
	app.Post("/watch_history", handlers.CreateWatchHistoryHandler(db))
	app.Get("/watch_history/:user_id", handlers.GetWatchHistoryHandler(db))
	app.Get("/update_subscription", handlers.UpdateSubscriptionHandler(db))

	app.Use(auth.AdminOnlyMiddleware)
	app.Get("/hi_admin", func(c *fiber.Ctx) error {
		return c.SendString("Hello admin")
	})
	app.Post("/contents", handlers.CreateContentHandler(db))
	app.Delete("/contents/:content_id", handlers.DeleteContentHandler(db))
	app.Put("/contents/:content_id", handlers.UpdateContentHandler(db))
	app.Get("/users", handlers.GetUsersHandler(db))
	app.Delete("/users/:user_id", handlers.DeleteUserHandler(db))

	log.Fatal(app.Listen(":8080"))
}
