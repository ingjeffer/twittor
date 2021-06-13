package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN */
var MongoCN = conectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://root:root@cluster0.z1ux7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD */
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

/* ChequeoConnection */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
