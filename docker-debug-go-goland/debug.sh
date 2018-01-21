while true; do
	go build -o main main.go
	dlv debug --headless --listen=:2345 --api-version=2
done