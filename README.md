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