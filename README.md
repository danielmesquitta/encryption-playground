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

Write a readme.md file for this project

I'll help you create a README.md file for this project. First, let me explore the project structure to understand what it's about.
Let me check the contents of the pkg directory to understand the main functionality.
Let me look at the go.mod file to understand the project dependencies and module name.
Let me check one of the encryption implementations to better understand the project's purpose.
Based on my exploration of the project, I'll create a comprehensive README.md file. Here's what I suggest:

```bash
├── cmd/      # Command-line applications
├── pkg/      # Core packages
│ ├── aes/    # AES encryption implementation
│ ├── argon2/ # Argon2 password hashing
│ ├── bcrypt/ # BCrypt password hashing
│ ├── rsa/    # RSA encryption implementation
│ └── sha256/ # SHA256 hashing implementation
├── go.mod    # Go module file
└── go.sum    # Go module checksum file

```
