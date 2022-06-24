package models

import (
	"context"
	"database/sql"
	"fmt"
	// "log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetAll(pages ...int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""

	if len(pages) > 0 {
		where = fmt.Sprintf(
			`WHERE id IN 
			(SELECT movie_id FROM movies_genres WHERE genre_id = %d)`, pages[0])
	}

	query := fmt.Sprintf(`
		SELECT id, title, description, year, release_date, 
				rating, runtime, mpaa_rating, created_at, updated_at, COALESCE(poster, '')
		FROM movies %s ORDER BY title`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Date,
			&user.Country,
		)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}
