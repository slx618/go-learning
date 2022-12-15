package main


 func main() {
	server := NewServer("0.0.0.0", 9970)
	server.Start()
 }