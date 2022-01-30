## golang CLI tool created with cobra

### Part 1: create a root and child command that uses curl get from an api

#### 1. set up the repo structure
- `go mod init randomjoke`
- `cobra init`

#### 2. set up the base command and test
- `cobra add get`: this will create a child command for your root command
- the rest of the steps are to finish the function that will execute if this command runs

#### 3. write the request 
- use this template for request and response

```go
import (
  "net/http"
  "io/ioutil"
)

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
```

#### 4. unmarshal its response from json format to go struct
- This tool instantly converts JSON into a Go type definition: https://mholt.github.io/json-to-go/
- The template is as follows:

```go
import "encoding/json"

type Joke struct {
  ID string `json:"id"`
  Joke string `json:"joke"`
  Status int `json:"status"`
}

func jsonUnmarshal(responseBytes []byte) string {
  joke := Joke{}

  if err := json.Unmarshal(responseBytes, &joke); err != nil {
    fmt.Printf("Could not unmarshal reponseBytes. %v", err)
  }
  
  return joke.Joke
}

```


#### 5. combine the above two parts together into one function, and put it in the `Run` field under getCmd

```go
// the main function goes under Run
var getCmd = &cobra.Command{
  Use: "get",
  Short: "get a randomjoke from the jokeapi",
  Long: `get a randomjoke from the jokeapi`,
  Run: func(cmd \*cobra.Command, args []string) {
    getJoke()
  },
}
```

### block that checks if a flag exists 
#### 6. Add a flag to pass a keyword parameter

In the subcommand.go file, add the below code to add a flag to the command

```go
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a randomjoke from the jokeapi",
	Long: `get a randomjoke from the jokeapi`,
	Run: func(cmd *cobra.Command, args []string) {
		keyword, _ := cmd.Flags().GetString("keyword")

		if keyword != "" {
      //function that works when keyword parameter is passed in
			getJokeWithKeyword(keyword)
		} else {
      //function for the original joke
			getJoke()
		}
	},
}

//add getCmd.PersistentFlags for flags that applies to both the root and child command
func init() {
	rootCmd.AddCommand(getCmd)
  getCmd.PersistentFlags().String("keyword", "", "search the jokes by keyword")
}
```
###### test if the above works with a fmt print function

```go
func getJokeWithKeyword(keyword string) {
  fmt.Printf("test if %v works",keyword)
}
```
###### make the function work with the search keyword

This part is very simple, just match the curl command in the http request and a different unmarshal command to marshal the returned response bytes



### how to use this cli locally
run `go run main.go get` or `go build main.go` then run  `./main get`
