# More, More Web Server Stuffs

I am trying to understand the low-level case to why using the w and r on the func and i found out (with GPT help) that the w (http.ResponseWriter) is responsible for handling the out-going response in which the web server would respond for a given request.

r (http.Request) is used for understanding the type of reqest that is being received by the web server.

## Assignment
1. Basic Task:

1. Create a web server with the following routes:
- / (responds with "Welcome!")
- /about (responds with "About Page")
- /search (takes a query parameter q and responds with "Searching for: {q}")
- /user/{id} (extracts {id} from the URL and responds with "User ID: {id}").

2. Challenge Task:

b. Add error handling for routes:
    If /search is accessed without q, respond with "Query parameter is missing."
    If /user/{id} is accessed with an invalid ID (e.g., empty or less than 3 characters), respond with "Invalid User ID."

## Solution
[server/](https://github.com/Lawsonredeye/100DaysOfGoLang/tree/main/day15/server)
