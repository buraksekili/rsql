# rsql

Easy way to walk through MySQL. 
rsql allows a quick check for MySQL tables, and it facilitates insertion operation for test purposes.

## Motivation

I am generally using MySQL through Docker. In order to check the result of my queries, I need to open a connection to MySQL container back and forth. 

rsql facilitates working on MySQL for the basic commands. It allows users to see available tables, the content of the tables, adding data to the table, etc.

## Usage
```
add     <TABLE> : adds data to <TABLE>
info    <TABLE> : displays the column informations of the <TABLE>
display <TABLE> : displays the data of the <TABLE>
tables          : displays available tables under the <DB> specified by user
help            : displays available commands after connection establishment.
q, exit         : exits the program
```

![rsql](https://user-images.githubusercontent.com/32663655/110378071-376f6980-8066-11eb-8853-7a53d7014c68.gif)

```bash
burak@burak-ZenBook:~$ rsql

===== STATS =====
Max Open Connections:	 0
Open Connection:	 1
Idle:			 1
In Use:			 0
COMMANDS
	add <TABLE>	: adds data to <TABLE>
	info <TABLE>	: displays the column informations of the <TABLE>
	display <TABLE>	: displays the data of the <TABLE>
	tables		: displays available tables under the <DB> specified by user
	help		: displays this message
	q, exit 	: exits the program
rsql> tables
+-------+
| posts |
+-------+
rsql> info posts

FETCHING INFORMATION FOR TABLE: posts
+------------+--------------+------+-----+---------+----------------+
|   FIELD    |     TYPE     | NULL | KEY | DEFAULT |     EXTRA      |
+------------+--------------+------+-----+---------+----------------+
| post_id    | int          | NO   |     |         | auto_increment |
| post_title | varchar(100) | NO   |     |         |                |
| post_body  | text         | NO   |     |         |                |
+------------+--------------+------+-----+---------+----------------+
rsql>  
```

### [Environment File Usage](https://github.com/buraksekili/rsql/blob/master/.env)

You can define the environment file while connecting the database. 
`rsql` reads `.env` file by default.

Following fields are required for MySQL connection. Missing fields will be asked while executing the app.
```bash
R_MYSQL_USER=root
R_MYSQL_ADDR=127.0.0.1
R_MSQL_PORT=8080
R_MYSQL_DB=posts
```

You can use any `.env` file by passing `-f` or `--envfile` flag.
```bash
$ rsql --envfile ~/your/.env
```

## Installation
```bash
$ git clone https://github.com/buraksekili/rsql.git 
$ cd rsql
$ make build
```

### Run Tests
```bash
$ make test
```

## Contribute

### Issues and Enhancement

- Feel free to post your issues and enhancement ideas on the `Issues` section.

### Contributing

1. *Fork* the repository
2. *Clone* the repository to your local environment.
3. *Create* your brach, e.g, `git branch -b fix-test`
4. *Commit* your changes `git commit -s -m "update test"` 
5. *Push* your branch `git push origin fix-test` 
6. *Create* a new  Pull Request.

> *Do not forget to sign your commits, you can sign your commits with `git commit -s`.*

### License
[License](https://github.com/buraksekili/rsql/blob/master/LICENSE)
