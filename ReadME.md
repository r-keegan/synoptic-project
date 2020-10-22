First you will need to install go mod, in your terminal type:

`go mod init`

To begin using the app, go to the project root directory and simply type in your terminal:

`go run main.go`

To run the tests there are multiple options:

In your terminal, cd into the directory you wish to test: ie: `cd Repository`
Next, in your terminal, type: `ginkgo`

To run all tests, from your root directory in terminal, type: `ginkgo -r`

Depending on your IDE: press the `play` button on each test file

Navigate to http://localhost:8080 and refer to the routes table or curl commands for manually testing the application.

Postman example to test:

    * In terminal enter command `go run main.go`
    * Go to browser and enter domain `http://localhost:8080/user`
    * Open postman
        * To create a new user, change request to POST and provide url (as above)
        * Click to Body; click raw; drop down JSON
        * Enter in details: 
            {"employeeID":3,"name":"OJ Old","email":"oj.old@gmail.com","phone":"0777000111","pin":1234}
        * Click Send
        * Change Response Body to JSON
        * Check status as expected (200)

Must run login GET request first to purchase/topup/getBalance with the same cardID and pin
