package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	// "github.com/joho/godotenv"
	"evolve-credit/pkg/models"
	"log"
	"time"
)

// func init() {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db, err := openDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	model := newModel(db)
	where := ""
	query := ""
	defaultLimit := 10
	var queryParams map[string]string 
	var limit int

	if len(params) > 0 {
		queryParams = params[0]
		if from, ok := queryParams["from"]; ok {
			if to, ok := queryParams["to"]; ok {
				where = fmt.Sprintf(`
				WHERE date::timestamp >= '%s'::timestamp AND  
				date::timestamp <= '%s'::timestamp`, 
				from, to)
			}
		} 

		query = fmt.Sprintf(`
			SELECT * 
			FROM users 
			%s
			LIMIT $1
			OFFSET $2
		`, where)
	}

	limit, err = strconv.Atoi(queryParams["limit"])
	if err != nil {
		limit = defaultLimit
	}
	offset, _ := strconv.Atoi(queryParams["offset"])


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

	row, err := model.db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer row.Close()

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
