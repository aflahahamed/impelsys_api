# impelsys_api

To run the API type the following command:
`make run`
This will create a binary and run the API on `localhost:4000` port.

## Create course
To create a course give the following payload with `POST` request to the `localhost:4000/course` path:
```
{
    "couresename":"Golang",
    "courseprice": 2000,
    "authorname":"aflah",
    "authorwebsite": "www.youtube.com"
}
```
![image](https://github.com/aflahahamed/impelsys_api/assets/61117867/6007208b-3c30-4075-a8a5-8fd2eb6011a9)


## Update course
To update a course give the following payload with `PATCH` request along with the course ID to the `localhost:4000/course` path:
```
{
    "id":2,
    "couresename":"Golang",
    "courseprice": 2000,
    "authorname":"Mike",
    "authorwebsite": "www.youtube.com/Mike"
}
```
![image](https://github.com/aflahahamed/impelsys_api/assets/61117867/8a8a0058-a544-4889-a410-b03ef4fe507c)

