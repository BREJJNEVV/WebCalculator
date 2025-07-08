package main

import (
	"CalculatorAppBackend/internal/calculationService"
	"CalculatorAppBackend/internal/db"
	"CalculatorAppBackend/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connected to DB: %v", err)
	}
	e := echo.New() // Создаем экземпляр Echo фреймворка

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHanders := handlers.NewCalculationHandler(calcService)

	// Подключаем middleware
	e.Use(middleware.CORS())   // Разрешаем кросс-доменные запросы
	e.Use(middleware.Logger()) // Логируем запросы

	// Регистрируем маршрут GET /calculations
	e.GET("/calculations", calcHanders.GetCalculations)
	e.POST("/calculations", calcHanders.PostCalculations)
	e.PATCH("/calculations/:id", calcHanders.PatchCalculations)
	e.DELETE("/calculations/:id", calcHanders.DeleteCalcations)
	// Запускаем сервер на порту 8083
	e.Start("localhost:8083")

}
