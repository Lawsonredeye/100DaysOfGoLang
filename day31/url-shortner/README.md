# URL Shortener in Go

This is my complete web server which takes in a link passed via JSON and then returns a short base62 encoded string which would be used to retrieve the original link from the database.

Built with Go net/http standard library, this opens a new chapter on how Go can be used for building a simple url service.

## Requirements
- Install MySQL
- MySQL user
- CURL or Postman for running test
- Browser for testing the shortened links

## How to use

To use the url shortener all you need to do is to follow these steps:
1. add your MySQL db username and password on the env using the `export` keyword:
    `export db_user="DB_USER_NAME" && export db_pwd="DB_PASSWORD"`
2. run the command `go run main.go`
3. open a new terminal and send a **POST** request with the key name `url` and the value as the url you want to shorten
```Bash
    curl -i -H "Content-Type: application/json" \
        -X POST localhost:8000/create \
        -d '{"url": "https://github.com/Lawsonredeye/100DaysOfGoLang/blob/main/day27/main.go"}'
```
Note: You can also use postman for this if you don't have cUrl installed locally. 
4. store the returned value some where
5. to use the returned shortened string open your browser and send a get request to the `http://localhost:8000/` together with the shortened value you stored
6. voila, thats how you use the local url shortener.

## Future update
I plan on implementing a basic frontend for this project instead of running it with either curl or postman.

## Stacks Used in this project
1. Go
2. MySQL
