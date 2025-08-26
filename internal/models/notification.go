package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID         uuid.UUID `json:"id"`
	Message    string    `json:"name"`
	SendAt     time.Time `json:"send_at"`
	Status     string    `json:"status"`
	RetryCount int       `json:"retry_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
