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

func (p *Postgres) GetPicturesByRange(ctx context.Context, leftBound, rightBound uint) (*[]entity.Picture, error) {
	query := fmt.Sprintf(`SELECT id, url, date FROM %s WHERE id BETWEEN $1 AND $2`, pictureTable)

	rows, err := p.Pool.Query(ctx, query, leftBound, rightBound)
	if err != nil {
		fmt.Println(2, err)
		return nil, err
	}
	defer rows.Close()

	var pictures []entity.Picture

	for rows.Next() {
		var pic entity.Picture
		if err := rows.Scan(&pic.ID, &pic.URL, &pic.Date); err != nil {
			fmt.Println(1, err)
			return nil, err
		}
		pictures = append(pictures, pic)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pictures, nil
}
