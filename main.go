package main

import (
	"log"
	"sqlitebenchmark/terminals"
	"sqlitebenchmark/terminals/repo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = true

	// db
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// repo
	repo, err := repo.New(db)
	if err != nil {
		log.Fatal(err)
	}

	// services
	terminals := terminals.New(repo)
	g := e.Group("/api/v1/terminals")
	terminals.AddRoutes(g)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
