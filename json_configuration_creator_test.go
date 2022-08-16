package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	json := []byte("{\"ProjectID\":\"my-project\",\"collectionName\":\"MyCollection\"}")

	jsonConfigurationCreator := JsonConfigurationCreator{}

	configuration := jsonConfigurationCreator.Create(json)

	assert.Equal(t, "my-project", configuration.ProjectID)
	assert.Equal(t, "MyCollection", configuration.CollectionName)
}
