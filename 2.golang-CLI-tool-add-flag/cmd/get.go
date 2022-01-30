/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/spf13/cobra"
  "math/rand"
  "time"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a randomjoke from the jokeapi",
	Long: `get a randomjoke from the jokeapi`,
	Run: func(cmd *cobra.Command, args []string) {
		keyword, _ := cmd.Flags().GetString("keyword")

		if keyword != "" {
			getJokeWithKeyword(keyword)
		} else {
			getJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
  getCmd.PersistentFlags().String("keyword", "", "search the jokes by keyword")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeRequest(url)
  joke := jsonUnmarshal(responseBytes)
	fmt.Println(joke)
}

//golang http request 
func getJokeRequest(url string) []byte {
  request,err := http.NewRequest(
    http.MethodGet,  //get method
    url, //url
    nil, //body
  )

  //Add a header to tell the API we want our data returned as JSON
  request.Header.Add("Accept", "application/json")
  //Add a custom User-Agent header to tell the API maintainers how we’re using their API
  request.Header.Add("User-Agent", "https://github.com/yanglyu520/golang-projects")
  
  //err handling for request not going through
  if err != nil {
    log.Printf("Could not make a request to get a random joke. %v", err)
  }

  //response
  response, err := http.DefaultClient.Do(request)
  //err handling for response
  if err != nil {
    log.Printf("Could not make a request. %v", err)
  }

  //response body
  responseBytes, err := ioutil.ReadAll(response.Body)

  //response body err handling
  if err != nil {
    log.Printf("Could not read the response body. %v", err)
  }

  return responseBytes
}

//unmarshal the joke from json to string
type Joke struct {
  ID string `json:"id"`
  Joke string `json:"joke"`
  Status int `json:"status"`
}

func jsonUnmarshal(responseBytes []byte) string {
  joke := Joke{}

  if err := json.Unmarshal(responseBytes, &joke); err != nil {
    fmt.Printf("Could not unmarshal responseBytes. %v", err)
  }
  
  return joke.Joke
}

//----function to get a joke with keyword------
func getJokeWithKeyword(keyword string) () {
  url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", keyword)
  responseBytes := getJokeRequest(url)
  joke := jsonUnmarshalWithKeyword(responseBytes)
  fmt.Println(joke)
}

// struct for joke api that returns json for joke with keyword
type JokeWithKeyword struct {
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
	Results      []Result `json:"results"`
	SearchTerm string `json:"search_term"`
	Status     int    `json:"status"`
	TotalJokes int    `json:"total_jokes"`
	TotalPages int    `json:"total_pages"`
}

type Result struct {
  ID   string `json:"id"`
	Joke string `json:"joke"`
}

//func that unmarshal the joke json to the golang struct with keyword
func jsonUnmarshalWithKeyword(responseBytes []byte) string {
  rand.Seed(time.Now().Unix())
  joke := JokeWithKeyword{}
  
  if err := json.Unmarshal(responseBytes, &joke); err != nil {
    fmt.Printf("Could not unmarshal responseBytes. %v", err)
  } 
  
  length := joke.TotalJokes
  min := 0
  max := length - 1
  
  if length <= 0 {
    return "Could not find joke with this keyword"
  } else {
    randomNum := min + rand.Intn(max-min)
    return joke.Results[randomNum].Joke
  }
  
}
