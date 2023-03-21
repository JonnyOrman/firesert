# firesert

Create Go microservices that receive Pub/Sub push messages and insert their data into Firestore collections.

Very easy to use, create a working service by adding just one line of code to your `main` method:
```
firesert.Run()
```

## Examples

Try working examples of firesert [here](https://github.com/JonnyOrman/firesert-examples)

## Getting started

Create a new Go project
```
mkdir firesert-example
cd firesert-example
go mod init firesert/example
```

Get `firesert`
```
go get github.com/jonnyorman/firesert
```

Add a `main.go` file with the following
```
package main

import "github.com/jonnyorman/firesert"

func main() {
	firesert.Run()
}
```

Add a `firesert-config.json` file with the following
```
{
    "projectID": "your-firebase-project",
    "collectionName": "FirestoreCollection"
}
```

Tidy and run with access to a Firebase project or emulator
```
    go mod tidy
    go run .
```

Submit a `POST` to the service with a Pub/Sub push body. You will see the data get inserted in the `FirestoreCollection` collection in your Firestore.

You can also create a struct with the data you want to insert into Firestore. Create a struct and use it:
```
package main

import "github.com/jonnyorman/firesert"

type DocumentModel struct {
	Prop1 string
	Prop2 int
}

func main() {
	firesert.RunTyped[DocumentModel]()
}
```

## Environment configuration

The configuration can also be provided by the environment with the following keys:
- `projectID` - `PROJECT_ID`
- `collectionName` - `COLLECTION_NAME`

A combination of the `firesert-config.json` file and environment variables can be used. For example, the project ID could be provided as the `PROJECT_ID` environment variable, while the collection name is provided with the following configuration file:
```
{
    "collectionName": "FirestoreCollection"
}
```

If a configuration value is provided in both `firesert-config.json` and the environment, then the configuration file with take priority. For example, if the `PROJECT_ID` envronment varialbe has value "env-project-id" and the following `firesert-config.json` file is provided:
```
{
    "projectID": "config-project-id",
    "collectionName": "FirestoreCollection"
}
```
then the project ID will be "config-project-id".

## Running integration tests

To run integration tests in Docker against a local Firebase emulator, run the following:
- For typed documents: `make test-typed`
- For untyped documents: `make test-untyped`