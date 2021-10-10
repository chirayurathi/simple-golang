# simple-golang

### A simple Instagram Backend API made on Golang and MongoDB for storage.

## Documentation :open_book:

  Checkout the documentation and examples [Here](https://documenter.getpostman.com/view/12247743/UV5Rketx) .

## Features :sparkles:	
- ## Users
  - Add User
  - Get User By User ID
- ## Posts
  - Add Post
  - Get Post By Post ID
  - Get List of Posts By User ID
- ## Security
  Passwords are encrypted before saving to Database and Decripted before serving back to client
- ## Pagination
  optional pagination on get posts by user ID API , pass the following params for pagination:
    - page : page number .
    - count: No. of documents each page.

## Sample :pencil2:	
- ### Add User
  ![Add User Sample](/sample/adduser.JPG)
- ### Get User By Id
  ![Get User Sample](/sample/getuser.JPG)
- ### Add Post
  ![Add Post Sample](/sample/addpost.JPG)
- ### Get Post By ID
  ![Get Post Sample](/sample/getpost.JPG)
- ### Get Posts By User ID
  ![Get Post List Sample](/sample/userpost.JPG)



    
