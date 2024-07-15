# identity-service

User management and authentication service.

## Quickstart

```bash
# Generate a keypair for JWT
openssl genpkey -algorithm ed25519 -outform PEM -out private_key.pem
openssl pkey -in private_key.pem -pubout -outform PEM -out public_key.pem

# Encode jwt keypair into base64
cat private_key.pem | base64
cat public_key.pem | base64

# Create .env file, update its content
cp .env.example .env

# Run the program
make run
```

## Features

- Authentication.
- Email verification.
- Session management.
- User management.
