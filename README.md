Install go (project used: v1.14) see website for details: https://golang.org/doc/install	

To quick start the application, enter  in your terminal:

`go run main.go`

Navigate to http://localhost:8080 and refer to the routes table or curl commands for manually testing the application.

TO RUN THE REQUESTS ON POSTMAN:

Refer to this guide for importing the postman collection:
https://learning.postman.com/docs/getting-started/importing-and-exporting-data/

Postman example to test:

    * In terminal enter command `go run main.go`
    * Open postman:
        * Enter URL: http://localhost:8080/user
        * To create a new user, change request to POST
        * Click to Body; click raw; drop down JSON
        * Enter in details: 
            {"employeeID":3,”cardID:”9876543210abcedp”"name":"OJ Old","email":"oj.old@gmail.com","phone":"0777000111","pin":1234}
        * Click Send
        * Change Response Body to JSON
        * Check status as expected (200)


TO CURL THE REQUESTS:

Data must be:

     EmployeeID: 	    unique int
     CardID: 	    unique string, alphanumeric, 16 characters in length
     Name: 	        string, must not be empty
     Email: 	    string, must be valid email syntax
     Phone:	        string, must not be empty
     Pin:		    string containing 4 digits
     Balance 	    int, set to zero when created, must top up after creating user

Start the application by running go run main.go

In your terminal, paste these commands or change the user details to whatever you would like to test.

HAPPY PATH CURL

Create User
curl --location --request POST 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{"employeeID":8,"cardID":"123b567c91B234b6","name":"Greg Harris","email":"greg.harris1@gmail.com","phone":"0799007007","pin":"1234","balance":0}'

Present Card
curl --location --request GET 'http://localhost:8080/cardPresented/123b567c91B234b6'

Login
curl --location --request GET 'http://localhost:8080/user/auth' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234"}'
 
TopUp
curl --location --request PUT 'http://localhost:8080/topup' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234","amount":50}'
 
View Balance
curl --location --request GET 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234"}'
 
Make Purchase
curl --location --request PUT 'http://localhost:8080/purchase' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234","amount":45}'
 
Logout
curl --location --request GET 'http://localhost:8080/logout/123b567c91B234b6'
 
To confirm the unhappy paths type in invalid information 
ie: the wrong pin, cardID, amounts that reduce the balance below zero, etc. 

Unhappy path tests are also covered in unit and integration tests.

Refer to the Route Path document or Test Plan for error messages.

If you want to go back to an initial state - delete the main.db inside the synoptic-project source code. It can be safely deleted.

To run the tests there are multiple options:

To run all tests, from your root directory in the terminal, type: `go test ./..`

In your terminal, cd into the directory you wish to test: ie: `cd Repository`
Next, in your terminal, type: `go test`
__
OR  you can run these commands: 

    `go get -u github.com/onsi/ginkgo/ginkgo`
    `go get github.com/onsi/gomega/...`

In the root directory run the command in terminal: `ginkgo -r`
