# password-manager

A simple password manager written in Go. Uses the PBKDF2 cryptographic algorithm to encrypt passwords.

This is a trustless password manager. The program does not store any passwords, only the encrypted version of them. The user is responsible for remembering the master password used to create the encrypted passwords.

Temporarily all passwords are stored to a JSON file, but all sensitive pieces of data are encrypted.

## Usage

TODO

## Installation

To simply run the program, download the keys executable binary from the releases page. To build from source, run `make build` in the root directory.