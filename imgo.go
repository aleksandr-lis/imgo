package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Wrong, you need to do it like this:\n./imgo image_url")
		os.Exit(1)
	}
	page_url := args[1]

	if strings.Contains(page_url, "instagram.com") {
		fmt.Println("Ok, it's an Instagram")
		proceed_instagram(page_url)

	} else if strings.Contains(page_url, "pinterest.") {
		fmt.Println("Ok, it's a Pinterest")
		proceed_pinterest(page_url)

	} else if strings.Contains(page_url, "facebook.com") {
		fmt.Println("Ok, it's a Facebook")
		proceed_facebook(page_url)

	} else {
		fmt.Println("I can't download images from this service, sorry")
	}

	fmt.Println("I'm done!'")
}

func proceed_instagram(page_url string) {
	html := string(get_content(page_url))
	re, _ := regexp.Compile("<meta property=\"og:(image|video)\" content=\"(https://(.+)/([0-9a-z_\\.]+)?(.+))\" />")
	index := re.FindAllStringSubmatchIndex(html, -1)
	if len(index) == 0 {
		fmt.Println("Wrong url!")
		os.Exit(1)
	}

	for _, i := range index {
		file_url := html[i[4]:i[5]]
		file_name := html[i[8]:i[9]]
		save_media(file_url, file_name)
	}
}

func proceed_pinterest(page_url string) {
	html := string(get_content(page_url))
	re, _ := regexp.Compile("<meta property=\"og:image\" name=\"og:image\" content=\"(https://i.pinimg.com/([0-9a-z\\/]+)/([0-9a-z\\.\\-]+))\" data-app=\"true\"/>")
	index := re.FindStringSubmatchIndex(html)
	if len(index) == 0 {
		fmt.Println("Wrong or private url!")
		os.Exit(1)
	}

	file_url := html[index[2]:index[3]]
	file_name := html[index[6]:index[7]]
	save_media(file_url, file_name)
}

func proceed_facebook(page_url string) {
	html := string(get_content(page_url))
	re, _ := regexp.Compile("<meta property=\"og:image\" content=\"(https://([0-9a-z\\.\\-\\/]+)/([0-9a-z_\\.]+)([?0-9a-zA-Z\\-\\.&;_=]+))\" />")
	index := re.FindStringSubmatchIndex(html)
	if len(index) == 0 {
		fmt.Println("Wrong or private url!")
		os.Exit(1)
	}

	file_url := strings.ReplaceAll(html[index[2]:index[3]], "&amp;", "&")
	file_name := html[index[6]:index[7]]
	save_media(file_url, file_name)
}

func save_media(file_url string, file_name string) {
	fmt.Println("Getting file...")
	file_content := get_content(file_url)
	err := ioutil.WriteFile("files/"+file_name, file_content, 0644)
	check(err)

	fmt.Println("Downloaded file saved in: files/" + file_name)
}

func get_content(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	dataInBytes, err := ioutil.ReadAll(response.Body)
	return dataInBytes
}

func check(e error) {
	if e != nil {
		fmt.Println("Wrong url!")
	}
}
