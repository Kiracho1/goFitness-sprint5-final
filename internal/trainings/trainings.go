package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// Разделили строку на слайс строк
	data := strings.Split(datastring, ",")

	if len(data) != 3 {
		return fmt.Errorf("длина слайса строк должна быть равна 3")
	}
	// Проверяем на ошибки преобразования строки в число
	if t.Steps, err = strconv.Atoi(data[0]); err != nil {
		return err
	}
	// Проверяем на отрицательное число
	if t.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше 0")
	}

	// Сохраняем значение типа тренировки в поле структуры
	t.TrainingType = data[1]

	// Проверяем на ошибки формата продолжительности времени
	if t.Duration, err = time.ParseDuration(data[2]); err != nil {
		return err
	}
	if t.Duration <= 0 {
		return fmt.Errorf("длительность тренировки должна быть больше 0")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// Создаем результирующую строку с выводными данными
	var result string

	switch t.TrainingType {
	case "Бег":
		// Рассчитываем дистанцию
		distance := spentenergy.Distance(t.Steps, t.Height)
		// Рассчитываем скорость
		meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		runSptCal, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Тип тренировки: %s\n" +
			"Длительность: %.2f ч.\n" +
			"Дистанция: %.2f км.\n" +
			"Скорость: %.2f км/ч\n" +
			"Сожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance,
			meanSpeed, runSptCal)
	case "Ходьба":
		// Рассчитываем дистанцию
		distance := spentenergy.Distance(t.Steps, t.Height)
		// Рассчитываем скорость
		meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		walkSptCal, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Тип тренировки: %s\n" +
			"Длительность: %.2f ч.\n" +
			"Дистанция: %.2f км.\n" +
			"Скорость: %.2f км/ч\n" +
			"Сожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance,
			meanSpeed, walkSptCal)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return result, nil
}
