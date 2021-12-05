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

var watch_twitter bool

func (user *AccountInfo) firstTwitterRun() {

	os.MkdirAll("src/account_data/twitter/"+user.Username, os.ModePerm)
	log.Println("First run for ", user.Username, " on ", user.Pname)

	page := rod.New().MustConnect().MustPage("about:blank")

	page.MustEmulate(rod_device)

	page.MustNavigate("https://twitter.com/i/flow/login").MustWaitLoad()
	// For debugging
	end_debug := make(chan bool)

	if watch_twitter {
		go user.watchBrowser(page, end_debug)
	}

	time.Sleep(time.Second * 3)

	page.MustElementX("/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[5]/label/div/div[2]/div/input").MustInput(user.Username)
	page.MustElementX("/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[6]/div").MustClick()
	page.MustElementX("/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[3]/div/label/div/div[2]/div[1]/input").MustInput(user.Password)
	page.MustElementX("/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div").MustClick()

	// This section makes the first login to the account and reports if the password is wrong
	time.Sleep(time.Second * 7)

	_, err := page.Timeout(time.Second * 5).ElementX("/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[1]/div/div[1]/div[1]/span")

	if err == nil {
		page.Close()
		log.Println("Username or password incorrect")
		return
	}
	// This section makes the first login to the account and reports if the password is wrong

	page.MustWaitLoad()
	log.Println("Wait 3 second for get cookies successful.")
	time.Sleep(time.Second * 3)

	f, err := os.Create("src/account_data/twitter/" + user.Username + "/cookie.dat")

	if err != nil {
		log.Fatal("We got an error at creating cookie data file\n", err)
	}

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
	user.cookieTwitterRun()

}

func (user *AccountInfo) cookieTwitterRun() {
	log.Println("Found cookie for ", user.Username)

	// first section
	page := rod.New().MustConnect().MustPage("about:blank")

	page.MustEmulate(rod_device)

	file, err := os.Open("src/account_data/twitter/" + user.Username + "/cookie.dat")

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

	page.MustNavigate("https://twitter.com").MustWaitLoad()

	// For debugging
	end_debug := make(chan bool)

	if watch_twitter {
		go user.watchBrowser(page, end_debug)
	}

	time.Sleep(time.Second * 5)

	page.MustElementX("/html/body/div/div/div/div[2]/main/div/div/div/div/div/div[2]/div/div[2]/div[1]/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/label/div[1]/div/div/div/div/div[2]/div/div/div/div").MustClick()

	page.MustElementX("/html/body/div/div/div/div[2]/main/div/div/div/div/div/div[2]/div/div[2]/div[1]/div/div/div/div[2]/div[1]/div/div/div/div/div/div/div/div/div/label/div[1]/div/div/div/div/div[2]/div/div/div/div").MustInput(user.Caption)
	time.Sleep(time.Second * 3)
	page.MustElementX("/html/body/div/div/div/div[2]/main/div/div/div/div/div/div[2]/div/div[2]/div[1]/div/div/div/div[2]/div[3]/div/div/div[1]/input").MustSetFiles(user.FilePath)
	time.Sleep(time.Second * 3)
	page.MustElementX("/html/body/div/div/div/div[2]/main/div/div/div/div/div/div[2]/div/div[2]/div[1]/div/div/div/div[2]/div[3]/div/div/div[2]/div[3]/div").MustClick()

	log.Println("All process are done, waiting 15 seconds")
	time.Sleep(time.Second * 15)

	if watch_instagram {
		end_debug <- true // for stop saving screenshots
	}

	log.Println("Everything ok, shutting down browser.")
	page.Close()

	log.Println("Login with Cookie and share successfully")
}

func (user *AccountInfo) shareTwitter(debug bool) {

	_, err := os.ReadFile("src/account_data/twitter/" + user.Username + "/cookie.dat")

	watch_twitter = debug

	if err != nil {
		user.firstTwitterRun()
	} else {
		user.cookieTwitterRun()
	}

}
