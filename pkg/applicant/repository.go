package applicant

import "database/sql"

type Repository interface {
	CreateBasicInfo()
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBasicInfo() {

}
