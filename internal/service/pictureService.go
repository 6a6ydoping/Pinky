package service

import (
	"context"
	"github.com/6a6ydoping/Pinky/internal/entity"
	"time"
)

func (m *Manager) CreatePicture(ctx context.Context, p *entity.Picture) error {
	currentTime := time.Now()
	p.Date = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)
	err := m.Repository.CreatePicture(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
