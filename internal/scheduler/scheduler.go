package scheduler

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"../email"
)

type ScheduledEmail struct {
	ID        string      `json:"id"`
	Email     email.Email `json:"email"`
	SendAt    time.Time   `json:"send_at"`
	Recurring string      `json:"recurring,omitempty"` // daily, weekly, monthly
}

type Scheduler struct {
	ScheduledEmails []ScheduledEmail `json:"scheduled_emails"`
	ConfigPath      string
}

func NewScheduler() (*Scheduler, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".email-cli")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, "scheduled.json")
	s := &Scheduler{ConfigPath: configPath}

	if err := s.load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return s, nil
}

func (s *Scheduler) Schedule(email email.Email, sendAt time.Time, recurring string) error {
	id := time.Now().Format("20060102150405")
	s.ScheduledEmails = append(s.ScheduledEmails, ScheduledEmail{
		ID:        id,
		Email:     email,
		SendAt:    sendAt,
		Recurring: recurring,
	})

	return s.save()
}

func (s *Scheduler) Cancel(id string) error {
	for i, se := range s.ScheduledEmails {
		if se.ID == id {
			s.ScheduledEmails = append(s.ScheduledEmails[:i], s.ScheduledEmails[i+1:]...)
			return s.save()
		}
	}
	return nil
}

func (s *Scheduler) load() error {
	data, err := os.ReadFile(s.ConfigPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.ScheduledEmails)
}

func (s *Scheduler) save() error {
	data, err := json.MarshalIndent(s.ScheduledEmails, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.ConfigPath, data, 0600)
}