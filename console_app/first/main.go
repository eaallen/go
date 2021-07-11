/*
	Code lesson from
		https://www.codementor.io/@opiumated/write-a-command-line-application-with-go-a05wdamn6
*/
package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var options struct {
	postCount int
	version   bool
}

func init() {
	flag.IntVar(&options.postCount, "count", 1, "Enter the number of posts you would love to see")
	flag.BoolVar(&options.version, "version", false, "View the version of this application")
	flag.Parse()
}

func main() {
	var posts = []string{"Using Docker and Docker Compose for Local Development and Small Deployments", "Trying Clean Architecture on Golang", "Understanding Depenedency Management and Package Management in Golang", "Dancing with Go's Mutexes", "Selenium: easy as pie", "Go Best Practises - Testing"}

	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[post_count]\n", os.Args[0])
		flag.PrintDefaults()
	}

	intPtr := &options.postCount

	// NO areguments so print usage
	if len(os.Args[1:]) == 0 {
		Usage()
		os.Exit(1)
	}

	// handle count
	if intPtr != nil {
		if *intPtr > len(posts) {
			fmt.Printf("%s\n", "Number of post specified is greater than what we currently have...")
		}
		buildPost := getPostByCount(posts, *intPtr)
		for post := range buildPost {
			fmt.Printf("%s\n", buildPost[post])
		}
	}

	// handle verison
	if options.version {
		fmt.Printf("%s", "goCmd installed version 0.0.1\n")
	}

	// end after making it through potential prints
	os.Exit(0)

}

func getPostByCount(posts []string, count int) []string {
	buildPost := []string{}
	for post := 0; post < count; post++ {
		buildPost = append(buildPost, posts[post])
	}
	return buildPost
}
