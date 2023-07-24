package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	dbConnection := getDBConnection()
	userRepo := db.NewUserRepositoryDB(dbConnection)
	paymentRepo := db.NewPaymentRepositoryDB(dbConnection)
	pagarmeGateway := pagarme.NewPagarMeGateway()

	createUser := usecases.NewCreateUser(userRepo)
	processPayment := usecases.NewProcessPayment(paymentRepo, pagarmeGateway)

	server := http.NewServer(createCompany, listUsers, processPayment, getPayment)
	server.SetupRoutes()
	server.Run(os.Getenv("PORT"))

}

func getDBConnection() *sql.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
