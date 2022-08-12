package firesert

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
)

type FiresertConfig struct {
	ProjectID      string
	CollectionName string
}

type PubSubBody struct {
	Message      pubsub.Message
	Subscription string
}

var firesertConfig FiresertConfig

func Run() {
	RunTyped[map[string]interface{}]()
}

func RunTyped[T any]() {
	firesertConfigJson, _ := ioutil.ReadFile("./firesert-config.json")

	_ = json.Unmarshal(firesertConfigJson, &firesertConfig)

	router := gin.Default()

	router.POST("/", handle[T])

	router.Run()
}

func handle[T any](ginContext *gin.Context) {
	client, _ := firestore.NewClient(context.Background(), firesertConfig.ProjectID)

	defer client.Close()

	bodyAsByteArray, _ := ioutil.ReadAll(ginContext.Request.Body)

	var pubsubBody PubSubBody
	_ = json.Unmarshal(bodyAsByteArray, &pubsubBody)

	var unmarshaledData T
	json.Unmarshal([]byte(string(pubsubBody.Message.Data)), &unmarshaledData)

	collection := client.Collection(firesertConfig.CollectionName)

	_, _, _ = collection.Add(context.Background(), unmarshaledData)
}
