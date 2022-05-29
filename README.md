# test
This is a go api project.
CRUD commends example:
  1. GET: curl --location --request GET 'http://localhost:8080/api/v1/comment/uuid'
  You can use online uuid generator to test if a certain id comment exists.
  
  2. POST: curl --location --request POST 'http://localhost:8080/api/v1/comment' --header 'Content-Type: application/json' --data-raw '{"slug": "hello", "body": "lala", "author": "qqllo"}'

  3. UPDATE: curl --location --request PUT 'http://localhost:8080/api/v1/comment/uuid' --header 'Content-Type: application/json' --data-raw '{"slug": "hello", "body": "I updated this comment", "author": "qqllo"}'

  4.
