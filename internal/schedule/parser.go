package schedule

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func ParseSchedule(s *Schedule) error {
	wb, err := xlsx.OpenFile("./currentSchedule/newSchedule.xlsx")
	if err != nil {
		return fmt.Errorf("open newSchedule.xlsx: %v", err)
	}
	sh, ok := wb.Sheet["4к 9кл, 3к 11кл"]
	if !ok {
		return fmt.Errorf("newSchedule.xlsx: sheet not found")
	}

	parseDate(&s.WeekDay, sh)
	parseTime(&s.WeekDay, sh)
	parseLesson(&s.WeekDay, sh)

	return nil
}

func parseDate(days *[6]Day, sh *xlsx.Sheet) {
	day := 0
	for i := 5; i < 33; i += 5 {

		theCell, err := sh.Cell(i, 0)
		if err != nil {
			fmt.Printf("error with parseDate: %v", err)
			panic(err)
		}
		days[day].Date = theCell.String()
		day++
	}
}

func parseTime(line *[6]Day, sh *xlsx.Sheet) {
	day := 0
	time := 0
	for i := 5; i < 35; i++ {
		if time > 4 {
			day++
			time = 0
		}
		theCell, err := sh.Cell(i, 1)
		if err != nil {
			fmt.Printf("error with parseTime: %v", err)
			panic(err)
		}
		line[day].Time[time] = theCell.String()
		time++
	}
}

func parseLesson(line *[6]Day, sh *xlsx.Sheet) {
	day := 0
	lesson := 0
	for i := 5; i < 35; i++ {
		if lesson > 4 {
			day++
			lesson = 0
		}
		theCell, err := sh.Cell(i, 3)
		if err != nil {
			fmt.Printf("error with parseTime: %v", err)
			panic(err)
		}
		line[day].Lesson[lesson] = theCell.String()
		lesson++
	}
}
