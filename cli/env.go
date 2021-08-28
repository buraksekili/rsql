package cli

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/buraksekili/rsql/data"
)

func ReadEnvFile(filename string) (*data.ConnInfo, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot get working directory: %v", err)
	}
	fp := path.Join(wd, filename)

	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("cannot open env file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	conn := data.ConnInfo{}
	for scanner.Scan() {

		line := scanner.Text()

		i := strings.Index(line, "=")
		key := line[:i]
		value := line[i+1:]

		switch key {
		case "USER":
			conn.User = value
		case "ADDR":
			conn.HostAddr = value
		case "PASSWORD":
			conn.Password = value
		case "DB":
			conn.DbName = value
		case "PORT":
			conn.Port = value
		}
	}
	return &conn, nil
}
