package stealth

import (
	"math/rand"
	"time"

	"linkedin-automation-demo/utils"
)

// Random delay between actions
func RandomDelay(minSec, maxSec int) {
	delay := rand.Intn(maxSec-minSec+1) + minSec
	utils.Info("Human delay: waiting " + time.Duration(delay).String() + " seconds")
	time.Sleep(time.Duration(delay) * time.Second)
}

// Simulate typing behavior
func SimulateTyping(text string) {
	for _, ch := range text {
		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)
		_ = ch
	}
}

// Check if current time is within work hours
func IsWithinWorkHours(start, end int) bool {
	hour := time.Now().Hour()
	return hour >= start && hour < end
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
