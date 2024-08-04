# identity-service

User management and authentication service.

## Quickstart

```bash
# Generate a keypair for JWT
openssl genpkey -algorithm ed25519 -outform PEM -out private_key.pem
openssl pkey -in private_key.pem -pubout -outform PEM -out public_key.pem

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
