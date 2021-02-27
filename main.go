package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/buraksekili/selog"
)

func main() {

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Host address: (127.0.0.1)")
	hostAddr, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if strings.Trim(hostAddr, "\n") == "" {
		hostAddr = "127.0.0.1"
	}

	fmt.Print("Port: (8080)")
	port, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if strings.Trim(port, "\n") == "" {
		port = "8080"
	}

	fmt.Print("Database: (mysql)")
	dbName, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read db name: %v", err)
	}
	if strings.Trim(dbName, "\n") == "" {
		dbName = "mysql"
	}

	fmt.Print("User: (root)")
	user, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read user: %v", err)
	}
	if strings.Trim(user, "\n") == "" {
		user = "root"
	}

	fmt.Print("Password: ")
	var password []byte
	for {
		password, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			l.Fatal("cannot read password: %v", err)
		}

		pw := string(password)

		if strings.Trim(pw, "\n") != "" {
			fmt.Println()
			break
		}
		fmt.Printf("\n\tInvalid password %s\nPassword: ", pw)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s\n", user, password, hostAddr, port, dbName))
	if err != nil {
		l.Fatal("cannot open connection to db: %v", err)
	}

	defer db.Close()

	fmt.Fprintf(os.Stdout, "%s:%s@tcp(%s:%s)/%s\n", user, password, hostAddr, port, dbName)
}
