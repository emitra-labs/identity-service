run:
	godotenv go run .

test:
	SKIP_AMQP_PUBLISHING=true godotenv go test -v -count=1 ./tests
