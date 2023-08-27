package service

import (
	"context"
	"github.com/6a6ydoping/Pinky/internal/entity"
)

type Service interface {
	PictureService
	UserService
}

type PictureService interface {
	CreatePicture(ctx context.Context, picture *entity.Picture) error
	GetPictureByPage(ctx context.Context, page, perPage uint) (*[]entity.Picture, error)
}

type UserService interface {
}
