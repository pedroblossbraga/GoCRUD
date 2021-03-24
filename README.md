# Go API 

![](images/Rest-API.png)

## CRUD operations with GORM

- Create (POST)
- Read (GET)
- Update (PUT)
- Delete (DELETE)

## API calls:
-   / 
![](images/homepic.jpeg)
- 	/users -> GET (All Users)
-	/user/{name}" -> "DELETE
-	/user/{name}/{lastname}/{city}/{email} -> PUT
-	/user/{name}/{lastname}/{city}/{email} -> POST

## Run the app

        go run *.go

## Examples

- Create
    - starting from scratch:

![](images/1.jpeg)
    
    - creating john johnson, from newyork, john@email.com:

![](images/2.jpeg)

    - creating annya solov, from moscow, annya@mail.com:

![](images/3.jpeg)

- Read

    - current data:

![](images/4.jpeg)

- Update

    - annya moved to barcelona, lets update that info:

![](images/5.jpeg)

    - current data:

![](images/6.jpeg)

- Delete

    - john moved to mars, lets delete him:

![](images/7.jpeg)

- result

![](images/8.jpeg)

## References:
- A great tutorial, and main source of study: https://tutorialedge.net/golang/golang-orm-tutorial/
- Go ORM querie tutorial : https://gorm.io/docs/query.html

