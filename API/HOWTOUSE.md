# API Driver
## Language 
Golang

## How to run
In the API directory run 
```
go run main.go
```

## File Definitions
### Directories
- 📁 conn/: database connection package
- 📁 json_templates/: templates for how data should be structured in request body
- 📁 models/: structs used throughout all packages
- 📁 rest/: API routes and handlers
- 📁 SQL/: SQL queries for creation of tables
### Files
- 📄 main.go: main driver for API package
- 📄 .env: database configuration
- 📄 go.mod: go packages
- 📄 go.sum: go packages

## How to structure requests
- GET Routes
    - Account:
        - Usage: grabbing all information of an account generic route grabs all accounts use query string id=NUM to grab specific account
        - URI: localhost:8080/account?[optional query strings]
        - Return Sample:
    ```[
        {
            "id": 1,
            "fname": "John",
            "lname": "Doe",
            "fullname": "John Doe",
            "email": "john.doe@example.com",
            "pwd": "password123",
            "pnum": 1234567890,
            "username": "johndoe",
            "acesslevel": "user"
        },
        {
            "id": 7,
            "fname": "guest",
            "lname": "guest",
            "fullname": "guest guest",
            "email": "guest@guest.com",
            "pwd": "guest",
            "username": "guest",
            "acesslevel": "guest"
        }
    ]
    ```

    - Posts
        - Usage: grabbing all the posts in database
        - URI: localhost:8080/posts?[optional query strings]
        - Return Sample:
    ```
    [
        {
            "id": 2,
            "title": "Post 2",
            "descr": "A post about science",
            "genre": "Science",
            "authorId": 2,
            "numUp": 5,
            "numDown": 0,
            "picId": 2,
            "postedDate": "2023-11-09 00:00:00"
        },
        {
            "id": 3,
            "title": "Post 3",
            "descr": "A fun post about nature.",
            "genre": "Nature",
            "authorId": 3,
            "numUp": 12,
            "numDown": 3,
            "picId": 3,
            "postedDate": "2023-11-09 00:00:00"
        }
    ]
    ```
    
    - Posts/Fullcontext
        - Usage: grabbing all the posts in database plus picture and author information
        - URI: localhost:8080/posts/fullcontext?[optional query strings]
        - Return Sample:
    ```
    [
        {
            "id": 2,
            "title": "Post 2",
            "descr": "A post about science",
            "genre": "Science",
            "authorInfo": {
                "id": 2,
                "fullname": "Alice Smith",
                "username": "alicesmith",
                "acesslevel": ""
            },
            "numUp": 5,
            "numDown": 0,
            "picInfo": {
                "id": 2,
                "imgname": "image2.png"
            },
            "postedDate": "2023-11-09 00:00:00"
        },
        {
            "id": 3,
            "title": "Post 3",
            "descr": "A fun post about nature.",
            "genre": "Nature",
            "authorInfo": {
                "id": 3,
                "fullname": "Bob Johnson",
                "username": "bobjohnson",
                "acesslevel": ""
            },
            "numUp": 12,
            "numDown": 3,
            "picInfo": {
                "id": 3,
                "imgname": "image3.jpeg"
            },
            "postedDate": "2023-11-09 00:00:00"
        },
    ]
    ```

- POST Routes
    - Account
        - Usage: Create a new account has checks for if account w same username exists
        - URI: localhost:8080/account/new
        - Example Body:
        ```
        {	
        "Fname":    "Gio",
        "Lname":    "Joseph",
        "Fullname": "Gio Joseph",
        "Email":    "gio@gmail.com",
        "pwd":      "password",
        "Pnum":     17187155801,
        "Username": "Shnybones"
        }
        ```

        - Return Sample:
       ```
       (if account is new)
       [
        {
            "id": 8,
            "fname": "james",
            "lname": "jamerson",
            "fullname": "james jamerson",
            "email": "guest@guest.com",
            "pwd": "guest",
            "pnum": 17187155801,
            "username": "jimm",
            "acesslevel": "user"
        }
        ]
       ```
        ```
        (if account with username already exits)
        "Account Already Exits"
        ```
    - Post
        - Usage: create a new post
        - URI: localhost:8080/posts/create
        - Example Body:
        ```
        {
            "id": 1,
            "title": "TITLE",
            "descr": "CONTENTS OF THE POST",
            "genre": "GENRE OF THE POST",
            "authorId": 1,
            "picId": 1
        }
        ```
        - Return Sample:
        ```
        "post created"
        ```
    - Login
        - Usage: using either username or email combined with password returns information about user if account exits and error if password is incorrect NOTE: if username and password are both guest -> login as guest
        - URI: localhost:8080/login
        - Example Body:
        ```
        (NOTE: request must contain all 3 properties, handler only checks for one so both are not needed but it must still be sent as an empty string)
        {
            "username": "johndoe",
            "password": "password123",
            "email" : ""
        }
        ```
        - Return Sample:
        ```
        (if user/email and password are correct)
        [
            {
            "id": 1,
            "fname": "John",
            "lname": "Doe",
            "fullname": "John Doe",
            "email": "john.doe@example.com",
            "pwd": "password123",
            "pnum": 1234567890,
            "username": "johndoe",
            "acesslevel": "user"
            }
        ]
        ```
        ```
        (if user/email and password are invalid)
       "invalid password"
        ```
- PATCH Routes
    - Account
        - Usage: update an account
        - URI: localhost:8080/account/update
        - Example Body:
        ```
        {
        "old":
            {	
            "id":		6,
            "fname":    "f",
            "lname":    "Jablonka",
            "fullname": "Adam Jablonka",
            "email":    "ajjablonka@gmail.com",
            "pwd":      "password",
            "pnum":     17187155801,
            "age":	    22,
            "username": "WebsiteLover"
            },
        "new":
            {
            "id":		6,
            "fname":    "Adam",
            "lname":    "Jablonka",
            "fullname": "Adam Jablonka",
            "email":    "ajjablonka@gmail.com",
            "pwd":      "password",
            "pnum":     17187155801,
            "age":	    22,
            "username": "WebsiteLover"
            }
        }
        ```
        - Return Sample:
        ```{
            "id": 6,
            "fname": "Adam",
            "lname": "Jablonka",
            "fullname": "Adam Jablonka",
            "email": "ajjablonka@gmail.com",
            "pwd": "password",
            "pnum": 17187155801,
            "age": 22,
            "username": "WebsiteLover"
        }
        ```
    - Post
        - Usage: update a post
        - URI: localhost:8080/post/update
        - Example Body:
        ```
        {
            "old":
                {
                    "id": 1,
                    "title": "How to get better at golf",
                    "descr": "I suck at golf how do i do better please help",
                    "genre": "Technology",
                    "authorId": 1,
                    "picId": 1
                },
            "new":
                {
                    "id": 1,
                    "title": "How to get better at golf",
                    "descr": "I suck at golf how do i do better please help",
                    "genre": "Technology",
                    "authorId": 1,
                    "picId": 1
                }
        }
        ```
        - Return Sample:
        ```
        {
            "id": 1,
            "title": "How to get better at golf",
            "descr": "I suck at golf how do i do better please help",
            "genre": "Technology",
            "authorId": 1,
            "picId": 1
        }
        ```
- DELETE Routes
    - Account
        - Usage: delete an account
        - URI: localhost:8080/account/delete/:ID_OF_USER
        - Return Sample:
        ```
        "account ACCOUNTID deleted"
        ```
    - Post
        - Usage: delete a post
        - URI: localhost:8080/account/delete/:ID_OF_POST
        - Return Sample:
        ```
        "post POSTID deleted"
        ```
