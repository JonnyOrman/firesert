package firesert

func BuildApplication[T any]() Application {
	configurationIoReaderGenerator := ConfigurationIoReaderGenerator{}

	reader := IoutilReader{}

	configurationFileReader := ConfigurationJsonFileReader{
		configurationIoReaderGenerator,
		reader}

	configuratonCreator := JsonConfigurationCreator{}

	configurationLoader := JsonConfigurationLoader{
		configurationFileReader,
		configuratonCreator}

	configuration := configurationLoader.Load()

	pubSubBodyDeserialiser := JsonDataDeserialiser[PubSubBody]{}

	pubSubBodyReader := GinPubSubBodyReader{
		reader,
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
