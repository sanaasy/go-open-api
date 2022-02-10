// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Enter something: ")

// 	reader := bufio.NewReader(os.Stdin)
// 	text, _ := reader.ReadString('\n')
// 	text = strings.Replace(text, "\n", "", -1)

// 	fmt.Println(text)

// 	os.Exit(0)
// }

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	api, err := ioutil.ReadFile("apikey.txt")
	apiKey := string(api)
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	text := getText()

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	// otherwise completion
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{text},
		MaxTokens: gpt3.IntPtr(60),
		Stop:      []string{"."},
		Echo:      true,
	})
	// scope issue - move err and resp variables to outside if block
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text + ".")
}

func getText() string {
	fmt.Println("Enter something: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}
