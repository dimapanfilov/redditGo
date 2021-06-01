# redditGo

Learning Golang by making a small reddit clone based off it.

To run this currently, run 
```
1. make postgres
2. make adminer
```
Log into server. If localhost:8080 does not work with original credentials, use the address seen when running 
```
docker ps
docker inspect /* address of postgres container*/
```
Using the obtained IP, log in.
Now, run ```make migrate``` to create all the tables.

Finally, running ```go run cmd/redditgo/main.go``` will permit you to access localhost:3000/threads 

To create threads go to where the tables are created, go to threads, and add a new one using a random uuid generator online. 
