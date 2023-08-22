package pgrepo

import (
	"context"
	"fmt"
	"github.com/6a6ydoping/Pinky/internal/entity"
	"time"
)

func (p *Postgres) CreatePicture(ctx context.Context, picture *entity.Picture) error {
	query := fmt.Sprintf(`INSERT INTO %s (
                url, -- 1
                date	-- 2
) VALUES ($1, $2)`, pictureTable)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := p.Pool.Exec(ctx, query, picture.URL, picture.Date)
	if err != nil {
		return err
	}
	return nil
}
