while true; do
	go build -o main main.go
	dlv debug . -l 0.0.0.0:2345 --headless=true --log=true -- server
done