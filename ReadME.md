To run the tests there are multiple options:

To run all tests, from your root directory in the terminal, type: `go test ./..`

In your terminal, cd into the directory you wish to test: ie: `cd Repository`
Next, in your terminal, type: `go test`

OR you can run these commands: 

`go get -u github.com/onsi/ginkgo/ginkgo`

`go get github.com/onsi/gomega/...`

In the root directory run the command in terminal: `ginkgo -r`
	

To quick start the application, in the terminal:

`go run main.go`

Navigate to http://localhost:8080 and refer to the routes table or curl commands for manually testing the application.

TO RUN THE REQUESTS ON POSTMAN:

Postman example to test:

    * In terminal enter command `go run main.go` or `go build main.go`
    * Go to browser and enter domain `http://localhost:8080/user`
    * Open postman
        * To create a new user, change request to POST and provide url (as above)
        * Click to Body; click raw; drop down JSON
        * Enter in details: 
            {"employeeID":3,"name":"OJ Old","email":"oj.old@gmail.com","phone":"0777000111","pin":1234}
        * Click Send
        * Change Response Body to JSON
        * Check status as expected (200)


TO CURL THE REQUESTS:

Data must be:
EmployeeID: 	unique int
CardID: 	    unique string, alphanumeric, 16 characters in length
Name: 		    string, must not be empty
Email: 		    string, must be valid email syntax
Phone:		    string
Pin:		    string containing 4 digits
Balance 	    int, set to zero when created, must top up after creating user

Start the application by running `go run main.go` or `go build main.go`

In your terminal, paste these commands or change the user details to whatever you would like to test.


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
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234","amount":50}
 
View Balance
curl --location --request GET 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234"}'
 
Make Purchase
curl --location --request PUT 'http://localhost:8080/purchase' \
--header 'Content-Type: application/json' \
--data-raw '{"cardID":"123b567c91B234b6","pin":"1234","amount":45}
 
Logout
curl --location --request GET 'http://localhost:8080/logout/123b567c91B234b6'
 
To confirm the unhappy paths type in invalid information 
ie: the wrong pin, cardID, amounts that reduce the balance below zero, etc. 
Refer to the routes table for error messages.


