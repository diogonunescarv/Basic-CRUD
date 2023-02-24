package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar faz a conexão com o Banco de Dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/teste1?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	// confere se a conexão com o banco de dados obteve exito, se não retorna um erro e um valor nulo para DB
	if erro != nil {
		return nil, erro
	}

	// confere se a conexão com o banco de dados obteve exito, se não retorna um erro e um valor nulo para DB
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
