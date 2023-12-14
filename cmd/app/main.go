package main

import (
	"Jepix/internal/database"
)

func main() {
	Database.CheackHealth()
	Database.Disconnect()
}
