package repository

import (
	"context"
	"github.com/6a6ydoping/Pinky/internal/entity"
)

type Repository interface {
	CreatePicture(ctx context.Context, picture *entity.Picture) error
	GetPicturesByRange(ctx context.Context, leftBound, rightBound uint) (*[]entity.Picture, error)
}
