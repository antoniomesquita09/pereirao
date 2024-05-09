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

func (p *PersonDb) Save(product application.PersonInterface) (application.PersonInterface, error) {
	var rows int
	p.db.QueryRow("Select count(*) from person where id=$1", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *PersonDb) create(person application.PersonInterface) (application.PersonInterface, error) {
	stmt, err := p.db.Prepare(`insert into person(id, name, email, cpf) values($1,$2,$3,$4)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		person.GetID(),
		person.GetName(),
		person.GetEmail(),
		person.GetCpf(),
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p *PersonDb) update(product application.PersonInterface) (application.PersonInterface, error) {
	_, err := p.db.Exec(
		"update person set name=$1, email=$2, cpf=$3 where id=$4",
		product.GetName(),
		product.GetEmail(),
		product.GetCpf(),
		product.GetID(),
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
