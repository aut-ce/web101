package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"

	"context"
	"fmt"
	"log"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	e := echo.New()

	// try to allow cors from anywhere
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The hallows.ir server is up")
	})

	e.GET("/chart", func(c echo.Context) error {
		return c.File("./data/chart.json")
	})
	e.GET("/occasion", func(c echo.Context) error {
		return c.File("./data/occasion.json")
	})
	e.GET("/house/:id", func(c echo.Context) error {
		return c.File("./data/house.json")
	})
	e.GET("/mags", func(c echo.Context) error {
		return c.File("./data/mags.json")
	})

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://local:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	e.Logger.Fatal(e.Start(":80"))
}
