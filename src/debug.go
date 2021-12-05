package sobot

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-rod/rod"
)

func (user *AccountInfo) watchBrowser(page *rod.Page, end chan bool) {
	dt := time.Now()
	date := dt.Format("02-01-2006_15.04.05")

	directory := "screenshots/" + user.Username + "/" + user.Pname + "/" + date
	err := os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		log.Println("We can't create a document for debugging")
	}

	final := false

	go func() {

		if <-end {
			final = true
		}
	}()

	for i := 0; ; i++ {
		if final {
			break
		}

		num := strconv.Itoa(i)
		name := directory + "/debug_" + num + ".png"
		page.MustScreenshot(name)
		log.Println("Captured:", name)
		time.Sleep(time.Second)

	}
}
