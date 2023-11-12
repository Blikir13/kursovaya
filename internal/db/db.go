package db

import(
	"database/sql"
	"fmt"
	"crypto/md5"
    "encoding/hex"
	"time"

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

func (d *db) Write (name string, port int, port_state string, bool_change bool) error {

	_, err := d.base.Exec("insert into logs (device_name, port, port_state, bool_change, date_time) values ($1, $2, $3, $4, $5)", name, port, port_state, bool_change, time.Now())
	return err
}

func (d *db) GetTable () *sql.Rows {
	rows, _ := d.base.Query("select * from logs")
	fmt.Println("error: ", rows)
	return rows
}