package utils

import (
	"log"
	"time"
)

func Info(msg string) {
	log.Printf("[INFO] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}

func Warn(msg string) {
	log.Printf("[WARN] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}

func Error(msg string) {
	log.Printf("[ERROR] %s | %s\n", time.Now().Format(time.RFC3339), msg)
}
