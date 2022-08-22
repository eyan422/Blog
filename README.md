mysql -uroot -p12345678 < creation.sql

go mod tidy

go test ./...

go run .

```curl
curl -X GET localhost:8080/articles/1

curl -X GET localhost:8080/articles/2

curl -X POST --location --request POST 'http://localhost:8080/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "test_title",
    "content": "test_content",
    "author": "test_author"
}'
```
