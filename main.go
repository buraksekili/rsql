package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/buraksekili/rsql/cli"
	"github.com/buraksekili/rsql/commands"
	"github.com/buraksekili/rsql/data"
	"github.com/buraksekili/selog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connInfo := &data.ConnInfo{
		User:     os.Getenv("R_MYSQL_USER"),
		Password: os.Getenv("R_MYSQL_PASSWORD"),
		HostAddr: os.Getenv("R_MYSQL_ADDR"),
		Port:     os.Getenv("R_MYSQL_PORT"),
		DbName:   os.Getenv("R_MYSQL_DB"),
	}

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)
	dbClient := commands.NewDbClient(l)

	switch v := cli.ParseFlag(os.Args[1:]).(type) {
	case cli.HelpOp:
		cli.PrintHelp(os.Stdout)
		return
	case cli.ConnInfoInput:
		getConnInfo(dbClient, connInfo)
	case cli.EnvFileOp:
		getConnInfo(dbClient, connInfo)
	case cli.UnknownOp:
		dbClient.Log.Fatal(fmt.Sprintf("Unknown operation: : %s\n", v.Error))
	case cli.InvalidOp:
		dbClient.Log.Fatal(fmt.Sprintf("Invalid operation: %s\n", v.Error))
	default:
		dbClient.Log.Fatal(fmt.Sprintf("invalid operation: %T", v))
	}
}

// getConnInfo takes required inputs to establish MySQL connection from terminal.
// Then, it opens connection to the DB. In case of any error, it prints fatal message
// and exits the program.
func getConnInfo(dbClient *commands.DbClient, connInfo *data.ConnInfo) {
	reader := bufio.NewReader(os.Stdin)

	dbClient.ConnInfo = connInfo
	if strings.Trim(dbClient.ConnInfo.HostAddr, " ") == "" {
		fmt.Print("Host address: (127.0.0.1) ")
		hostAddr, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read host address: %v", err)
		}
		if hostAddr = strings.Trim(hostAddr, "\r\n"); hostAddr == "" {
			hostAddr = "127.0.0.1"
		}
		dbClient.ConnInfo.HostAddr = hostAddr
	}

	if strings.Trim(dbClient.ConnInfo.Port, " ") == "" {
		fmt.Print("Port: (8080) ")
		port, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read host address: %v", err)
		}
		if port = strings.Trim(port, "\r\n"); port == "" {
			port = "8080"
		}
		dbClient.ConnInfo.Port = port
	}

	if strings.Trim(dbClient.ConnInfo.DbName, " ") == "" {
		fmt.Print("Database: (mysql) ")
		dbName, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read db name: %v", err)
		}
		if dbName = strings.Trim(dbName, "\r\n"); dbName == "" {
			dbName = "mysql"
		}
		dbClient.ConnInfo.DbName = dbName
	}

	if strings.Trim(dbClient.ConnInfo.User, " ") == "" {
		fmt.Print("User: (root) ")
		user, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read user: %v", err)
		}
		if user = strings.Trim(user, "\r\n"); user == "" {
			user = "root"
		}
		dbClient.ConnInfo.User = user
	}

	if strings.Trim(dbClient.ConnInfo.Password, " ") == "" {
		fmt.Print("Password: ")
		for {
			var err error
			var password []byte
			if runtime.GOOS == "windows" {
				password, err = term.ReadPassword(int(syscall.Stdin))
			} else {
				password, err = terminal.ReadPassword(int(syscall.Stdin))
			}
			if err != nil {
				dbClient.Log.Fatal("cannot read password: %v", err)
			}

			pw := string(password)
			dbClient.ConnInfo.Password = pw
			break
		}
	}

	if err := dbClient.OpenConnection(); err != nil {
		dbClient.Log.Fatal(fmt.Sprintf("cannot open connection: %v\n", err))
	}
}
