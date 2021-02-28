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

type ConnInfo struct {
	User     string
	Password string
	HostAddr string
	Port     string
	DbName   string
}

type DbClient struct {
	ci *ConnInfo
	l  *selog.Selog
}

func NewDbClient(l *selog.Selog) *DbClient {
	return &DbClient{&ConnInfo{}, l}
}

func main() {

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)
	dbClient := NewDbClient(l)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Host address: (127.0.0.1) ")
	hostAddr, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if hostAddr = strings.Trim(hostAddr, "\n"); hostAddr == "" {
		hostAddr = "127.0.0.1"
	}
	dbClient.ci.HostAddr = hostAddr

	fmt.Print("Port: (8080) ")
	port, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read host address: %v", err)
	}
	if port = strings.Trim(port, "\n"); port == "" {
		port = "8080"
	}
	dbClient.ci.Port = port

	fmt.Print("Database: (mysql) ")
	dbName, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read db name: %v", err)
	}
	if dbName = strings.Trim(dbName, "\n"); dbName == "" {
		dbName = "mysql"
	}
	dbClient.ci.DbName = dbName

	fmt.Print("User: (root)")
	user, err := reader.ReadString('\n')
	if err != nil {
		l.Fatal("cannot read user: %v", err)
	}
	if user = strings.Trim(user, "\n"); user == "" {
		user = "root"
	}
	dbClient.ci.User = user

	fmt.Print("Password: ")
	var password []byte
	for {
		password, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			l.Fatal("cannot read password: %v", err)
		}

		pw := string(password)
		if pw = strings.Trim(pw, "\n"); pw != "" {
			dbClient.ci.Password = pw
			fmt.Println()
			break
		}
		fmt.Print("\n\tInvalid password\nPassword: ")
	}

	dbClient.openConnection()
}

func (c *DbClient) openConnection() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.ci.User, c.ci.Password, c.ci.HostAddr, c.ci.Port, c.ci.DbName))

	if err != nil {
		c.l.Fatal("cannot open connection to db: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		c.l.Fatal("cannot establish connection: %v", err)
	}

}
