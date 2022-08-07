set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -gcflags "-N -l" -o build/rn-agent ./