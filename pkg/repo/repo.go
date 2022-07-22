package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"evolve-credit/pkg/models"
	"log"
	"time"
)

type DBModel struct {
	db *sql.DB
}

func newModel(db *sql.DB) DBModel {
	return DBModel{db: db}
}


func openDB(conString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conString)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAll(params ...map[string]string) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queryParams := params[0]
	model := newModel(db)
	where := ""
	query := ""
	defaultLimit := 10
	defaultOffset := 0
	var limit int
	var offset int

	if from := queryParams["from"]; from != "" {
		if to := queryParams["to"]; to != "" {
			where = fmt.Sprintf(`
			WHERE date::timestamp >= '%s'::timestamp AND  
			date::timestamp <= '%s'::timestamp`, 
			from, to)
		}
	} 

	// Set to default limit if value was not provided
	limit, err = strconv.Atoi(queryParams["limit"])
	if err != nil {
		limit = defaultLimit
	}

	// Set to default offset if value was not provided
	offset, err = strconv.Atoi(queryParams["offset"])
	if err != nil {
		offset = defaultOffset
	} 

	query = fmt.Sprintf(`
		SELECT * 
		FROM users 
		%s
		LIMIT $1
		OFFSET $2
	`, where)

	rows, err := model.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
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
		users = append(users, &user)
	}

	return users, nil
}

func Get(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	model := newModel(db)
	
	query := `
		SELECT * 
		FROM users 
		WHERE email = $1
	`

	row := model.db.QueryRowContext(ctx, query, email)

	var user models.User

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Date,
		&user.Country,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
