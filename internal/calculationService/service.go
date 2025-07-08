package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id string, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.CalculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	// Создаем новый экземпляр структуры, кладем туда всё, что нужно
	calc := Calculation{
		ID:         uuid.NewString(), // Генерируем случайную строку в ID
		Expression: expression,
		Result:     result,
	}
	err = s.repo.CreateCalculation(calc)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

// DeleteCalculation implements CalculationService.
func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)

}

func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calcService) GetAllCalculation() ([]Calculation, error) {
	return s.repo.GetAllCalculations()
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.CalculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc.Expression = expression
	calc.Result = result

	s.repo.SaveCalculation(calc)

	return calc, nil

}

// CalculateExpression вычисляет математическое выражение из строки
func (s *calcService) CalculateExpression(expression string) (string, error) {
	// Создаем вычислимое выражение с помощью библиотеки govaluate
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err // Возвращаем ошибку парсинга
	}
	// Вычисляем выражение без параметров (nil)
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err // Возвращаем ошибку вычисления
	}
	// Конвертируем результат в строку и возвращаем
	return fmt.Sprintf("%v", result), err

}
