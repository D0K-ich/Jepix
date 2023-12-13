package Database

import (
	cfg "Jepix/internal/Config"
	models "Jepix/internal/Models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dsn = cfg.GetConfDB()
var database, err = sqlx.Connect("postgres", dsn)

func СheckHealth() string {
	if err != nil {
		Disconnect()
		return (fmt.Sprintf("error %s", err))
	}
	return "Database is healthy"
}

func Disconnect() {

	database.Close()
}

func AddNewUser(Name string, Surname string) {
	createDatabaseCommand := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name text,
		surname text,
		email text,
		dateofregister int,	
		balancemoney int,
		balancebonus int,
		promocodes int
	)
	`

	database.Queryx(createDatabaseCommand)

	persst := models.Customer{
		Name: Name, Surname: Surname, Email: "test", Dateofregister: 1, Balancemoney: 1,
		Balancebonus: 0, Promocodes: 0}
	_, err = database.NamedExec(`INSERT INTO users (name, surname, email, dateofregister, balancemoney, balancebonus, promocodes)
        VALUES (:name, :surname, :email, :dateofregister, :balancemoney, :balancebonus, :promocodes)`, persst)

	t := models.Customer{}
	database.Get(&t, "SELECT name, surname FROM users")

	fmt.Println(t)
}
