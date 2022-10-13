# How to build/start server and client
Ensure you have GNU Make, Go (>= 1.19), NodeJs and npm is installed.

Before building node modules must be installed with ```npm i```.

Client: run 'make' in the server directory to both build and run the client.
Server: run 'make' in the client directory to both build and run the server.

Server will run on port 3000 and web client will run on port 8080;

# How to use
Open localhost:8080 in a browser and send a password with the form input field.

Ive hardcoded a server hash of the client hashed password 'pass'. This means only the password 'pass' will be accepted.

I have not done the voluntary tasks:)
