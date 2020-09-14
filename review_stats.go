// TODO: How do I deal with cursor pagination?

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Results struct {
	Success       int           `json:"success"`
	Query_Summary Query_Summary `json:"query_summary"`
	Reviews       []Reviews     `json:"reviews"`
	Cursor        string        `json:"cursor"`
}

type Query_Summary struct {
	Num_Reviews       int    `json:"num_reviews"`
	Review_Score      int    `json:"review_score"`
	Review_Score_Desc string `json:"review_score_desc"`
	Total_Positive    int    `json:"total_positive"`
	Total_Negative    int    `json:"total_negative"`
	Total_Reviews     int    `json:"total_reviews"`
}

type Reviews struct {
	Recommendationid            string `json:"recommendationid"`
	Author                      Author `json:"author"`
	Language                    string `json:"language"`
	Review                      string `json:"review"`
	Timestamp_Created           int    `json:"timestamp_created"`
	Timestamp_Updated           int    `json:"timestamp_updated"`
	Voted_Up                    bool   `json:"voted_up"`
	Votes_Up                    int    `json:"votes_up"`
	Votes_Funny                 int    `json:"votes_funny"`
	Weighted_Vote_Score         int    `json:"weighted_vote_score"`
	Comment_Count               int    `json:"comment_count"`
	Steam_Purchase              bool   `json:"steam_purchase"`
	Received_For_Free           bool   `json:"received_for_free"`
	Written_During_Early_access bool   `json:"written_during_early_access"`
}

type Author struct {
	Steamid                 string `json:"steamid"`
	Num_Games_Owned         int    `json:"num_games_owned"`
	Num_Reviews             int    `json:"num_reviews"`
	Playtime_Forever        int    `json:"playtime_forever"`
	Playtime_Last_Two_Weeks int    `json:"playtime_last_two_weeks"`
	Playtime_At_Review      int    `json:"playtime_at_review"`
	Last_Played             int    `json:"last_played"`
}

func GetReviews(appid int, cursor string) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?json=1&filter=updated&cursor=%v&num_per_page=100", appid, cursor)

	// Steam API docs https://partner.steamgames.com/doc/store/getreviews
	// TODO: Put all stuff here, including filling the structs, iterating through the cursors, and filtering out 0 minute playtime reviews! Get it out of the main function
	// TODO: return the struct? Do I need pointers?

	return http.Get(url)
}

func main() {

	response, err := GetReviews(201510, "*")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results Results

	err = json.Unmarshal(responseData, &results)
	if err != nil {
		fmt.Print(err.Error())
	}

	var sum_seconds int

	for i := 0; i < len(results.Reviews); i++ {
		if !results.Reviews[i].Voted_Up {
			sum_seconds += results.Reviews[i].Author.Playtime_Forever
			fmt.Println("Text: " + strconv.Itoa((results.Reviews[i].Author.Playtime_Forever / 60)) + " minutes")
		}
	}

	sum_minutes := sum_seconds / 60
	average_minutes := sum_minutes / len(results.Reviews)

	fmt.Printf("Average playtime of reviews is about %d minutes. \n", average_minutes)
}
