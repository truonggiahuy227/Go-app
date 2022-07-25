package repository

import (
	"database/sql"
	"fmt"

	"akawork.io/dto"
	"akawork.io/infrastructure/logger"
	"akawork.io/infrastructure/util"
	_ "github.com/lib/pq"
)

type PostgreSqlRepository struct {
	DbContext *sql.DB
}

type PostgreSqlRowScanner interface {
	Scan(dest ...interface{}) error
}

/**
 * Sets Postgresql Db context
 */
func (repository *PostgreSqlRepository) SetDbContext(dbContext *sql.DB) {
	repository.DbContext = dbContext
}

/**
 * Handles MySql error
 */
func (respository *PostgreSqlRepository) HandleError(err error) {
	logger.Error("[PostgreSql]", err.Error())
}

/**
 * Initializes a MySql infrastructure
 */
func ConnectPostgreSql(host string, userName string, password string, database string, maxOpenConnections int, maxIdleConnections int) (db *sql.DB) {
	db, err := sql.Open("postgres", "postgres://"+userName+":"+password+"@"+host+":5432/"+database+"?sslmode=disable")
	if err != nil {
		logger.Error("Failed to connect to PostgreSql", err.Error())
		return nil
	}
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	return db
}

func InsertAccount(db *sql.DB, user *dto.AccountDto) error {

	query := `INSERT INTO accounts (Username, Password) VALUES (?, ?)`
	_, err := db.Exec(query, user.Username, user.Password)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func GetByUsername(db *sql.DB, pUsername string) (*dto.AccountDto, error) {
	user := dto.AccountDto{}

	var username, password sql.NullString

	err := db.QueryRow(`SELECT Username, Password FROM sso_user WHERE Username = ? 
                                  		LIMIT 1`, pUsername).Scan(&username, &password)

	if err != nil {
		return nil, err
	}

	user.Username = util.ConvertNullString(username)
	user.Password = util.ConvertNullString(password)

	return &user, nil
}
