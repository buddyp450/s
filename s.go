package main

import (
	"fmt"
	"os"

	// Load necessary providers.
	_ "github.com/buddyp450/s/providers/amazon"
	_ "github.com/buddyp450/s/providers/bing"
	_ "github.com/buddyp450/s/providers/digg"
	_ "github.com/buddyp450/s/providers/dockerhub"
	_ "github.com/buddyp450/s/providers/duckduckgo"
	_ "github.com/buddyp450/s/providers/github"
	_ "github.com/buddyp450/s/providers/go"
	_ "github.com/buddyp450/s/providers/google"
	_ "github.com/buddyp450/s/providers/npm"
	_ "github.com/buddyp450/s/providers/npmsearch"
	_ "github.com/buddyp450/s/providers/pinterest"
	_ "github.com/buddyp450/s/providers/reddit"
	_ "github.com/buddyp450/s/providers/soundcloud"
	_ "github.com/buddyp450/s/providers/stackoverflow"
	_ "github.com/buddyp450/s/providers/twitter"
	_ "github.com/buddyp450/s/providers/wikipedia"
	_ "github.com/buddyp450/s/providers/yahoo"
	_ "github.com/buddyp450/s/providers/youtube"
	"github.com/buddyp450/s/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.SearchCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
