package main

import(
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"strings"
)

func main() {
	//fileName is a pointer to string, so *fileName is the value
	fileName := flag.String("file", "problems.csv", "a csv file that has the format of 'question, answer'")
	flag.Parse()
	fmt.Println("the csv file name is: ", *fileName)
	
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *fileName))
	}
    r := csv.NewReader(file)
	//parse the csv files
    lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseLines(lines)

	// count the number of right answered problem
	counter := 0

	for i,p := range problems {
	    //this prints out the questions
		fmt.Printf("Problem #%d: %s = ?\n", i+1, p.question)
		//this reads the user input
		var answer string
		fmt.Scanf("%s\n", &answer)
        //validate if user's input is the right answer
		if answer ==p.answer {
			fmt.Println("Correct")
			counter++
		} else {
			fmt.Println("Wrong")
		}
	}
	fmt.Printf("the number of your correct answers in total is %d\n", counter)
	fmt.Printf("there are in total %d questions\n", len(lines))
}

//this is similar to log.Fatal(err) function, but more explicitly
func exit(msg string){
	fmt.Println(msg)
    os.Exit(1)
}

type problem struct {
	question string
	answer string
}

func parseLines(lines [][]string)[]problem {
	//initialise an array
	ret := make([]problem, len(lines))
	//fill in the array of struct
    for i, line := range lines {
		ret[i] = problem {
			//the first colume is question, and second column is answer
			question: line[0],
			//use the trimspace to get rid of the answer space
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
