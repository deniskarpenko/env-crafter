package main

import (
	"context"
	"fmt"
	"log"

	"bitbucket.org/go-scrapper/docker-compose/pkg/db"
	"bitbucket.org/go-scrapper/docker-compose/pkg/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx       context.Context
	imageRepo *db.ImageRepository
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	dbGorm, err := gorm.Open(sqlite.Open("./../images.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	a.imageRepo = db.NewImageRepository(dbGorm)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetAllImages() []models.Image {
	images, err := a.imageRepo.GetAllImages()

	if err != nil {
		log.Fatal(err)
	}

	return images
}

func (a *App) GetTagByImageId(imageId int) []models.Tag {
	tags, err := a.imageRepo.GetTagsByImageId(imageId)
	if err != nil {
		log.Fatal(err)
	}

	return tags
}

func (a *App) Build(app ApplicationConfig) {
	fmt.Println(app)
}
