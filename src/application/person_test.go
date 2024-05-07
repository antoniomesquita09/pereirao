package application_test

import (
	"pereirao/src/application"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerson_IsValid(t *testing.T) {
	person := application.NewPerson("dummy name", "email", "cpf")

	ok, err := person.IsValid()
	require.Equal(t, false, ok)
	require.Equal(t, "invalid CPF", err.Error())

	person.Cpf = "39727465048"
	ok, err = person.IsValid()
	require.Equal(t, false, ok)
	require.Equal(t, "invalid email", err.Error())

	person.Email = "dummy@gmail.com"
	ok, err = person.IsValid()
	require.Equal(t, true, ok)
	require.Nil(t, err)
}
