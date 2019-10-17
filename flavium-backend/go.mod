module github.com/dhackz/flavium/flavium-backend

go 1.13

require (
	github.com/dhackz/flavium/flavium-backend/server v0.0.0
	github.com/dhackz/flavium/flavium-backend/torrents v0.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	google.golang.org/grpc v1.24.0 // indirect
)

replace github.com/dhackz/flavium/flavium-backend/server v0.0.0 => ./pkg/server/

replace github.com/dhackz/flavium/flavium-backend/torrents v0.0.0 => ./pkg/torrents/
