# Govie
Welcome to Govie :)


## How to run the project?
```shell
> git clone https://github.com/vit0rr/Govie
> cd Govie
> go get .
> go run .
```

## How to add a movie?
```shell
curl http://localhost:8080/movies \

    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "10","title": "Spider-Man","price": 49.99}'
```
## How to get every movie?
```shell
curl http://localhost:8080/movies
```

## How to delete a movie?
```shell
curl http://localhost:8080/movies/2 \ 
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE"
```