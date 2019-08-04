# dotVault

dotVault is an a wrapper script that I wrote in go. Its purpose is to enable anyone to easily store & retrieve sensitive configuration files from their LastPass vault (https://www.lastpass.com/)

## Getting Started

### Prerequisites

__Disclaimer:__ _this has only been tested on Ubuntu Linux._

* Unix-like operating system that supports LastPass-CLi i.e. Linux

### Dependencies

dotVault is just a go wrapper script so it depends on lastpass-cli and go so install it using the below command:
```
sudo apt install lastpass-cli go
```

### Installation

```
go run main.go
```
```
go build main.go
```

### Ussage

#### files.json

### Why would I want my sensitive configuration files in LastPass?

Your sensitiver configuration files might be the most important files on your machine. dotVault enables you to back them up and store them in a secure manner, its also a very convienent way of being able to download them onto new devices as and when you need to.
