package helpers

import "time"

func CalculateTimeSinceCreation(t string) (int, error) {
	parsedTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return 0, err
	}

	duration := time.Since(parsedTime)
	minutes := int(duration.Minutes())

	return minutes, nil
}
