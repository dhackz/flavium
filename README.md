# Flavium
Flavium is a project to automate and simplify downloads, formatting and provisioning of media files. \
The project should have few requirements to set up to make it easy for the user to install. This is accomplished by using infrastructure as code in the form of dockerfiles. Currently the only requirements for the user are docker, docker-compose and make. Zuul-CI is used to automate tests and make sure that no changes that breaks the system is merged to the master branch.

## Setup
### Env
This needs to be in your .env:  
```
GOOGLE_CLIENT_ID=googleid
GOOGLE_CLIENT_SECRET=googlesecret
TRANSMISSION_HOST=transmission
REACT_APP_MOVIE_KEY=secret-key
APPROVED_EMAILS=example@gmail.com,example2@gmail.com
```

### Transmission cli
You need to edit the transmission/config/settings.json that should be generated when running `docker-compose up transmission`   
* rpc-whitelist-endabled should be false   
* rpc-host-whitelist should be "transmission"

## Build
`make`  

`docker-compose up --build backend`  
`docker-compose up --build dashboard`  

## Running

### Dryrun
`docker-compose up backend`
`docker-compose up dashboard`


### Full system
Edit the docker-compose.yaml to set `--dry_run=false`

`docker-compose up backend`
`docker-compose up dashboard`
`docker-compose up transmission`

### Requests
`curl -X GET 'http://localhost:8080/v1/torrent' -b oauthstate=pseudo-random`  

`curl -X POST -d '{"magnetLink":"link"}' 'http://localhost:8080/v1/torrent' -b oauthstate=pseudo-random`  

