package social_medias

import (
	"database/sql"
	"github.com/avtara-kw/social-media-api/businesses/social_medias"
)

type repoSocialMedias struct {
	DB *sql.DB
}

func NewRepoPostgresql(db *sql.DB) social_medias.Repository {
	return &repoSocialMedias{
		DB: db,
	}
}

func (ru *repoSocialMedias) Store(data *social_medias.Domain) (*social_medias.Domain, error) {
	socialMedia := fromDomain(data)

	sqlStatement := `
INSERT INTO social_medias (user_id, social_media_url, name)
VALUES ($1, $2, $3) Returning *;
`

	if err := ru.DB.QueryRow(sqlStatement, socialMedia.UserID, socialMedia.SocialMediaURL, socialMedia.Name).
		Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserID,
			&socialMedia.CreatedAt, &socialMedia.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(socialMedia)
	return result, nil
}

func (ru *repoSocialMedias) GetAllByID(id string) ([]social_medias.Domain, error) {
	var result []social_medias.Domain
	sqlStatement := `
SELECT social_medias.id, social_medias.name, social_media_url, user_id, social_medias.created_at, social_medias.updated_at, username
FROM social_medias
JOIN users u on u.id = social_medias.user_id
WHERE user_id = $1;
`

	rows, err := ru.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var socialMedia social_medias.Domain
		err = rows.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserID,
			&socialMedia.CreatedAt, &socialMedia.UpdatedAt, &socialMedia.Users.Username)
		if err != nil {
			return nil, err
		}
		result = append(result, socialMedia)
	}

	return result, nil
}

func (ru *repoSocialMedias) Update(data *social_medias.Domain, socialMediaID string) (*social_medias.Domain, error) {
	var socialMedia SocialMedias
	sqlStatement := `
UPDATE social_medias
SET name = $2, social_media_url = $3, updated_at=current_timestamp
WHERE id = $1 AND user_id = $4 returning *;
`
	if err := ru.DB.QueryRow(sqlStatement, socialMediaID, data.Name, data.SocialMediaURL, data.UserID).
		Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaURL, &socialMedia.UserID,
			&socialMedia.CreatedAt, &socialMedia.UpdatedAt); err != nil {
		return nil, err
	}

	result := toDomain(&socialMedia)
	return result, nil
}

func (ru *repoSocialMedias) DeleteByID(ID int) error {
	sqlStatement := `
DELETE FROM social_medias WHERE id=$1;
`
	if _, err := ru.DB.Exec(sqlStatement, ID); err != nil {
		return err
	}
	return nil
}

func (ru *repoSocialMedias) GetByID(ID int) (*social_medias.Domain, error) {
	var photo SocialMedias

	sqlStatement := `
SELECT * FROM social_medias where user_id=$1;
`
	if err := ru.DB.QueryRow(sqlStatement, ID).
		Scan(&photo.ID, &photo.Name, &photo.SocialMediaURL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
		return nil, err
	}
	result := toDomain(&photo)
	return result, nil
}
