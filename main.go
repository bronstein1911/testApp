package main

func main() {
	if err := run(); err != nil {
		debugLogger.Fatal(err)
	}
}
