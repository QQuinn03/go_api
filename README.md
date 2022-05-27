# test
This is a go api project.
CRUD commends:
  GET: curl --location --request GET 'http://localhost:8080/api/v1/comment/uuid'.
  
  You can use online uuid generator to test if a certain comment exists.
  
  POST: curl --location --request POST 'http://localhost:8080/api/v1/comment' --header 'Content-Type: application/json' --data-raw '{"slug": "hello", "body": "lala", "author": "qqllo"}'
