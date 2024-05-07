package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/klassmann/cpfcnpj"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type PersonInterface interface {
	IsValid() (bool, error)
	GetID() uuid.UUID
	GetName() string
	GetEmail() string
	GetCpf() string
}

type PersonServiceInterface interface {
	GetByID(id uuid.UUID) (PersonInterface, error)
	GetByCpf(cpf string) (PersonInterface, error)
	Create(name, email, cpf string) (PersonInterface, error)
}

type PersonReader interface {
	Get(id uuid.UUID) (PersonInterface, error)
}

type PersonWriter interface {
	Save(person PersonInterface) (PersonInterface, error)
}

type PersonPersistenceInterface interface { // composition
	PersonReader
	PersonWriter
}

type Person struct {
	ID    uuid.UUID `valid:"uuidv4"`
	Name  string    `valid:"required"`
	Email string    `valid:"required,email"`
	Cpf   string    `valid:"required"`
}

func NewPerson(name, email, cpf string) *Person {
	person := Person{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
		Cpf:   cpf,
	}

	return &person
}

func (p *Person) IsValid() (bool, error) {
	if !p.isValidCpf() {
		return false, errors.New("invalid CPF")
	}

	if !govalidator.IsEmail(p.Email) {
		return false, errors.New("invalid email")
	}

	if len(p.Name) > 50 {
		return false, errors.New("name is too long")
	}

	return true, nil
}

func (p *Person) isValidCpf() bool {
	return cpfcnpj.ValidateCPF(p.Cpf)
}

func (p *Person) GetID() uuid.UUID {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetEmail() string {
	return p.Email
}

func (p *Person) GetCpf() string {
	return p.Cpf
}
