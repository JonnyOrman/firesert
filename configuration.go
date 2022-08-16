package firesert

type Configuration struct {
	ProjectID      string
	CollectionName string
}

func NewConfiguration(projectID string, collectionName string) *Configuration {
	configuration := new(Configuration)
	configuration.ProjectID = projectID
	configuration.CollectionName = collectionName
	return configuration
}
