package main

import (
	"complaint_portal/config"
	"complaint_portal/controller"
	custommw "complaint_portal/middleware"
	"complaint_portal/models"
	"complaint_portal/repository"
	"complaint_portal/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("db setup failed: %v", err)
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if err := db.AutoMigrate(&models.UserModel{}, &models.ComplaintModel{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	userRepo := repository.NewUserRepo(db)
	compRepo := repository.NewComplaintRepo(db)

	userUC := usecase.NewUserUsecase(userRepo)
	compUC := usecase.NewComplaintUsecase(compRepo, userRepo)

	userCtrl := controller.NewUserController(userUC)
	compCtrl := controller.NewComplaintController(compUC, userUC)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", userCtrl.Register)
	e.POST("/login", userCtrl.Login)

	e.POST("/createAdmin", userCtrl.CreateAdmin)

	api := e.Group("")
	api.Use(custommw.AuthMiddleware(userUC))

	api.POST("/submitComplaint", compCtrl.SubmitComplaint)
	api.GET("/getAllComplaintsForUser", compCtrl.GetAllComplaintsForUser)
	api.GET("/viewComplaint/:id", compCtrl.ViewComplaint)

	admin := api.Group("")
	admin.Use(custommw.AdminOnlyMiddleware(userUC))

	admin.GET("/getAllComplaintsForAdmin", compCtrl.GetAllComplaintsForAdmin)
	admin.POST("/resolveComplaint/:id", compCtrl.ResolveComplaint)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
