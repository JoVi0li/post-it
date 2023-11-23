package repository

import (
	"database/sql"

	"github.com/joviolis/post-it/entity"
	"github.com/joviolis/post-it/schema"
)

type PostItRepository struct {
	db *sql.DB
}

func (p *PostItRepository) Create(data schema.PostItSchema) (post entity.PostItEntity, err error) {
	defer p.db.Close()
	return entity.PostItEntity{}, nil
}

func (p *PostItRepository) Get(id string) (post entity.PostItEntity, err error) {
	return entity.PostItEntity{}, nil
}

func (p *PostItRepository) List() (posts []entity.PostItEntity, err error) {
	return []entity.PostItEntity{}, nil
}

func (p *PostItRepository) Update(id string, data schema.PostItSchema) (post entity.PostItEntity, err error) {
	return entity.PostItEntity{}, nil
}

func Delete(id string) (post entity.PostItEntity, err error) {
	return entity.PostItEntity{}, nil
}
