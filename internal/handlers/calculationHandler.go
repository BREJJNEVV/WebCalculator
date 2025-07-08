package handlers

import (
	"CalculatorAppBackend/internal/calculationService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

// Обработчик для получения истории вычислений
func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest
	// расшифровка - декодировка того, что передали
	if err := c.Bind(&req); err != nil {
		// Проблема с данными, мы не можем их декодировать
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	calc, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		//что-то с запросом
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculations"})
	}
	return c.JSON(http.StatusCreated, calc)
}

func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculationService.CalculationRequest
	// расшифровка - декодировка того, что передали
	if err := c.Bind(&req); err != nil {
		// Проблема с данными, мы не можем их декодировать
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	updateCalc, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		// Проблема с выражением - bad request
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	return c.JSON(http.StatusOK, updateCalc)
}

func (h *CalculationHandler) DeleteCalcations(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculations"})
	}

	return c.NoContent(http.StatusNoContent)
}
