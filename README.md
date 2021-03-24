# gotweet

## Requirements and first steps
- You need a Twitter Developer account to have access to the Twitter API.
- From the Developer Portal you can either Projects or Standalone Apps.
- I created a Standalone App.
- It will provide you the `API_KEY` and the `API_KEY_SECRET`.
- Besides, in order to authenticate your transactions you need to generate the `ACCESS_TOKEN` and the `ACCESS_TOKEN_KEY`, also in the Developer Portal.


## Getting started with Twitter API
```
# twurl example for v1.1
sudo gem install twurl
twurl -u "$USER_gotweet" -c $API_KEY -s $API_KEY_SECRET -a $ACCESS_TOKEN -S $ACCESS_TOKEN_SECRET  -X GET "/1.1/search/tweets.json?q=%40twitterapi"
```

## Go code
I use tutorial [1] as an example for this tool.
```
# Run the code
go run main.go
```
Note that for unit testing the file needs to be called <name>_test.go
```
# Run the unit testing
go test
```

## References

* [1] [Using the Twitter API to get tweets](https://laptrinhx.com/using-the-twitter-api-to-get-tweets-921505982/)
