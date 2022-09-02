export GOARCH=amd64

export GOOS=darwin
go build -o ./bin/mac/bench .

export GOOS=linux
go build -o ./bin/linux/bench .

export GOOS=windows
go build -o ./bin/windows/bench.exe .