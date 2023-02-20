package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/lucasszmt/transfeera-challenge/domain/vo"
	"github.com/lucasszmt/transfeera-challenge/infra/db"
	"github.com/lucasszmt/transfeera-challenge/infra/log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	logger := log.PrettyLogger()
	// Init DB connection
	dbConn := db.Must(db.NewPostgresConn(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	initialMigration(dbConn)

	// Init repositories
	receiverRepo := db.NewReceiver(dbConn)

	// Init services
	receiverService := receiver.NewService(&logger, receiverRepo)
	GenerateData(receiverService)
}

func GenerateData(receiverService *receiver.Service) {
	for i := 0; i < 30; i++ {
		item := dtos.CreateReceiverRequest{
			Name:       gofakeit.Name(),
			Email:      gofakeit.Email(),
			Doc:        gofakeit.Regex(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`),
			PixKeyType: vo.PixKeyType(gofakeit.RandomString([]string{"cpf", "cnpj", "random_key", "email", "phone"})),
		}

		if gofakeit.RandomString([]string{"cpf", "cnpj"}) == "cpf" {
			item.Doc = gofakeit.Regex(vo.CPFRegexp.String())
		} else {
			item.Doc = gofakeit.Regex(vo.CNPJRegexp.String())
		}

		switch string(item.PixKeyType) {
		case "cpf":
			item.PixKey = gofakeit.Regex(vo.CPFRegexp.String())
		case "cnpj":
			item.PixKey = gofakeit.Regex(vo.CNPJRegexp.String())
		case "random_key":
			item.PixKey = gofakeit.Regex(vo.RandomKeyRegexp.String())
		case "email":
			item.PixKey = gofakeit.Regex(item.Email)
		case "phone":
			item.PixKey = gofakeit.Regex(vo.PhoneRegexp.String())
		}
		receiverService.CreateReceiver(item)
	}
}

func initialMigration(db *sqlx.DB) {
	tx := db.MustBegin()
	tx.MustExec(CreateExtension)
	tx.MustExec(CreateTablePixKeyType)
	tx.MustExec(InsertPixKeyItems)
	tx.MustExec(CreateReceiverTable)
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
