# tiny-url

A basic tiny URL in memory generator.

### Usage
Create new tiny URL via CURL command
```shell
curl http://localhost:8080/ \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"url": "https://go.dev/doc/tutorial"}' 
```

Get all in memory tiny URLs
```shell
curl http://localhost:8080
``` 

### Docker
In order to run the code inside docker container please run:
```shell
docker build --tag tiny-url .
docker run --rm -p 8080:8080 tiny-url
```