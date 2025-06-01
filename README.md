# Encryption Playground

A Go-based playground for experimenting with various encryption and hashing algorithms. This project provides implementations of common cryptographic functions for educational and testing purposes.

## Features

The project includes implementations of the following cryptographic algorithms:

- **RSA**: Asymmetric encryption with 2048-bit key pairs
- **AES**: Symmetric encryption
- **Argon2**: Password hashing
- **SHA256**: Cryptographic hashing
- **BCrypt**: Password hashing

## Prerequisites

- Go 1.24.1 or later
- `golang.org/x/crypto` package

## Installation

1. Clone the repository:

```bash
git clone https://github.com/danielmesquitta/encryption-playground.git
cd encryption-playground
```

2. Install dependencies:

```bash
go mod download
```

## Project Structure

```bash
├── cmd/      # Application entrypoint
├── pkg/      # Core packages
│ ├── aes/    # AES encryption implementation
│ ├── argon2/ # Argon2 password hashing
│ ├── bcrypt/ # BCrypt password hashing
│ ├── rsa/    # RSA encryption implementation
│ └── sha256/ # SHA256 hashing implementation
├── go.mod    # Go module file
└── go.sum    # Go module checksum file

```
