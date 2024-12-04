# Building a web server that creates a url shortener in Go

So, with my 30th day mark i'll be building a golang web application for a url shortener.

So, without further ado, lets dive right in.

## Steps for building the url shortener
1. Create a function to handle the unique ID generator from a string.
2. Create a function to handle the encoding of a unique ID into a base62 encoding.
3. Create a function that uses both function internally and returns a unique string
4. Create a database to store the unique ID, URL to be shorten and the shortened url
