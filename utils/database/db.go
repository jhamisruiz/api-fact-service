package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL
)

// InitDB inicializa y devuelve una conexión a la base de datos
func InitDB(config map[string]string) (*sql.DB, error) {
	// Construye el Data Source Name (DSN) para conectar con MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config["user"], config["password"], config["host"], config["port"], config["dbname"])

	// Abre la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica que la conexión sea exitosa
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
