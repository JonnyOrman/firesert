package firesert

type ApplicationConfigurationLoader struct {
	projectIDProvider      ValueProvider
	collectionNameProvider ValueProvider
}

func NewApplicationConfigurationLoader(
	projectIDProvider ValueProvider,
	collectionNameProvider ValueProvider,
) *ApplicationConfigurationLoader {
	this := new(ApplicationConfigurationLoader)

	this.projectIDProvider = projectIDProvider
	this.collectionNameProvider = collectionNameProvider

	return this
}

func (this ApplicationConfigurationLoader) Load() Configuration {
	projectID, _ := this.projectIDProvider.Get()
	collectionName, _ := this.collectionNameProvider.Get()

	return Configuration{projectID, collectionName}
}
