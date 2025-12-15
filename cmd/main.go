package main

import (
	"os"

	"linkedin-automation-demo/auth"
	"linkedin-automation-demo/config"
	"linkedin-automation-demo/messaging"
	"linkedin-automation-demo/search"
	"linkedin-automation-demo/stealth"
	"linkedin-automation-demo/storage"
	"linkedin-automation-demo/utils"
)

func main() {
	utils.Info("Starting LinkedIn Automation Demo")

	cfg := config.LoadConfig()

	if cfg.Email == "" || cfg.Password == "" {
		utils.Error("Missing credentials. Set environment variables.")
		os.Exit(1)
	}

	if !stealth.IsWithinWorkHours(cfg.WorkStartHour, cfg.WorkEndHour) {
		utils.Warn("Outside working hours. Automation paused.")
		return
	}

	if err := auth.Login(cfg.Email, cfg.Password); err != nil {
		utils.Error("Login failed")
		return
	}

	profiles := search.SearchProfiles()

	messenger := messaging.Messenger{
		DailyLimit: cfg.DailyConnectionLimit,
	}

	for i := range profiles {
		messenger.SendConnection(&profiles[i])
		messenger.SendMessage(profiles[i])
	}

	state := storage.State{
		ConnectionsSent: messenger.SentCount,
		Profiles:        profiles,
	}

	if err := storage.SaveState(state); err != nil {
		utils.Error("Failed to save state")
	}

	utils.Info("Automation demo completed safely")
}
