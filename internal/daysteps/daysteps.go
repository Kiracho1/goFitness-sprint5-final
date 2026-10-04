package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	stepAndTime := strings.Split(datastring, ",")
	if len(stepAndTime) != 2 {
		return fmt.Errorf("длина слайса должна быть равна 2")
	}

	step, err := strconv.Atoi(stepAndTime[0])
	if err != nil {
		return err
	}
	// Проверяем на отрицательное значение количества шагов
	if step <= 0 {
		return fmt.Errorf("количество шагов должно быть строго больше 0")
	}
	ds.Steps = step

	duration, err := time.ParseDuration(stepAndTime[1])
	if err != nil {
		return err
	}
	// Проверяем на отрицательное значение длительности тренировки
	if duration <= 0 {
		return fmt.Errorf("длительность тренировки должна быть больше 0")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var result string
	// Рассчитываем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Рассчитываем количество калорий
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result = fmt.Sprintf(
		"Количество шагов: %d.\n" +
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, spentCalories,
	)
	return result, nil
}