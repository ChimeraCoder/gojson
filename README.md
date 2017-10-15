[![Build Status](https://travis-ci.org/ChimeraCoder/gojson.svg?branch=master)](https://travis-ci.org/ChimeraCoder/gojson)
gojson
======

gojson generates go struct definitions from json or yaml documents.

Example
----------

```sh
$ curl -s https://api.github.com/repos/chimeracoder/gojson | gojson -name=Repository

package main

type Repository struct {
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CloneURL         string      `json:"clone_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	CreatedAt        string      `json:"created_at"`
	DefaultBranch    string      `json:"default_branch"`
	Description      string      `json:"description"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	Fork             bool        `json:"fork"`
	Forks            float64     `json:"forks"`
	ForksCount       float64     `json:"forks_count"`
	ForksURL         string      `json:"forks_url"`
	FullName         string      `json:"full_name"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	HasDownloads     bool        `json:"has_downloads"`
	HasIssues        bool        `json:"has_issues"`
	HasWiki          bool        `json:"has_wiki"`
	Homepage         interface{} `json:"homepage"`
	HooksURL         string      `json:"hooks_url"`
	HtmlURL          string      `json:"html_url"`
	ID               float64     `json:"id"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	Language         string      `json:"language"`
	LanguagesURL     string      `json:"languages_url"`
	MasterBranch     string      `json:"master_branch"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	MirrorURL        interface{} `json:"mirror_url"`
	Name             string      `json:"name"`
	NetworkCount     float64     `json:"network_count"`
	NotificationsURL string      `json:"notifications_url"`
	OpenIssues       float64     `json:"open_issues"`
	OpenIssuesCount  float64     `json:"open_issues_count"`
	Owner            struct {
		AvatarURL         string  `json:"avatar_url"`
		EventsURL         string  `json:"events_url"`
		FollowersURL      string  `json:"followers_url"`
		FollowingURL      string  `json:"following_url"`
		GistsURL          string  `json:"gists_url"`
		GravatarID        string  `json:"gravatar_id"`
		HtmlURL           string  `json:"html_url"`
		ID                float64 `json:"id"`
		Login             string  `json:"login"`
		OrganizationsURL  string  `json:"organizations_url"`
		ReceivedEventsURL string  `json:"received_events_url"`
		ReposURL          string  `json:"repos_url"`
		SiteAdmin         bool    `json:"site_admin"`
		StarredURL        string  `json:"starred_url"`
		SubscriptionsURL  string  `json:"subscriptions_url"`
		Type              string  `json:"type"`
		URL               string  `json:"url"`
	} `json:"owner"`
	Private         bool    `json:"private"`
	PullsURL        string  `json:"pulls_url"`
	PushedAt        string  `json:"pushed_at"`
	Size            float64 `json:"size"`
	SshURL          string  `json:"ssh_url"`
	StargazersURL   string  `json:"stargazers_url"`
	StatusesURL     string  `json:"statuses_url"`
	SubscribersURL  string  `json:"subscribers_url"`
	SubscriptionURL string  `json:"subscription_url"`
	SvnURL          string  `json:"svn_url"`
	TagsURL         string  `json:"tags_url"`
	TeamsURL        string  `json:"teams_url"`
	TreesURL        string  `json:"trees_url"`
	UpdatedAt       string  `json:"updated_at"`
	URL             string  `json:"url"`
	Watchers        float64 `json:"watchers"`
	WatchersCount   float64 `json:"watchers_count"`
}
```

CLI Installation
----------------

```sh
$ go get github.com/ChimeraCoder/gojson/gojson
```

Assuming `$GOPATH/bin` is in your `PATH`, you can now invoke `gojson` directly.


API Installation
----------------

```sh
$ go get github.com/ChimeraCoder/gojson/gojson
```

Development
-----------

```
$ git clone https://github.com/ChimeraCoder/gojson.git
$ cd gojson
$ go test
```

**Building CLI**

```
$ go build -o _build/gojson ./gojson
```

**Installing CLI**

```
$ go install ./gojson
```

**Formatting**

```
$ gofmt -w -e -s -l .
```

Related Work
------------

github.com/str1ngs/jflect

License
----------

gojson is free software distributed under Version 3 of the GNU Public License.

As of the time of writing, this is the same license used for gcc (and therefore gccgo), so it is unlikely to restrict use in any way. Note that the GPL does not extend to any output generated by gojson; the GPL only applies to software which includes copies of gojson itself.
