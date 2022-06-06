package photos

import (
	"database/sql"
	"github.com/avtara-kw/social-media-api/businesses/photos"
)

type repoPhotos struct {
	DB *sql.DB
}

func NewRepoPostgresql(db *sql.DB) photos.Repository {
	return &repoPhotos{
		DB: db,
	}
}

func (ru *repoPhotos) Store(data *photos.Domain) (*photos.Domain, error) {
	photo := fromDomain(data)

	sqlStatement := `
INSERT INTO photos (title, photo_url, caption, user_id)
VALUES ($1, $2, $3, $4) Returning *;
`

	if err := ru.DB.QueryRow(sqlStatement, photo.Title, photo.PhotoURL, photo.Caption, photo.UserID).
		Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(photo)
	return result, nil
}

func (ru *repoPhotos) GetAll() ([]photos.Domain, error) {
	var result []photos.Domain
	sqlStatement := `
select photos.id, title, caption, photo_url, user_id, photos.created_at, photos.updated_at, email, username
FROM photos
JOIN users u on u.id = photos.user_id;
`

	rows, err := ru.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var photo photos.Domain
		err = rows.Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID,
			&photo.CreatedAt, &photo.UpdatedAt, &photo.Users.Email, &photo.Users.Username)
		if err != nil {
			return nil, err
		}
		result = append(result, photo)
	}

	return result, nil
}

func (ru *repoPhotos) DeleteByID(ID int) error {
	sqlStatement := `
DELETE FROM photos WHERE id=$1;
`
	if _, err := ru.DB.Exec(sqlStatement, ID); err != nil {
		return err
	}
	return nil
}
func (ru *repoPhotos) Update(data *photos.Domain, photoID string) (*photos.Domain, error) {
	var photo Photos
	sqlStatement := `
UPDATE photos
SET title = $2, caption = $3, photo_url=$4, updated_at=current_timestamp
WHERE id = $1 AND user_id = $5 returning *;
`
	if err := ru.DB.QueryRow(sqlStatement, photoID, data.Title, data.Caption, data.PhotoURL, data.UserID).
		Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(&photo)
	return result, nil
}

func (ru *repoPhotos) GetByID(ID int) (*photos.Domain, error) {
	var photo Photos

	sqlStatement := `
SELECT * FROM photos where user_id=$1;
`
	if err := ru.DB.QueryRow(sqlStatement, ID).
		Scan(&photo.ID, &photo.Title, &photo.Caption, &photo.PhotoURL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
		return nil, err
	}
	result := toDomain(&photo)
	return result, nil
}
