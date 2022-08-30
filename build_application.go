package firesert

func BuildApplication[T any]() Application {
	configurationFilePathProvider := ConfigurationFilePathProvider{"firesert-config"}

	configurationFileReader := ConfigurationFileReader{configurationFilePathProvider}

	configurationJsonFileReader := ConfigurationJsonFileReader{configurationFileReader}

	configurationJson := configurationJsonFileReader.Read()

	projectIDProvider := CreateConfigurationValueProvider("projectID", "PROJECT_ID", configurationJson)

	collectionNameProvider := CreateConfigurationValueProvider("collectionName", "COLLECTION_NAME", configurationJson)

	configurationLoader := ApplicationConfigurationLoader{
		projectIDProvider,
		collectionNameProvider,
	}

	configuration := configurationLoader.Load()

	pubSubBodyDeserialiser := JsonDataDeserialiser[PubSubBody]{}

	ioutilReader := IoutilReader{}

	pubSubBodyReader := GinPubSubBodyReader{
		ioutilReader,
		pubSubBodyDeserialiser}

	dataDeserialiser := JsonDataDeserialiser[T]{}

	dataReader := HttpRequestBodyDataReader[T]{
		pubSubBodyReader,
		dataDeserialiser}

	dataInserter := FirestoreDataInserter[T]{configuration}

	requestHandler := PubSubPushRequestHandler[T]{
		dataReader:   dataReader,
		dataInserter: dataInserter,
	}

	routerBuilder := GinRouterBuilder[T]{requestHandler}

	router := routerBuilder.Build()

	app := Application{router}

	return app
}
