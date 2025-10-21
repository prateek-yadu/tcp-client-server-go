## Project Structure

```
tcp-client-server-go/
├── client/
│ └── client.go # TCP client code
├── server/
│ └── server.go # TCP server code
├── go.mod # Go module file
└── README.md
```

### Run Server
```bash
cd server
go run server.go
```

Server will start listening on `127.0.0.1:8000`.

### Run Client
```bash
cd client
go run client.go
```

Client connects to server on `127.0.0.1:8000`.
