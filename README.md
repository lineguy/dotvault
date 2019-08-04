# dotVault

dotVault is an a wrapper script that I wrote in go (1st Go Project). Its purpose is to enable anyone to easily store & retrieve sensitive configuration files from their LastPass vault (https://www.lastpass.com/)

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

### Configuration

dotVaults main configuration is setup in the `files.json` file. The file default file includes some examples to show how to configure the files you would like to manage with dotVault. If you need to find the `lpass_id` you can see each item in your LastPass vault by using the `lpass ls` command, these items will need to exist in your LastPass before configuring `files.json`.

#### LastPass Item Example



#### files.json Example
```
{
  "files": [
    {
      "name": "Give this part of the config a name",
      "lpass_id": "LastPass id of the stored file",
      "path": "Local path to store the file",
      "owner": "Username of the user to own file",
      "group": "Group name to associate with file",
      "chmod": "File mode bits for securing file"
    },
    {
      "name": "File Name 2",
      "lpass_id": "3495167510810510691",
      "path": "/home/username/.ssh/id_rsa",
      "owner": "username",
      "group": "username",
      "chmod": "0700"
    }
  ]
}
```

### Why would I want my sensitive configuration files in LastPass?

Your sensitiver configuration files might be the most important files on your machine. dotVault enables you to back them up and store them in a secure manner, its also a very convienent way of being able to download them onto new devices as and when you need to.
