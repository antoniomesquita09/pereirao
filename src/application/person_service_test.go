package application_test

import (
	"pereirao/src/application"
	mock_application "pereirao/src/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPersonService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	person := mock_application.NewMockPersonInterface(ctrl)
	persistence := mock_application.NewMockPersonPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(person, nil).AnyTimes()

	service := application.PersonService{Persistence: persistence}

	result, err := service.GetByID(uuid.New())
	require.Nil(t, err)
	require.Equal(t, person, result)
}
