# OpenPass 🔓

OpenPass is a simple password manager written in Go. Uses the PBKDF2 cryptographic algorithm to encrypt passwords. As of the latest release, it is a CLI tool only.

This is a trustless password manager. Every password is saved in an encrypted form, and can only be properly decrypted with the master password. The master password is never saved, and is only used to decrypt the passwords when needed.

Note: Any password you input will receive a "valid" output. That is, the program will not tell you if you have entered the wrong password. This is to prevent brute force attacks.

## Usage ⚙️

To see a full list of commands, run `keys help` and type in any arbitrary password. This will show you the help menu.

## Installation 🛠️

### Downloading the binary

Head over to the 'releases' page and download the latest binary. You can then run the binary with `./keys` or `keys.exe` on Windows.

If you prefer the convenience, you can add the binary to a `usr/bin` directory, or add it to your PATH, and run without the `./` prefix.

### Building from source

Building from source requires that you have the following tools installed:
- git
- make
- go

1. Start by cloning the repository:
```
git clone github.com/skjdfhkskjds/openpass
```

2. Navigate to the cloned repository and build the binary:
``` 
cd openpass
make build
```

That's it! The binary will be in the `bin` directory, and you can run it with `keys` or `keys.exe` on Windows.

## ⚠️ Warning ⚠️

This is a work in progress. It is not recommended to use this program for storing sensitive passwords. I am not responsible for any lost passwords or data.
