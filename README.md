# zenTotem
test for iConText group
task: https://observant-hero-c7c.notion.site/Go-3cc65a7d7c3e44c19b2e0543a98be2d2

# flags parameters
the following flags can be used to run the database and redis application with certain parameters
	-user default value="dobb2", description: "a user postgress"
	-hostDB" default value="localhost", description:  "a host of postgress"
	-password default value="root", description: "a password user postgress"
	-portDB default value="54320", description: "a port of postgress"
	-address default value="127.0.0.1:8080", description: "address of postgress"
	-db default value="testWB", description: "name db"
	-host default value="127.0.0.1", description: "a host redis"
	-port default value="6380", description: "a port redis"
# run aplication 
```
cd cmd
go run main.go -port 8080
```