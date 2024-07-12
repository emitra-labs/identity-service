run:
	godotenv go run .

test:
	godotenv go test -v -count=1 ./rest
