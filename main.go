package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Credentials stores the credentials to log into the developer account
type Credentials struct {
	ConsumerKey       string `json:"consumerkey"`
	ConsumerSecret    string `json:"consumersecret"`
	AccessToken       string `json:"accesstoken"`
	AccessTokenSecret string `json:"accesstokensecret"`
}

//Data stores data
type Data struct {
	Days                int `json:"days"`
	TotalQuestionSolved int `json:"totalquestionsolved"`
}

func getTwitterClient(cred *Credentials) (*twitter.Client, error) {
	//Pass your Api key
	config := oauth1.NewConfig(cred.ConsumerKey, cred.ConsumerSecret)
	//Pass your Access key
	token := oauth1.NewToken(cred.AccessToken, cred.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		return nil, err
	}
	fmt.Printf("Login Succesfull!! \n\tID:\t%v\n\tName: \t%v\n", user.ScreenName, user.Name)
	return client, nil

}

func parseHTML(url string) int {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	//check statuscode
	if resp.StatusCode != 200 {
		log.Fatalln("Response code: ", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result string
	doc.Find(".progress-bar-success").Each(func(i int, s *goquery.Selection) {
		if i == 3 {
			result = strings.TrimSpace(s.Text())
			result = strings.Split(result, "/")[0]
			result = strings.TrimSpace(result)
		}
	})
	solved, _ := strconv.Atoi(result)
	return solved
}

func tweetTweet(client *twitter.Client, url string) {

	solved := parseHTML(url)
	data := Data{}

	fl, err := os.Open("./data.json")
	if err != nil {
		log.Fatalln(err)
	}

	bycred, err := ioutil.ReadAll(fl)
	if err != nil {
		log.Fatalln(err)
	}
	fl.Close()

	err = json.Unmarshal(bycred, &data)
	if err != nil {
		log.Fatalln(err)
	}

	msg := "Day: " + strconv.Itoa(data.Days) + "/100\n\n"
	badMsgs := []string{
		"I broke my leetcoding streak :(\n\n",
		"I forgot to solve problems at @LeetCode today. :( \nBut I will make it up tommorow.\n\n",
		"Arghhhhh...\n\n I forgot to solve problems at @LeetCode.\n\n",
		"Didn't solved any problems today. :(\n\n",
		"I am lazzzzzzzyyyyy.\n\nI haven't solved a single question from @LeetCode. :(\n\n",
		"Failed to solve questions at @LeetCode.\n\n :(",
	}
	goodMsgs := []string{
		"Hurray!!!\n\n I am on step closer to my goal. Solved questions at @LeetCode.\n\n",
		"Another day spent solving questions at @LeetCode.\n\n",
		"Solved questions at @LeetCode, I am on step closer to my goal.\n\n",
		"Wooo-Hooo! Solving DS-Algo problems is so exciting. Best way to practice language fundamentals.\n\n",
		"Solved problems in @LeetCode using #golang and #cpp.\n\n",
		"Leetcoding is best way to prepare for you tech interviews. Solved few questions at @LeetCode today.\n\n",
	}

	if solved == data.TotalQuestionSolved {
		indx := rand.Intn(len(badMsgs))
		pick := badMsgs[indx]
		msg += pick
		msg += "Question solved : " + strconv.Itoa(solved-data.TotalQuestionSolved) + " 😔😔😔😔\n\n"
	} else if solved > data.TotalQuestionSolved {
		indx := rand.Intn(len(goodMsgs))
		pick := goodMsgs[indx]
		msg += pick
		msg += "Questions solved : " + strconv.Itoa(solved-data.TotalQuestionSolved) + " 😃😃🎉🎉\n\n"
	}

	msg += "Join me in #100daysofleetcode challenge and start prepairing for your upcoming interviews.\n\n"
	msg += "#leetcode #100daysofcode #CodeNewbie"

	if data.Days == 0 {
		msg = "I am commiting to the #100daysofcode challenge starting from today. I will be solving data structures and algorithm at @LeetCode. \n\nI will track my progress using a tweeter bot. \n\n#CodeNewbie #golang #100daysofleetcode #100daysofcode #interviewprep"
	}

	_, _, err = client.Statuses.Update(msg, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Day:", data.Days, " -- Msg Sent")
	log.Println(msg)

	data.Days++
	data.TotalQuestionSolved = solved

	bycred, err = json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	err = ioutil.WriteFile("./data.json", bycred, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	fmt.Println("Go-Tweet-Bot: LeetCode v0.01")

	url := "https://leetcode.com/sonukumarsaw"

	fl, err := os.Open("./cred.json")
	if err != nil {
		log.Fatalln(err)
	}

	bycred, err := ioutil.ReadAll(fl)
	if err != nil {
		log.Fatalln(err)
	}

	creds := Credentials{}
	err = json.Unmarshal(bycred, &creds)
	if err != nil {
		log.Println(err)
	}

	for {

		client, err1 := getTwitterClient(&creds)

		if err1 != nil {
			log.Fatalln(err1)
		}

		tweetTweet(client, url)
		time.Sleep(24 * time.Hour)

	}

}
