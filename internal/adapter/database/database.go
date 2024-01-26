// internal/adapter/database/database.go
package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db               *sql.DB
	once             sync.Once
	connectionString = "postgres://root:secret@localhost:5432/bank_db?sslmode=disable"
)

func GetDBConnection() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			// Manejar error de conexión a la base de datos.
			fmt.Println("Error al conectar a la base de datos:", err)
		}

		// Realizar una conexión de prueba para asegurarse de que todo está configurado correctamente.
		err = db.Ping()
		if err != nil {
			// Manejar error de ping a la base de datos.
			fmt.Println("Error al hacer ping a la base de datos:", err)
		}
	})

	return db
}
