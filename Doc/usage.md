# SE-project2

Software Engineering Methodology Project 2: Debug

Team Members: Angel Bajracharya, Derek Feliu, Giovanny Joseph, Nicole Magpantay, Christopher Miles

# How to use

**Requirements:**
- GoLang
- JavaScript
- VScode
    - extension [live server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer)

Clone the repository
```
git clone https://github.com/GiovannyJ/SE-project2
```

Open the project in the terminal wherever it is downloaded

Once in the project start the server by going to the API directory and running main.go
```
cd API
go run main.go
```

Now start the web server at [SE-Project2/Frontend/src/login.html](../Frontend/src/login.html) using the [live server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) extension
 and you can begin using the application


# Access Levels

- **Guest:**
  - Can view posts.

- **User:**
  - Can view posts.
  - Can create posts.
  - Can edit own account.

- **Verified User:**
  - Can view posts.
  - Can create posts.
  - Can edit own account.
  - Can delete posts.
  - Can delete comments.

- **Admin:**
  - Can view posts.
  - Can create posts.
  - Can edit own account.
  - Can delete posts.
  - Can delete comments.
  - Can delete other accounts.
  - Can edit access level of other accounts.

# Pages

## login.html

- Login as guest: Binds to guest account.
- Login as user:
  - Checks if user/pass exists in the database and are good.
- Button: [resetpassword.html](#resetpasswordhtml)
- Button: [createacct.html](#createaccthtml)

## resetpassword.html

- Using username, create a new password.

## createacct.html

- Fill form to make an account -> Goes to [menu.html](#menuhtml) bound with the account.
  - Checks if already logged in.

## menu.html

- Buttons to:
  - [searchresults.html](#searchresultshtml) (Guest, User, Verified User, Admin)
  - [ask.html](#askhtml) (User, Verified User, Admin)
  - [updateAccount.html](#updateaccounthtml) (User, Verified User, Admin)
  - [admin_panel.html](#adminpanelhtml) (Verified User, Admin)

## searchresults.html

- Loads all posts.
- If category/text box filled and search -> Loads filtered values.
- Buttons to post in context ([question.html](#questionhtml)).

## question.html

- Loads post in context.
- Comments allowed for everyone except guest.

## ask.html

- Allows users with proper access levels to submit questions.
- Fill form to ask a question.
- When a post is made, goes back to [searchresults.html](#searchresultshtml).

## updateAccount.html

- Fill form to modify account details (first, last, email, pnum, username).

## admin_panel.html

- Delete account (Admin).
- Delete post (Admin, Verified User).
- Delete comment (Admin, Verified User).
- Change access level (Admin).



# Extras

For more information on how to use the API [click here](https://github.com/GiovannyJ/SE-project2/blob/main/API/HOWTOUSE.md)