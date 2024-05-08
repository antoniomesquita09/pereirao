package db

import (
	"database/sql"
	"pereirao/src/application"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PersonDb struct {
	db *sql.DB
}

func (p *PersonDb) Get(id uuid.UUID) (application.PersonInterface, error) {
	var person application.Person
	stmt, err := p.db.Prepare("select id, name, email, cpf from person where id=$1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&person.ID, &person.Name, &person.Email, &person.Cpf)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
