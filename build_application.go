package firesert

import "github.com/jonnyorman/fireworks"

func BuildApplication[T any]() *fireworks.Application {
	configuration := fireworks.GenerateConfiguration("firesert-config")

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

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddPost("/", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
