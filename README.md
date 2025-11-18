# Store a key-value pair
```bash
# Create a value
go run main.go -action=put -key=name -value=John -filename=data.txt -server=server1
go run main.go -action=put -key=age -value=23 -filename=data.txt -server=server1

# Retrieve a value
go run main.go -action=get -key=name -filename=data.txt -server=server1
go run main.go -action=get -key=age -filename=data.txt -server=server1

# Update a value
go run main.go -action=put -key=name -value=Tom -filename=data.txt -server=server1
```