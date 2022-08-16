package firesert

func BuildApplication[T any]() Application {
	configurationFilePathProvider := ConfigurationFilePathProvider{"firesert-config"}

	configurationFileReader := ConfigurationJsonFileReader{
		configurationFilePathProvider,
	}

	configuratonCreator := JsonConfigurationCreator{}

	configurationLoader := JsonConfigurationLoader{
		configurationFileReader,
		configuratonCreator}

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
