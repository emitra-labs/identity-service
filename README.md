# identity-service

User management and authentication service.

## Quickstart

```bash
# Create .env file, update its content
cp .env.example .env

# Load .env file and run the program
godotenv go run .

# Re-init OpenAPI docs
swag init -g rest/rest.go
```
