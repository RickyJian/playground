# Init

this project is my first golang with docker practice

## How to use

```
## build golang container
$ docker build -t ricky/init .

## run web service, type `localhost:8080` will link to service
$ docker run --rm -p 8080:8080 ricky/init 
```
