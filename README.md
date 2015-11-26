# cmpe273-lab3

Setup

##Server
go get github.com/aggarwalsomya/cmpe273-lab3/server

cd src/github.com/aggarwalsomya/cmpe273-lab3/server

go run *


##Client

go get github.com/aggarwalsomya/cmpe273-lab3/Client

cd src/github.com/aggarwalsomya/cmpe273-lab3/Client

go run *




##How to use:

Server here is a multi threaded process which listens at localhost on three ports: 3000, 3001, 3002

Client is a light weight shard manager service which routes the request to the appropriate server using consistent hashing. I used fnv package to generate a 32 bit numeric hashing of any string. The consistent hashing algorithm is very simple and doesn't consider any replication factor. 


###To put a key to server through client (client takes care of sharding)

http://localhost:8080/keys/1/a

Response: 200 OK

###To put a key directly to a server:

http://localhost:300x/keys/1/a

Response: 200 OK

### To get a key through client (client takes care of sharding)

http://localhost:8080/keys/1

Response: 
Status 200 OK
{"key":1,"value":"a"}


### To get a key directly from server

http://localhost:300x/keys/1			(call appropriate server for key 1)
Response: 200 OK
{"key":1,"value":"a"}



###To get all keys from a server

http://localhost:300x/keys

Response: 200 OK
[{"key":2,"value":"b"},{"key":4,"value":"d"},{"key":7,"value":"g"},{"key":9,"value":"i"},{"key":10,"value":"j"},{"key":1,"value":"a"}]









