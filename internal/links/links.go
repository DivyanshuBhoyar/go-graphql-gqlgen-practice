package links

import (
	database "github.com/DivyanshuBhoyar/gqlgen-prac/internal/pkg/db"
	"github.com/DivyanshuBhoyar/gqlgen-prac/internal/users"
	"log"
)

// struct for database
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	// prepeare sql command
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	// execute sql command
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	// get id of last inset
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	log.Println("New entry inserted")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []Link  // db link[]
	var username string
	var id string

	for rows.Next() {
		var link Link   // db link struct
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username) // changed
		if err != nil{
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
