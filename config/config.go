package config

import (
	"os"
	"strconv"
)

type Config struct {
	Email                string
	Password             string
	DailyConnectionLimit int
	WorkStartHour        int
	WorkEndHour          int
}

func LoadConfig() Config {
	limit, _ := strconv.Atoi(getEnv("DAILY_CONNECTION_LIMIT", "10"))
	startHour, _ := strconv.Atoi(getEnv("WORK_START_HOUR", "9"))
	endHour, _ := strconv.Atoi(getEnv("WORK_END_HOUR", "18"))

	return Config{
		Email:                getEnv("LINKEDIN_EMAIL", ""),
		Password:             getEnv("LINKEDIN_PASSWORD", ""),
		DailyConnectionLimit: limit,
		WorkStartHour:        startHour,
		WorkEndHour:          endHour,
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
