package application

import "github.com/google/uuid"

type PersonService struct {
	Persistence PersonPersistenceInterface // dependency injection
}

func NewPersonService(persistence PersonPersistenceInterface) *PersonService {
	return &PersonService{Persistence: persistence}
}

func (s *PersonService) GetByID(id uuid.UUID) (PersonInterface, error) {
	person, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return person, err
}

func (s *PersonService) Create(name, email, cpf string) (PersonInterface, error) {
	person := NewPerson(name, email, cpf)

	isValid, err := person.IsValid()
	if !isValid {
		return nil, err
	}

	p, err := s.Persistence.Save(person)
	if err != nil {
		return nil, err
	}

	return p, nil
}
