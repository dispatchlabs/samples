- Run
	- `Go-NaiveChain ":9999"`

- Get blockchain
	- `curl http://localhost:9999/blocks`

- Create block
	- `curl -H "Content-type:application/json" --data '{"PreviousHash": "", "Timestamp": "", "Data" : "Some Data", "Hash": ""}' http://localhost:9999/mineBlock`
