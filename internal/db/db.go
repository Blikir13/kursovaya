package db

import(
	"database/sql"
	"fmt"
	"crypto/md5"
    "encoding/hex"

	_ "github.com/lib/pq"
)


type db struct {
	base *sql.DB
}

func NewDB() *db {
	return &db{
		base: Connection(),
	}
}

func Connection() *sql.DB {
	connection := "user=postgres password=1234 dbname=trps sslmode=disable"
	database, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	//defer database.Close()
	return database
}

func (d *db) Login (login, Password string) (error) {
	hash := md5.Sum([]byte(Password))
	hashedPassword := hex.EncodeToString(hash[:])
	var row string
	err := d.base.QueryRow(`select id from users where login = $1 and hashed_password = $2`, login, hashedPassword).Scan(&row)

	// fmt.Println(a)
	if err != nil {
		if err == sql.ErrNoRows {
            return fmt.Errorf("canPurchase unknown album")
        }
		return fmt.Errorf("error", login, Password, err)
	}
	return nil
}