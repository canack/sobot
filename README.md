# - SOBOT - Automated Social Media Sharing Tool

Social media post sharing tool

## Features

- The application has a stable beta version. Errors that will occur will be on the selector codes.
- For now it only has instagram and twitter support.
- Saves screenshots for bug tracking
- Username and password are used for first login only. Then it continues over the cookie.

## Usage/Examples

```go
package main

import (
    "github.com/canack/sobot/pkg"
)

func main(){
    // User creds
    username := "nickname"
    password := "secretpass123"

    // Post parameters
    file := "/home/user/Desktop/photo.jpg"
    comment := "Have a nice day!"

    // you can activate debugging with Share(true)
    // Sharing on Instagram
    sobot.Instagram(username, password).SetFile(file).SetCaption(comment).Share(true)
	
    // Sharing on Twitter
    sobot.Twitter(username, password).SetFile(file).SetCaption(comment).Share(true)
}
```

---
### By default, Rod will disable the browser's UI to maximize the performance. But when developing an automation task we usually care more about the ease of debugging. Rod provides a lot of solutions to help you debug the code.

 Let's create a ".rod" config file under the current working directory. The content is:

```show```

---

The automated operations are too fast for human eyes to catch, to debug them we usually enable the visual trace config, let's update the ".rod" file:

```
show
trace
```
---

## Routemap

- [+] Instagram support (currently stable-beta)

- [+] Twitter support (currently stable-beta)


## Contributors

- [@canack](https://www.github.com/canack)