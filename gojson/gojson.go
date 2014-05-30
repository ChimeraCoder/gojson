// gojson generates go struct defintions from JSON documents
//
// Reads from stdin and prints to stdout
//
// Example:
// 	curl -s https://api.github.com/repos/chimeracoder/gojson | gojson -name=Repository
//
// Output:
// 	package main
//
// 	type User struct {
// 		AvatarURL         string      `json:"avatar_url"`
// 		Bio               interface{} `json:"bio"`
// 		Blog              string      `json:"blog"`
// 		Company           string      `json:"company"`
// 		CreatedAt         string      `json:"created_at"`
// 		Email             string      `json:"email"`
// 		EventsURL         string      `json:"events_url"`
// 		Followers         float64     `json:"followers"`
// 		FollowersURL      string      `json:"followers_url"`
// 		Following         float64     `json:"following"`
// 		FollowingURL      string      `json:"following_url"`
// 		GistsURL          string      `json:"gists_url"`
// 		GravatarID        string      `json:"gravatar_id"`
// 		Hireable          bool        `json:"hireable"`
// 		HtmlURL           string      `json:"html_url"`
// 		ID                float64     `json:"id"`
// 		Location          string      `json:"location"`
// 		Login             string      `json:"login"`
// 		Name              string      `json:"name"`
// 		OrganizationsURL  string      `json:"organizations_url"`
// 		PublicGists       float64     `json:"public_gists"`
// 		PublicRepos       float64     `json:"public_repos"`
// 		ReceivedEventsURL string      `json:"received_events_url"`
// 		ReposURL          string      `json:"repos_url"`
// 		StarredURL        string      `json:"starred_url"`
// 		SubscriptionsURL  string      `json:"subscriptions_url"`
// 		Type              string      `json:"type"`
// 		UpdatedAt         string      `json:"updated_at"`
// 		URL               string      `json:"url"`
// 	}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ChimeraCoder/gojson"
)

var (
	name = flag.String("name", "Foo", "the name of the struct")
	pkg  = flag.String("pkg", "main", "the name of the package for the generated code")
)

func main() {
	flag.Parse()

	if isInteractive() {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Expects input on stdin")
		os.Exit(1)
	}

	if output, err := json2struct.Generate(os.Stdin, *name, *pkg); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	} else {
		fmt.Print(string(output))
	}
}

// Return true if os.Stdin appears to be interactive
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}
