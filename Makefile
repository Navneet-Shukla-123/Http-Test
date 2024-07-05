test:
	# Run tests and generate coverage profile
	go test -coverprofile=coverage.out

	# View coverage details in the terminal
	go tool cover -func=coverage.out

	# Generate and open an HTML coverage report
	go tool cover -html=coverage.out -o coverage.html

run:
	go run main.go