# - SOBOT - Automated Social Media Sharing Tool

Social media post sharing tool

## Features

- The application has a stable beta version. Errors that will occur will be on the selector codes.
- For now it only has instagram and twitter support.
- Saves screenshots for bug tracking

## Usage/Examples

```go
package main

import (
    "github.com/canack/sobot/src"
)

func main(){
    // User creds
    username := "nickname"
    password := "secretpass123"

    // Post parameters
    file := "/home/user/Desktop/photo.jpg"
    comment := "Have a nice day!"

    // Sharing on Instagram
    sobot.Instagram(username, password).SetFile(file).SetCaption(comment).Share().Debug()
	
    // Sharing on Twitter
    sobot.Twitter(username, password).SetFile(file).SetCaption(comment).Share().Debug()
}
```


## Routemap

- [+] Instagram support (currently stable-beta)

- [+] Twitter support (currently stable-beta)


## Contributors

- [@canack](https://www.github.com/canack)