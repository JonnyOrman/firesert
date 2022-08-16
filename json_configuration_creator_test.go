package firesert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	json := []byte("{\"ProjectID\":\"my-project\",\"collectionName\":\"MyCollection\"}")

	sut := JsonConfigurationCreator{}

	result := sut.Create(json)

	assert.Equal(t, "my-project", result.ProjectID)
	assert.Equal(t, "MyCollection", result.CollectionName)
}
