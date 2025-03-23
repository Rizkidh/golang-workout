package cmd

import (
	"context"
	"golang-crud-clean-arch/config"
	"golang-crud-clean-arch/database"
)

func main() {
	client := database.dbConn()                   // Panggil fungsi yang benar
	defer client.Disconnect(context.Background()) // Gunakan context.Background()
}
