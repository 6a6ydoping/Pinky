package service

import (
	"context"
	"errors"
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

func (m *Manager) GetPictureByPage(ctx context.Context, page, perPage uint) (*[]entity.Picture, error) {
	if page == 0 {
		return nil, errors.New("incorrect page value")
	}
	idsLeftBound := (page-1)*perPage + 1
	idsRightBound := idsLeftBound + perPage - 1
	pics, err := m.Repository.GetPicturesByRange(ctx, idsLeftBound, idsRightBound)
	if err != nil {
		return nil, err
	}

	return pics, nil
}
