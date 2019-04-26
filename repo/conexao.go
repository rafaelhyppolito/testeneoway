package repo

import(
    "fmt"
    "database/sql"
	_ "github.com/lib/pq"
)

// Retornar uma forma de conexão com o banco
func Connect() *sql.DB {
	driverConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "postgres", "root", "localhost", "5432", "base")
	connection, err := sql.Open("postgres", driverConfig)

    if err != nil {
        fmt.Printf("database.Connect ERROR: %s", err)
	}
	
	sqlStatement := "INSERT INTO base(cpf,priv,incompleto) VALUES('01249885043',0,0)"
	_, err = connection.Exec(sqlStatement)
	if err != nil {
		  panic(err)
	}

    return connection
}