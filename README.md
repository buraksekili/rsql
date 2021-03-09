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
![rsql](https://user-images.githubusercontent.com/32663655/110378071-376f6980-8066-11eb-8853-7a53d7014c68.gif)



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

    `$ cd ./cmd && go build -o $GOPATH/bin/rsql`

## Contribute

### Issues and Enhancement

- Feel free to post your issues and enhancement ideas on the `Issues` section.

### Contributing

1. *Fork* the repository
2. *Clone* the repository to your local environment.
3. *Create* your brach, e.g, `git branch -b fix-test`
4. *Commit* your changes `git commit -m "update test"` 
5. *Push* your branch `git push origin fix-test` 
6. *Create* a new  Pull Request.


### License
[License](https://github.com/buraksekili/rsql/blob/master/LICENSE)
