package repo

import(
    "fmt"
    "database/sql"
	_ "github.com/lib/pq"
)

// Retornar uma forma de conexão com o banco
func Connect() *sql.DB {
//	var user = "postgres"
//	var pass = "root"
//	var base = "base"
//	var host = "localhost"
//	var port = 5432
//	var base = "base"
//	var sslm = "enabled"

	//driverConfig := fmt.Sprintf("server=%s;user=%s;password=%s;port=%d;database=%s;sslmode=%s", host, user, pass, port, base, sslm)
	driverConfig := "user=postgres password=root port=5432 database=postgres sslmode=disable"
	connection, err := sql.Open("postgres", driverConfig)

    if err != nil {
        fmt.Printf("database.Connect ERROR: %s\r\n", err)
	}

	err = connection.Ping()

	if err != nil {
        fmt.Printf("Deu ruim: %s", err)
	}

    return connection
}

func ExecSQL(sql string, connection *sql.DB) {
	sqlStatement := sql
	_, err := connection.Exec(sqlStatement)
	if err != nil {
		  panic(err)
	}	
}