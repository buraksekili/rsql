# rsql

Easy way to walk through MySQL. 
rsql allows a quick check for MySQL tables, and it facilitates insertion operation for test purposes.

## Motivation

I am generally using MySQL through Docker. In order to check the result of my queries, I need to open a connection to MySQL container back and forth. 

rsql facilitates working on MySQL for the basic commands. It allows users to see available tables, the content of the tables, adding data to the table, etc.

## Usage

### Available Flags

```
rsql -h, --help                         : displays usage message and exits
rsql -f <FNAME>, --envfile <FNAME>      : reads environment file to establish MySQL connection.
```

### Commands

```
add <TABLE>     : adds data to <TABLE>
info <TABLE>    : displays the column informations of the <TABLE>
display <TABLE> : displays the data of the <TABLE>
tables          : displays available tables under the <DB> specified by user
help            : displays available commands after connection establishment.
q, exit         : exits the program
```

### Examples
![rsql](https://user-images.githubusercontent.com/32663655/110255037-e3e81780-7fa2-11eb-92b7-60b94a6d42f8.gif)



### [Environment File Usage](https://github.com/buraksekili/rsql/blob/master/env.list)

You can define the environment file while connecting the database. The program reads the environment file (case sensitive). Partial environment fields are OK. 

Following fields can be used. Missing fields are asked while executing the app.

```
USER=root
PASSWORD=yourpassword
ADDR=127.0.0.1
PORT=8080
DB=your_db_name
```

See [env.list](https://github.com/buraksekili/rsql/blob/master/env.list) for example usage

## Installation

```shell script
$ git clone https://github.com/buraksekili/rsql.git 
```

- If you have access to `bash`
    
    `$ bash install.sh`
- or;

    `$ cd ./cmd/cli && go build -o $GOPATH/bin/rsql`

### License
[License](https://github.com/buraksekili/rsql/blob/master/LICENSE)
