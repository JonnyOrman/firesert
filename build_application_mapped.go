package firesert

import "github.com/jonnyorman/fireworks"

func BuildApplicationMapped[TReceive any, TInsert any](dataMapper DataMapper[TReceive, TInsert]) *fireworks.Application {
	configuration := fireworks.GenerateConfiguration("firesert-config")

	pubSubBodyDeserialiser := fireworks.JsonDataDeserialiser[fireworks.PubSubBody]{}

	ioutilReader := fireworks.IoutilReader{}

	pubSubBodyReader := fireworks.NewGinPubSubBodyReader(
		ioutilReader,
		pubSubBodyDeserialiser)

	dataDeserialiser := fireworks.JsonDataDeserialiser[TReceive]{}

	dataReader := fireworks.NewHttpRequestBodyDataReader[TReceive](
		pubSubBodyReader,
		dataDeserialiser)

	dataInserter := FirestoreDataInserter[TInsert]{configuration}

	requestHandler := PubSubPushRequestHandlerMapped[TReceive, TInsert]{
		dataReader:   dataReader,
		dataMapper:   dataMapper,
		dataInserter: dataInserter,
	}

	routerBuilder := fireworks.NewGinRouterBuilder()

	routerBuilder.AddPost("/", requestHandler.Handle)

	router := routerBuilder.Build()

	application := fireworks.NewApplication(router)

	return application
}
