# Building A URL Shortner in Go

So, i have been learning golang for 30 days and for my 30th day i decided to create a function handlers for handling shortening of links like a url shortener.

So, i had to learn about different concepts such as:
1. what are the ways of building a shortener
2. what is a good hash function?
3. How to create a good hash function.
4. base62 encoding
and so on.

This topics gave me some insights on how url shortening services works and today i focused on learning and understanding how to do the basics of hashing and finding a good hash function.

## Steps to take for building the hash function
For building this function, i have created a layout to follow to build a good url shortener.
URL SHORTNER IN GO
1. Get the url
2. check if the url exists in db
3. if it doesnt then generate a unique ID
4. convert ID to base64
5. store both URL && ID
6. Return the shortened URL.

## Whats Next?
I have been able to map out the functions needed for creating the url shortener, the next step is to create a web server that would be responsible for hashing and also adding a frontend using basic html and css. 


## Resources
1. [How to build a url shortener](https://www.linkedin.com/pulse/building-url-shortener-using-hash-functions-base62-conversion-singh-y01oc/)
2. [How to create a hash function in go](https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go)
3. [Hash Tables and hash functions - CS50 lecture](https://youtu.be/nvzVHwrrub0?si=q9KpWg9DZ6GbiY24)
