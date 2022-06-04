package users

import (
	"database/sql"
	"github.com/avtara-kw/social-media-api/businesses/users"
)

type repoUsers struct {
	DB *sql.DB
}

func NewRepoMySQL(db *sql.DB) users.Repository {
	return &repoUsers{
		DB: db,
	}
}

func (ru *repoUsers) Store(data *users.Domain) (*users.Domain, error) {
	user := fromDomain(data)

	sqlStatement := `
INSERT INTO users (username, email, password, age)
VALUES ($1, $2, $3, $4) Returning *
`

	if err := ru.DB.QueryRow(sqlStatement, user.Username, user.Email, user.Password, user.Age).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(user)
	return result, nil
}
