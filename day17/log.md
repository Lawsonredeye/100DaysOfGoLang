# Basic Authentication on web servers - Day 17

Authentication is a basic && important part of backend implementation and using the standard library we can make use of net/http [BasicAuth](https://pkg.go.dev/net/http#Request.BasicAuth) for adding security to our simple web server application

## Things to note when using Basic Authentication
To help keep things secure you should:

1. Only ever use it over HTTPS connections. If you don't use HTTPS, the Authorization header can potentially be intercepted and decoded by an attacker, who can then use the username and password to gain access to your protected resources.

2. Use a strong password that is difficult for attackers to guess or brute-force.

3. Consider adding rate limiting to your application, to make it harder for an attacker to brute-force the credentials.

## Resources
1. [Basic authentication in GO](https://www.alexedwards.net/blog/basic-authentication-in-go)
