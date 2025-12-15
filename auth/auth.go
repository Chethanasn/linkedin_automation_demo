package auth

import (
	"errors"

	"linkedin-automation-demo/utils"
)

func Login(email, password string) error {
	utils.Info("Attempting login")

	if email == "" || password == "" {
		return errors.New("invalid credentials")
	}

	// Simulation only
	utils.Info("Login successful (simulated)")
	return nil
}
