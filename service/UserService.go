package service

import (
	"github.com/ergildo/go-lang-app-mysql-example/database"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id   int64
	Name string
}

func ListAll() []User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select Id, Name from users")
	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}
	return users
}

func ListById(id int64) User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	var u User
	db.QueryRow("select Id, Name from users where Id=?", id).Scan(&u.Id, &u.Name)

	return u
}

func Save(user User) User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	stm, err := db.Prepare("insert into users(Name) values(?)")
	rs, err := stm.Exec(user.Name)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u := ListById(lastId)
	return u
}

func Update(user User) User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	stm, err := db.Prepare("update users set name =? where id =?")
	rs, err := stm.Exec(user.Name, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u := ListById(lastId)
	return u
}

func Delete(id int64) {

	db, err := database.GetDB()
	defer database.CloseDB()

	stm, err := db.Prepare("delete from users where id =?")
	_, err = stm.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

}
