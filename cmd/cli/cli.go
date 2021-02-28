package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/buraksekili/rsql/data"

	"github.com/buraksekili/rsql/cmd/client"
	"github.com/buraksekili/selog"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	c := &data.ConnInfo{}
	switch v := parseFlag().(type) {
	case HelpOp:
		printHelp(os.Stdout)
		return
	case EnvFileOp:
		connInfo, err := readEnvFile(v.Filename)
		if err != nil {
			log.Fatal("cannot read env file: %v", err)
		}
		c = connInfo
	}

	logger := log.New(os.Stdout, "rsql ", log.LstdFlags|log.Lshortfile)
	l := selog.NewLogger(logger)

	dbClient := client.NewDbClient(l)

	if c != nil {
		fmt.Println("FILE LOADED")
	}

	getConnInfo(dbClient, c)
}

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
			if pw = strings.Trim(pw, "\n"); pw != "" {
				dbClient.ConnInfo.Password = pw
				fmt.Println()
				break
			}
			fmt.Print("\n\tInvalid password\nPassword: ")
		}
	}
	dbClient.OpenConnection()
}
