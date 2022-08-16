package firesert

type Configuration struct {
	ProjectID      string
	CollectionName string
}

func NewConfiguration(projectID string, collectionName string) *Configuration {
	this := new(Configuration)

	this.ProjectID = projectID
	this.CollectionName = collectionName

	return this
}
