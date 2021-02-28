package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/buraksekili/rsql/cmd/client"
	"github.com/buraksekili/selog"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)
	dbClient := client.NewDbClient(l)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Host address: (127.0.0.1) ")
	hostAddr, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if hostAddr = strings.Trim(hostAddr, "\n"); hostAddr == "" {
		hostAddr = "127.0.0.1"
	}
	dbClient.ConnInfo.HostAddr = hostAddr

	fmt.Print("Port: (8080) ")
	port, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if port = strings.Trim(port, "\n"); port == "" {
		port = "8080"
	}
	dbClient.ConnInfo.Port = port

	fmt.Print("Database: (mysql) ")
	dbName, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read db name: %v", err)
	}
	if dbName = strings.Trim(dbName, "\n"); dbName == "" {
		dbName = "mysql"
	}
	dbClient.ConnInfo.DbName = dbName

	fmt.Print("User: (root)")
	user, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read user: %v", err)
	}
	if user = strings.Trim(user, "\n"); user == "" {
		user = "root"
	}
	dbClient.ConnInfo.User = user

	fmt.Print("Password: ")
	var password []byte
	for {
		password, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			l.Fatal("cannot read password: %v", err)
		}

		pw := string(password)
		if pw = strings.Trim(pw, "\n"); pw != "" {
			dbClient.ConnInfo.Password = pw
			fmt.Println()
			break
		}
		fmt.Print("\n\tInvalid password\nPassword: ")
	}

	dbClient.OpenConnection()
}
