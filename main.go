package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Client.
// Теперь, если передать объект Client в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (c Client) String() string {
	return fmt.Sprintf("ID: %d FIO: %s Login: %s Birthday: %s Email: %s",
		c.ID, c.FIO, c.Login, c.Birthday, c.Email)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// добавление нового клиента
	newClient := Client{
		FIO:      "Щербаков Владлен Александрович", // укажите ФИО
		Login:    "blackenkeeper",                  // укажите логин
		Birthday: "19980723",                       // укажите день рождения
		Email:    "test@test.ru",                   // укажите почту
	}

	id, err := insertClient(db, newClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err := selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// обновление логина клиента
	newLogin := "deathlaughfear" // укажите новый логин
	err = updateClientLogin(db, newLogin, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// удаление клиента
	err = deleteClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	_, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertClient(db *sql.DB, client Client) (int64, error) {
	insertQuery := "insert into clients (fio, login, birthday, email) values (?, ?, ?, ?)"
	res, err := db.Exec(insertQuery, client.FIO, client.Login, client.Birthday, client.Email)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return res.LastInsertId()
}

func updateClientLogin(db *sql.DB, login string, id int64) error {
	updateQuery := "update clients set login = :login where id = :id"
	_, err := db.Exec(updateQuery, sql.Named("login", login), sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func deleteClient(db *sql.DB, id int64) error {
	deleteQuery := "delete from clients where id = :id"
	_, err := db.Exec(deleteQuery, sql.Named("id", id))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func selectClient(db *sql.DB, id int64) (Client, error) {
	client := Client{}

	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)

	return client, err
}
