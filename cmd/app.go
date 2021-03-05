package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/buraksekili/rsql/cmd/cli"

	"github.com/buraksekili/rsql/data"

	"github.com/buraksekili/rsql/cmd/client"
	"github.com/buraksekili/selog"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)

	dbClient := client.NewDbClient(l)

	c := &data.ConnInfo{}
	switch v := cli.ParseFlag(os.Args[1:]).(type) {
	case cli.HelpOp:
		cli.PrintHelp(os.Stdout)
		return
	case cli.EnvFileOp:
		connInfo, err := cli.ReadEnvFile(v.Filename)
		if err != nil {
			dbClient.Log.Fatal("cannot read env file: %v", err)
		}
		c = connInfo
		fmt.Println("FILE LOADED")
		getConnInfo(dbClient, c)
	case cli.UnknownOp:
		dbClient.Log.Fatal("Unknown operation: ", v.Error)
	case cli.InvalidOp:
		dbClient.Log.Fatal("Invalid operation: ", v.Error)
	case cli.ConnInfoInput:
		getConnInfo(dbClient, c)
	default:
		dbClient.Log.Fatal("invalid operation: %T", v)
	}
}

// getConnInfo takes required inputs to establish MySQL connection from terminal.
// Then, it opens connection to the DB. In case of any error, it prints fatal message
// and exits the program.
func getConnInfo(dbClient *client.DbClient, connInfo *data.ConnInfo) {
	reader := bufio.NewReader(os.Stdin)

	dbClient.ConnInfo = connInfo
	if strings.Trim(dbClient.ConnInfo.HostAddr, " ") == "" {
		fmt.Print("Host address: (127.0.0.1) ")
		hostAddr, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read host address: %v", err)
		}
		if hostAddr = strings.Trim(hostAddr, "\n"); hostAddr == "" {
			hostAddr = "127.0.0.1"
		}
		dbClient.ConnInfo.HostAddr = hostAddr
	}

	if strings.Trim(dbClient.ConnInfo.Port, "") == "" {
		fmt.Print("Port: (8080) ")
		port, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read host address: %v", err)
		}
		if port = strings.Trim(port, "\n"); port == "" {
			port = "8080"
		}
		dbClient.ConnInfo.Port = port
	}

	if strings.Trim(dbClient.ConnInfo.DbName, "") == "" {
		fmt.Print("Database: (mysql) ")
		dbName, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read db name: %v", err)
		}
		if dbName = strings.Trim(dbName, "\n"); dbName == "" {
			dbName = "mysql"
		}
		dbClient.ConnInfo.DbName = dbName
	}

	if strings.Trim(dbClient.ConnInfo.User, "") == "" {
		fmt.Print("User: (root)")
		user, err := reader.ReadString('\n')
		if err != nil {
			dbClient.Log.Fatal("cannot read user: %v", err)
		}
		if user = strings.Trim(user, "\n"); user == "" {
			user = "root"
		}
		dbClient.ConnInfo.User = user
	}

	if strings.Trim(dbClient.ConnInfo.Password, "") == "" {
		fmt.Print("Password: ")
		for {
			password, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				dbClient.Log.Fatal("cannot read password: %v", err)
			}

			pw := string(password)
			dbClient.ConnInfo.Password = pw
			break
		}
	}
	if err := dbClient.OpenConnection(); err != nil {
		dbClient.Log.Fatal("cannot open connection: %v", err)
	}
}
