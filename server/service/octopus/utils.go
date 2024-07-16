package octopus

import (
	"fmt"
	"time"
)

func Today() (today string) {
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()

	today = fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	return today
}
