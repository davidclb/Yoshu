package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/davidclb/Yoshu/account/handlers"
	repository "github.com/davidclb/Yoshu/account/repository"
	"github.com/davidclb/Yoshu/account/service"
	"github.com/davidclb/Yoshu/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func main(){
	


////////////////////////////////////////// Data source//////////////////////////////////////////////////////////////////////////////////


	log.Printf("Initializing data sources\n")
	// load env variables - we could pass these in,
	// but this is sort of just a top-level (main package)
	// helper function, so I'll just read them in here


	log.Printf("Connecting to Postgresql\n")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost user=root password=root dbname=yoshu port=5432 sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	conn.Ping(ctx)
	conn.Exec(ctx, ";")
	fmt.Println("Le ping est passé")


	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	queries := sqlc.New(conn)
	


////////////////////////////////////////// Dependency injection //////////////////////////////////////////////////////////////////////////////////

// Repository layer 

   userRepository := repository.NewUserRepository(&repository.UserRepo{
		Querie : queries,
	})


// Service layer 

   userService := service.NewUserService(&service.USConfig{
		Repository:  userRepository,
	}) 

// Handler layer

   app := fiber.New()

   app.Use(cors.New(cors.Config{
	AllowCredentials: true,
   }))

 	handlers.NewHandler(&handlers.Config{
		R: app,
		UserService: userService,
	})
 
////////////////////////////////////////// Server initialisation//////////////////////////////////////////////////////////////////////////////////

	log.Println("Starting server...")

    app.Listen(":3001")

	log.Println("Server started")


}