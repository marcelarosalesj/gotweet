# gotweet

## Objective
The `gotweet` tool shall be a CLI utility that let me do basic Twitter operations. For example, it should show me tweets from my home page and let me tweet too.

## Requirements and first steps
- You need a Twitter Developer account to have access to the Twitter API.
- From the Developer Portal you can either Projects or Standalone Apps.
- I created a Standalone App.
- It will provide you the `API_KEY` and the `API_KEY_SECRET`.
- Besides, in order to authenticate your transactions you need to generate the `ACCESS_TOKEN` and the `ACCESS_TOKEN_KEY`, also in the Developer Portal.

## Getting started with Twitter API
Check out Twitter API documentation [1].
```
# twurl example for v1.1
sudo gem install twurl
twurl -u "$USER_gotweet" -c $API_KEY -s $API_KEY_SECRET -a $ACCESS_TOKEN -S $ACCESS_TOKEN_SECRET  -X GET "/1.1/search/tweets.json?q=%40twitterapi"
```

## Go code and implementation
I used tutorial and code from resource [2] to get started with the tool I want to build. It describes how to authenticate with Twitter and how to do search queries using Twitter API. To accomplish this, the tutorial uses [oauth1 library](https://github.com/dghubble/oauth1) for the authentication part and  [go-twitter library](https://github.com/dghubble/go-twitter/) as the Go client for the API. As of now go-twitter library supports Twitter API v1.1

Then, tutorial [3] has information about how to build a CLI tool using go.

## References
* [1][Twitter API v1.1 documentation](https://developer.twitter.com/en/docs/twitter-api/v1)
* [2] [Using the Twitter API to get tweets](https://laptrinhx.com/using-the-twitter-api-to-get-tweets-921505982/)
* [3] [Building a Simple CLI Tool with Golang](https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/)