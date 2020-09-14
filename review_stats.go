package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func GetReviews(appid int) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://store.steampowered.com/appreviews/%d?json=1&filter=updated&cursor=*&num_per_page=5", appid)

	// TODO: put that data into a struct, then start over and replace "cursor=*" the * with the id from "cursor" at the end of the response, do it as long as there are reviews and put it all into a nice struct, then give the struct back
	// Steam API docs https://partner.steamgames.com/doc/store/getreviews

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

	var results Results

	json.Unmarshal(responseData, &results)

	for i := 0; i < len(results.Reviews); i++ {
		fmt.Println("Text: " + results.Reviews[i].Review)
	}
}
