package main

import (
	"context"
	"fmt"
	// "log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)


func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6438967140116-6422028971239-2FDVF5RTzTycMgZM43hAj5IB")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06CWUR1BCL-6429888359718-1642825a779a637fbefb15117fcd229b50a9a592399d87c1056512d7ff9782ba")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my job is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example:	 "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("Year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				println("error")
			}
			age := 2021-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		fmt.Println(err)
	}
}