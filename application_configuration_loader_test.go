package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ValueProviderMock struct {
	mock.Mock
}

func (this ValueProviderMock) Get() (string, bool) {
	args := this.Called()
	return args.Get(0).(string), args.Get(1).(bool)
}

func TestApplicationConfigurationLoaderLoac(t *testing.T) {
	projectID := "my-project"

	collectionName := "MyCollection"

	projectIDProvider := new(ValueProviderMock)
	projectIDProvider.On("Get").Return(projectID, true)

	collectionNameProvider := new(ValueProviderMock)
	collectionNameProvider.On("Get").Return(collectionName, true)

	expectedResult := Configuration{projectID, collectionName}

	sut := ApplicationConfigurationLoader{projectIDProvider, collectionNameProvider}

	result := sut.Load()

	assert.Equal(t, expectedResult, result)
}
