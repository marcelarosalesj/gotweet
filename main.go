package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type TweetData struct {
	Tweet        string
	LikeCount    int
	RetweetCount int
}

func main() {
	creds := Credentials{
		ConsumerKey:       os.Getenv("API_KEY"),
		ConsumerSecret:    os.Getenv("API_KEY_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	demoPtr := flag.Bool("demo", false, "Show tweets related to Formula 1")
	homePtr := flag.Bool("home", false, "Displays home page tweets")
	flag.Parse()

	if *demoPtr {
		searchTweets(client)
	}
	if *homePtr {
		err := showHomePage(client)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Bye!")

}

func getClient(creds *Credentials) (*twitter.Client, error) {

	// These values are the API key and API key secret
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// These values are the consumer access token and consumer access token secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verify := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verify)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// print out the Twitter handle of the account we have used to authenticate with
	fmt.Println("Successfully authenticated using the following account : ", user.ScreenName)
	return client, nil
}

func searchTweets(client *twitter.Client) error {
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "Formula 1",
		Lang:  "en",
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, v := range search.Statuses {
		tweet := TweetData{
			Tweet:        v.Text,
			LikeCount:    v.FavoriteCount,
			RetweetCount: v.RetweetCount,
		}
		fmt.Printf("%+v\n", tweet)
	}
	return nil
}

func showHomePage(client *twitter.Client) error {
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 5,
	})
	if err != nil {
		return err
	}
	for _, tweet := range tweets {
		displayTweet(tweet)
	}
	return nil
}

func displayTweet(tweet twitter.Tweet) {
	fmt.Println("----------")
	fmt.Println("User: " + tweet.User.Name)
	fmt.Println("Tweet:" + tweet.Text)
	fmt.Println("----------")
}
func Sum(x, y int) int {
	return x + y
}
