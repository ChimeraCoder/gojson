/*
GoJson generates go struct defintions from JSON documents

It reads from stdin and prints to stdout.

Usage:
	gojson [flags]

The flags are:
	-name
		the name of the struct
	-pkg
		the name of the package for the generated code
	-inputName
		the name of the input file containing JSON (if input not provided via STDIN)
	-outputName
		the name of the file to write the output to (outputs to STDOUT by default)

Examples

Convert a simple json string

	echo '{"hello":"world"}' | gojson

Output:

	type Foo struct {
	      Hello string `json:"hello"`
	}

Convert resulting json of an endpoint,
giving the resulting struct the name "Repository"

	curl -s https://api.github.com/repos/chimeracoder/gojson | gojson -name=Repository

Output:

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
*/
package main
