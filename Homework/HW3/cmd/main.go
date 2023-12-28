package main

import (
	"database/sql"
	"fmt"
	"tiket/repository"
	"tiket/ticket"
	"tiket/user"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

var (
	User   int = 1
	Ticket int = 2
	Report int = 3
)

func main() {
	db, err := connectDb()
	if err != nil {
		fmt.Println("error while opening db file : ", err)
		return
	}

	defer db.Close()

	repo := repository.Repository{Db: db}

	fmt.Println("1 => User")
	fmt.Println("2 => Ticket")
	fmt.Println("3 => Report")
	fmt.Println()
	var numberFromInput int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&numberFromInput)

	switch numberFromInput {
	case User:
		handleUserOperations(repo)
	case Ticket:
		handleTicketOperations(repo)
	case Report:
		handleReportOperation(repo)
	default:
		fmt.Println("Invalid option selected!")
	}
}

func handleUserOperations(repo repository.Repository) {
	fmt.Println("1 => User add")
	fmt.Println("2 => Get user by id")
	fmt.Println("3 => List of users")
	fmt.Println("4 => Update user by id")
	fmt.Println("5 => Delete user")
	fmt.Println()
	fmt.Print("Enter your choice: ")
	var inputForUser int
	fmt.Scan(&inputForUser)
	switch inputForUser {
	case 1:
		u := user.User{}
		if err := repo.AddUser(u.NewUser()); err != nil {
			fmt.Println("error while adding user : ", err)
			return
		}
		fmt.Println("User successfully added to db!")
	case 2:
		var id string
		fmt.Print("enter user id : ")
		fmt.Scan(&id)
		u, err := repo.GetUserById(uuid.MustParse(id))
		if err != nil {
			fmt.Println("error while user by id : ", err)
			return
		}
		fmt.Println(u)
	case 3:
		u, err := repo.GetAllUsers()
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(u)
	case 4:
		var id string
		var firstName string
		var lastName string
		var email string
		fmt.Print("id : ")
		fmt.Scan(&id)
		fmt.Print("First name : ")
		fmt.Scan(&firstName)
		fmt.Print("Last name : ")
		fmt.Scan(&lastName)
		fmt.Print("email : ")
		fmt.Scan(&email)

		u := user.User{
			Id:        uuid.MustParse(id),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		}
		if err := repo.UpdateUsers(u); err != nil {
			panic(err)
			return
		}
		fmt.Println("user updated ! ")
	case 5:
		var id string
		fmt.Print("id : ")
		fmt.Scan(&id)
		if err := repo.DeleteUser(uuid.MustParse(id)); err != nil {
			panic(err)
			return
		}
		fmt.Println("User deleted !")
	default:
		fmt.Println("Invalid option selected!")
	}
}

func handleTicketOperations(repo repository.Repository) {
	fmt.Println("1 => Ticket add")
	fmt.Println("2 => Get ticket by id")
	fmt.Println("3 => List of tickets")
	fmt.Println("4 => Update ticket by id")
	fmt.Println("5 => Delete ticket")
	fmt.Println()
	fmt.Print("Enter your choice: ")
	var inputForUser int
	fmt.Scan(&inputForUser)
	switch inputForUser {
	case 1:
		t := ticket.Ticket{}
		if err := repo.AddTicket(t.NewTicket()); err != nil {
			fmt.Println("error while adding ticket : ", err)
			return
		}
		fmt.Println("Ticket successfully added to db!")
	case 2:
		var id string
		fmt.Print("enter ticket id : ")
		fmt.Scan(&id)
		t, err := repo.GetTicketById(uuid.MustParse(id))
		if err != nil {
			fmt.Println("error while ticket by id : ", err)
			return
		}
		fmt.Println(t)
	case 3:
		t, err := repo.GetAllTickets()
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(t)
	case 4:
		t := ticket.Ticket{}
		if err := repo.UpdateTicket(t.NewTicket()); err != nil {
			panic(err)
			return
		}
	case 5:
		fmt.Print("id : ")
		var id string
		fmt.Scan(&id)
		if err := repo.DeleteTicket(uuid.MustParse(id)); err != nil {
			panic(err)
			return
		}
		fmt.Println("Ticket deleted !")
	default:
		fmt.Println("Invalid option selected!")
	}
}

func handleReportOperation(repo repository.Repository) {
	fmt.Print("from : ")
	var from string
	fmt.Scan(&from)
	fmt.Print("to : ")
	var to string
	fmt.Scan(&to)
	report, err := repo.ReportAll(from, to)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(report)
}

func connectDb() (*sql.DB, error) {
	connactionString := "host=localhost port=5432 user=myuser password=0 dbname=mydatabase sslmode=require"
	db, err := sql.Open("postgres", connactionString)
	if err != nil {
		fmt.Println("error while opening db file : ", err)
		return nil, err
	}
	return db, nil
}
