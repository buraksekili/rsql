# rsql

Easy way to walk through MySQL. 
rsql allows a quick check for MySQL tables, and it facilitates insertion operation for test purposes.

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
q, exit         : exits the program
```

### Examples
![rsql](https://user-images.githubusercontent.com/32663655/109490848-df4ecb00-7a99-11eb-8b32-1434cbe7b626.png)

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
