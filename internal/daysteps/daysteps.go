package daysteps

import (
	"errors"
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
	parseData := strings.Split(datastring, ",")

	if len(parseData) != 2 {
		return errors.New("invalid data")
	}

	steps, err := strconv.Atoi(parseData[0])

	if err != nil {
		return errors.New("invalid steps")
	}

	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}

	duration, err := time.ParseDuration(parseData[1])

	if err != nil {
		return errors.New("invalid duration")
	}

	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
