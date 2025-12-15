package search

import (
	"linkedin-automation-demo/models"
	"linkedin-automation-demo/utils"
)

func SearchProfiles() []models.Profile {
	utils.Info("Searching profiles (simulated)")

	profiles := []models.Profile{
		{Name: "John Doe", Title: "Software Engineer", Company: "ABC Corp"},
		{Name: "Jane Smith", Title: "Backend Developer", Company: "XYZ Ltd"},
		{Name: "Rahul Kumar", Title: "Full Stack Dev", Company: "TechSoft"},
		{Name: "Anita Rao", Title: "QA Engineer", Company: "Innovate"},
		{Name: "David Lee", Title: "Cloud Engineer", Company: "CloudNine"},
	}

	utils.Info("Profiles fetched successfully")
	return profiles
}
