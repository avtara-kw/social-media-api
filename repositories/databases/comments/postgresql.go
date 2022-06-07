package comments

import (
	"database/sql"
	"github.com/avtara-kw/social-media-api/businesses/comments"
)

type repoComments struct {
	DB *sql.DB
}

func NewRepoPostgresql(db *sql.DB) comments.Repository {
	return &repoComments{
		DB: db,
	}
}

func (ru *repoComments) Store(data *comments.Domain) (*comments.Domain, error) {
	comment := fromDomain(data)

	sqlStatement := `
INSERT INTO comments (user_id, photo_id, message)
VALUES ($1, $2, $3) Returning *;
`

	if err := ru.DB.QueryRow(sqlStatement, comment.UserID, comment.PhotoID, comment.Message).
		Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(comment)
	return result, nil
}

func (ru *repoComments) GetAll() ([]comments.Domain, error) {
	var result []comments.Domain
	sqlStatement := `
SELECT c.id, c.message, c.photo_id, c.user_id, c.created_at, c.updated_at, u.id, u.email, u.username, p.id, p.title, p.caption, p.photo_url, p.user_id
FROM comments c
JOIN photos p on c.photo_id = p.id
JOIN users u on u.id = c.user_id;
`

	rows, err := ru.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment comments.Domain
		err = rows.Scan(&comment.ID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt,
			&comment.UpdatedAt, &comment.Users.ID, &comment.Users.Email, &comment.Users.Username, &comment.Photo.ID,
			&comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.PhotoURL, &comment.Photo.UserID)
		if err != nil {
			return nil, err
		}
		result = append(result, comment)
	}

	return result, nil
}

func (ru *repoComments) Update(data *comments.Domain, socialMediaID string) (*comments.Domain, error) {
	var comment Comments
	sqlStatement := `
UPDATE comments
SET message = $2, updated_at=current_timestamp
WHERE id = $1 AND user_id = $3 returning *;
`
	if err := ru.DB.QueryRow(sqlStatement, socialMediaID, data.Message, data.UserID).
		Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message,
			&comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(&comment)
	return result, nil
}

func (ru *repoComments) DeleteByID(ID int) error {
	sqlStatement := `
DELETE FROM comments WHERE id=$1;
`
	if _, err := ru.DB.Exec(sqlStatement, ID); err != nil {
		return err
	}
	return nil
}

func (ru *repoComments) GetByID(ID int) (*comments.Domain, error) {
	var comment Comments

	sqlStatement := `
SELECT * FROM comments where user_id=$1;
`
	if err := ru.DB.QueryRow(sqlStatement, ID).
		Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message,
			&comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}
	result := toDomain(&comment)
	return result, nil
}
