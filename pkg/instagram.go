package sobot

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

var (
	watch_instagram bool
)

func (user *AccountInfo) firstInstagramRun() {

	os.MkdirAll("account_data/instagram/"+user.Username, os.ModePerm)
	log.Println("First run for ", user.Username, " on ", user.Pname)

	// first section
	page := rod.New().SlowMotion(time.Second * 2).MustConnect().MustPage("about:blank")

	page.MustEmulate(rod_device)

	page.MustNavigate("https://instagram.com").MustWaitLoad()

	// For debugging
	end_debug := make(chan bool)

	if watch_instagram {
		go user.watchBrowser(page, end_debug)
	}
	time.Sleep(time.Second * 3)

	// username input
	page.MustElementX("/html/body/div[1]/section/main/article/div[2]/div[1]/div/form/div/div[1]/div/label/input").MustInput(user.Username)

	// password input
	page.MustElementX("/html/body/div[1]/section/main/article/div[2]/div[1]/div/form/div/div[2]/div/label/input").MustInput(user.Password)

	// login button
	page.MustElementX("/html/body/div[1]/section/main/article/div[2]/div[1]/div/form/div/div[3]/button").MustClick()

	// This section makes the first login to the account and reports if the password is wrong
	time.Sleep(time.Second * 7)
	_, err := page.Timeout(time.Second * 5).ElementX("/html/body/div[1]/section/main/article/div[2]/div[1]/div/form/div[1]/div[1]/div/label/input")

	if err == nil {
		page.Close()
		log.Println("Username or password incorrect")
		return
	}
	// This section makes the first login to the account and reports if the password is wrong

	page.MustWaitLoad()
	log.Println("Wait 3 second for get cookies successful.")
	time.Sleep(time.Second * 3)

	f, err := os.Create("account_data/instagram/" + user.Username + "/cookie.dat")

	if err != nil {
		log.Fatal("We got an error at creating cookie data file\n", err)
	}

	// Write cookies into file
	for _, v := range page.MustCookies() {
		_, err2 := f.WriteString(v.Name + "#" + v.Value + "#" + v.Domain + "\n")

		if err2 != nil {
			log.Fatal("Cookies doesn't writed on file\n", err2)
		}
	}
	log.Println("Cookies saved.")
	f.Close()

	if watch_instagram {
		end_debug <- true // for stop saving screenshots
	}

	page.Close()

	log.Println("Login successful and cookies registered")
	user.cookieInstagramRun()

}

func (user *AccountInfo) cookieInstagramRun() {
	log.Println("Found cookie for ", user.Username)

	// first section
	page := rod.New().SlowMotion(time.Second * 2).MustConnect().MustPage("about:blank")

	page.MustEmulate(rod_device)

	file, err := os.Open("account_data/instagram/" + user.Username + "/cookie.dat")

	if err != nil {
		log.Fatalf(user.Username, user.Pname, " cookie read error: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		saved_dat := strings.Split(scanner.Text(), "#")
		page.MustSetCookies([]*proto.NetworkCookieParam{{
			Name:   saved_dat[0],
			Value:  saved_dat[1],
			Domain: saved_dat[2],
		}}...)

	}

	file.Close()
	log.Println("Cookies loaded")

	page.MustNavigate("https://instagram.com").MustWaitLoad()
	// For debugging
	end_debug := make(chan bool)

	if watch_instagram {
		go user.watchBrowser(page, end_debug)
	}
	time.Sleep(time.Second * 5)

	// Pop-up control
	go func(page *rod.Page) {
		page.MustElementX("/html/body/div[5]/div/div/div/div[3]/button[2]").MustClick()
	}(page)

	// Add post button
	page.MustElementX("/html/body/div[1]/section/nav/div[2]/div/div/div[3]/div/div[3]/div/button").MustClick()

	// Set files
	page.MustElementX("/html/body/div[8]/div[2]/div/div/div/div[2]/div[1]/form/input").MustSetFiles(user.FilePath)

	// Next button
	page.MustElementX("/html/body/div[6]/div[2]/div/div/div/div[1]/div/div/div[2]/div/button").MustClick()

	// Next button
	page.MustElementX("/html/body/div[6]/div[2]/div/div/div/div[1]/div/div/div[2]/div/button").MustClick()

	// Write a caption
	page.MustElementX("/html/body/div[6]/div[2]/div/div/div/div[2]/div[2]/div/div/div/div[2]/div[1]/textarea").MustInput(user.Caption)

	// Share button
	page.MustElementX("/html/body/div[6]/div[2]/div/div/div/div[1]/div/div/div[2]/div/button").MustClick()

	log.Println("All process are done, waiting 15 seconds")
	time.Sleep(time.Second * 15)

	if watch_instagram {
		end_debug <- true // for stop saving screenshots
	}
	log.Println("Everything ok, shutting down browser.")
	page.Close()

	log.Println("Login with Cookie and share successfully")

}

func (user *AccountInfo) shareInstagram(debug bool) {

	_, err := os.ReadFile("account_data/instagram/" + user.Username + "/cookie.dat")
	watch_instagram = debug
	if err != nil {

		user.firstInstagramRun()

	} else {

		user.cookieInstagramRun()

	}

}
