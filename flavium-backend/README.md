## Build
`cd flavium-backend`  

`make`  

`cd ..`  

`docker-compose up --build backend`  

## Running
`docker-compose up backend`

## Requests
`curl -X GET 'http://localhost:8080/v1/torrent' -b oauthstate=pseudo-random`  

`curl -X POST -d '{"magnetLink":"link"}' 'http://localhost:8080/v1/torrent' -b oauthstate=pseudo-random`  

## Transmission cli
You need to edit the transmission/config/settings.json that should be generated when running `docker-compose up transmission`   
* rpc-whitelist-endabled should be false   
* rpc-host-whitelist should be "transmission"
