package Database

//todo Сделать checkheals, disconnect, init и ping
//todo Перенести файл базы в другую папку

import (
	//cfg "Jepix/internal/Config"
	//models "Jepix/internal/Models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Disconnect() {
	var database, err = sqlx.Connect("sqlite3", "check")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("ok")
	q1, _ := database.Prepare(`
	CREATE TABLE IF NOT EXISTS users (id INTEGER, name TEXT)
`)
	q1.Exec()

	q2, _ := database.Prepare(`
	SELECT id FROM users
`)
	q2.Exec()
	database.Close()
}
