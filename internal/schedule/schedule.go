package schedule

import (
	"strings"
)

type Day struct {
	Date   string
	Time   [5]string
	Lesson [5]string
}

func (d Day) ToMessage() string {
	var b strings.Builder

	b.WriteString(d.Date + "\n")
	for i := 0; i < len(d.Time); i++ {
		if d.Time[i] == "" && d.Lesson[i] == "" {
			continue // пропуск пустых строк
		}
		b.WriteString(d.Time[i] + " — " + d.Lesson[i] + "\n")
		b.WriteString("------------------------------" + "\n")

	}

	return b.String()
}

type Schedule struct {
	WeekDay [6]Day
}

func New() *Schedule {
	return &Schedule{}
}

func (s *Schedule) LoadSchedule() error {
	err := ParseSchedule(s)
	if err != nil {
		return err
	}
	return nil
}
