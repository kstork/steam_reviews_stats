package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetReviews(appid int) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?filter=updated&cursor=*&num_per_page=100", appid)

	// TODO: put that data into a struct, then start over and replace "cursor=*" the * with the id from "cursor" at the end of the response, do it as long as there are reviews and put it all into a nice struct, then give the struct back

	return http.Get(url)
}

func main() {

	response, err := GetReviews(427520)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
