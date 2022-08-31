package firesert

import "github.com/jonnyorman/fireworks"

func BuildApplication[T any]() *fireworks.Application {
	configurationFilePathProvider := fireworks.NewConfigurationFilePathProvider("firesert-config")

	configurationFileReader := fireworks.NewConfigurationFileReader(configurationFilePathProvider)

	configurationJsonFileReader := fireworks.NewConfigurationJsonFileReader(configurationFileReader)

	configurationJson := configurationJsonFileReader.Read()

	projectIDProvider := fireworks.CreateConfigurationValueProvider("projectID", "PROJECT_ID", configurationJson)

	collectionNameProvider := fireworks.CreateConfigurationValueProvider("collectionName", "COLLECTION_NAME", configurationJson)

	configurationLoader := fireworks.NewApplicationConfigurationLoader(
		projectIDProvider,
		collectionNameProvider,
	)

	configuration := configurationLoader.Load()

	pubSubBodyDeserialiser := fireworks.JsonDataDeserialiser[fireworks.PubSubBody]{}

	ioutilReader := fireworks.IoutilReader{}

	pubSubBodyReader := fireworks.NewGinPubSubBodyReader(
		ioutilReader,
		pubSubBodyDeserialiser)

	dataDeserialiser := fireworks.JsonDataDeserialiser[T]{}

	dataReader := fireworks.NewHttpRequestBodyDataReader[T](
		pubSubBodyReader,
		dataDeserialiser)

	dataInserter := FirestoreDataInserter[T]{configuration}

	requestHandler := PubSubPushRequestHandler[T]{
		dataReader:   dataReader,
		dataInserter: dataInserter,
	}

	routerBuilder := GinRouterBuilder[T]{requestHandler}

	router := routerBuilder.Build()

	app := fireworks.NewApplication(router)

	return app
}
