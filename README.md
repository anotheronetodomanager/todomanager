# TodoManager

Currently just a simple echo server

## Usage

You need only Go downloaded.

### Download

1. Clone it `git clone https://github.com/anotheronetodomanager/todomanager.git`
2. Install dependencies: `go mod download`

### Run

Run it with command `go run ./cmd/todomanager/`. Use `-port` flag to specify port for server. Flag `-help` shows output:
```
Usage of ./todomanager:
  -port string
        server port (default "8080")
```

### Build

Build executable with `go build ./cmd/todomanager/`

### Example

1. Run server. You will see log message `Starting echo server on port 8080`
2. Open new console and make a request to server: `curl localhost:8080 -d "hello"`
3. Make sure you received "Echo: hello"