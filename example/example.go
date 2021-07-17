package main

import (
	"fmt"
	"time"

	"github.com/pioneer-io/go_sdk"
)

/*
run a simple program like this to test out the SDK
*/

func main() {
	// Initialize an SDK client
	client := go_sdk.InitMember("http://localhost:3030/features", "JazzyElksRule")

	// connect SDK client to Scout to listen for SSE updates
	client.Connect()
	client.Listen()

	// this example supposes that you have an existing flag
	// called 'test this flag' that is toggled off

	fmt.Println(client.Get("test this flag")) // false

	time.Sleep(12 * time.Second)
	// wait so you can toggle "test this flag" von ia UI
	// and create "a_new_flag" and toggle it on

	// after we're finished sleeping, we can log the updated ruleset
	fmt.Println(client.Get("a_new_flag"))     // true
	fmt.Println(client.Get("test this flag")) // true

}
