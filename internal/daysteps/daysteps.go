package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dsSlice := strings.Split(datastring, ",")
	if len(dsSlice) != 2 {
		return errors.New("количество полей в строке после разделения по ',' не равно 2")
	}
	steps, err := strconv.Atoi(dsSlice[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("некорректное количество шагов")
	}
	duration, err := time.ParseDuration(dsSlice[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("некорректное значение переменной duration")
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
