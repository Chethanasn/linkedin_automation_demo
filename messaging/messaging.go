package messaging

import (
	"linkedin-automation-demo/models"
	"linkedin-automation-demo/stealth"
	"linkedin-automation-demo/utils"
)

type Messenger struct {
	DailyLimit int
	SentCount  int
}

func (m *Messenger) SendConnection(profile *models.Profile) {
	if m.SentCount >= m.DailyLimit {
		utils.Warn("Daily connection limit reached. Cooling down.")
		return
	}

	stealth.RandomDelay(2, 5)

	utils.Info("Sending connection to " + profile.Name)
	profile.Accepted = true
	m.SentCount++
}

func (m *Messenger) SendMessage(profile models.Profile) {
	if profile.Accepted {
		stealth.RandomDelay(1, 3)
		utils.Info("Typing message to " + profile.Name)
		stealth.SimulateTyping("Hi " + profile.Name + ", happy to connect!")
		utils.Info("Message sent to " + profile.Name)
	}
}
