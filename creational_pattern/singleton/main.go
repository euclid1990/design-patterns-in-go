package main

func main() {
	connectionString := "myuser@example.com:3306/main-schema"
	GetInstanceByDoOnce(connectionString)
	GetInstanceByDoLock(connectionString)
}
