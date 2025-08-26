package service

import (
	"BioMihanoid/DelayedNotifier/internal/models"
	"context"

	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateNotify(ctx context.Context, notify models.Notification) error {
	return nil
}

func (s *Service) GetNotifyStatus(ctx context.Context, id uuid.UUID) (string, error) {
	return "", nil
}

func (s *Service) DeleteNotify(ctx context.Context, id uuid.UUID) error {
	return nil
}
