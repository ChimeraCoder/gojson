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
// 	type Repository struct {
//     	ArchiveURL       string      `json:"archive_url"`
//     	AssigneesURL     string      `json:"assignees_url"`
//     	BlobsURL         string      `json:"blobs_url"`
//     	BranchesURL      string      `json:"branches_url"`
//     	CloneURL         string      `json:"clone_url"`
//     	CollaboratorsURL string      `json:"collaborators_url"`
//     	CommentsURL      string      `json:"comments_url"`
//     	CommitsURL       string      `json:"commits_url"`
//     	CompareURL       string      `json:"compare_url"`
//     	ContentsURL      string      `json:"contents_url"`
//     	ContributorsURL  string      `json:"contributors_url"`
//     	CreatedAt        string      `json:"created_at"`
//     	DefaultBranch    string      `json:"default_branch"`
//     	Description      string      `json:"description"`
//     	DownloadsURL     string      `json:"downloads_url"`
//     	EventsURL        string      `json:"events_url"`
//     	Fork             bool        `json:"fork"`
//     	Forks            float64     `json:"forks"`
//     	ForksCount       float64     `json:"forks_count"`
//     	ForksURL         string      `json:"forks_url"`
//     	FullName         string      `json:"full_name"`
//     	GitCommitsURL    string      `json:"git_commits_url"`
//     	GitRefsURL       string      `json:"git_refs_url"`
//     	GitTagsURL       string      `json:"git_tags_url"`
//     	GitURL           string      `json:"git_url"`
//     	HasDownloads     bool        `json:"has_downloads"`
//     	HasIssues        bool        `json:"has_issues"`
//     	HasWiki          bool        `json:"has_wiki"`
//     	Homepage         interface{} `json:"homepage"`
//     	HooksURL         string      `json:"hooks_url"`
//     	HtmlURL          string      `json:"html_url"`
//     	ID               float64     `json:"id"`
//     	IssueCommentURL  string      `json:"issue_comment_url"`
//     	IssueEventsURL   string      `json:"issue_events_url"`
//     	IssuesURL        string      `json:"issues_url"`
//     	KeysURL          string      `json:"keys_url"`
//     	LabelsURL        string      `json:"labels_url"`
//     	Language         string      `json:"language"`
//     	LanguagesURL     string      `json:"languages_url"`
//     	MasterBranch     string      `json:"master_branch"`
//     	MergesURL        string      `json:"merges_url"`
//     	MilestonesURL    string      `json:"milestones_url"`
//     	MirrorURL        interface{} `json:"mirror_url"`
//     	Name             string      `json:"name"`
//     	NetworkCount     float64     `json:"network_count"`
//     	NotificationsURL string      `json:"notifications_url"`
//     	OpenIssues       float64     `json:"open_issues"`
//     	OpenIssuesCount  float64     `json:"open_issues_count"`
//     	Owner            struct {
//         	AvatarURL         string  `json:"avatar_url"`
//         	EventsURL         string  `json:"events_url"`
//         	FollowersURL      string  `json:"followers_url"`
//         	FollowingURL      string  `json:"following_url"`
//         	GistsURL          string  `json:"gists_url"`
//         	GravatarID        string  `json:"gravatar_id"`
//         	HtmlURL           string  `json:"html_url"`
//         	ID                float64 `json:"id"`
//         	Login             string  `json:"login"`
//         	OrganizationsURL  string  `json:"organizations_url"`
//         	ReceivedEventsURL string  `json:"received_events_url"`
//         	ReposURL          string  `json:"repos_url"`
//         	SiteAdmin         bool    `json:"site_admin"`
//         	StarredURL        string  `json:"starred_url"`
//         	SubscriptionsURL  string  `json:"subscriptions_url"`
//         	Type              string  `json:"type"`
//         	URL               string  `json:"url"`
//     } `	json:"owner"`
//     	Private         bool    `json:"private"`
//     	PullsURL        string  `json:"pulls_url"`
//     	PushedAt        string  `json:"pushed_at"`
//     	Size            float64 `json:"size"`
//     	SshURL          string  `json:"ssh_url"`
//     	StargazersURL   string  `json:"stargazers_url"`
//     	StatusesURL     string  `json:"statuses_url"`
//     	SubscribersURL  string  `json:"subscribers_url"`
//     	SubscriptionURL string  `json:"subscription_url"`
//     	SvnURL          string  `json:"svn_url"`
//     	TagsURL         string  `json:"tags_url"`
//     	TeamsURL        string  `json:"teams_url"`
//     	TreesURL        string  `json:"trees_url"`
//     	UpdatedAt       string  `json:"updated_at"`
//     	URL             string  `json:"url"`
//     	Watchers        float64 `json:"watchers"`
//     	WatchersCount   float64 `json:"watchers_count"`
// 	}
package gojson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

var ForceFloats bool

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"NTP":   true,
	"DB":    true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// Options includes the flexible settings required
// for the representation of an object.
type Options struct {
	PkgName   string
	Tags      []string
	SubStruct bool
	Omitempty bool
}

type Parser func(io.Reader) (interface{}, error)

func ParseJson(input io.Reader) (interface{}, error) {
	var result interface{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func ParseYaml(input io.Reader) (interface{}, error) {
	var result interface{}
	b, err := readFile(input)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func readFile(input io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, input)
	if err != nil {
		return []byte{}, nil
	}
	return buf.Bytes(), nil
}

// Given a JSON string representation of an object and a name structName,
// attemp to generate a struct definition
func Generate(input io.Reader, parser Parser, structName, pkgName string, tags []string, subStruct bool) ([]byte, error) {
	return GenerateWithOptions(input, parser, structName, &Options{
		PkgName:   pkgName,
		SubStruct: subStruct,
		Tags:      tags,
		Omitempty: false,
	})
}

// Given a JSON string representation of an object and a name structName,
// attemp to generate a struct definition. Using Options for flexible setup.
func GenerateWithOptions(input io.Reader, parser Parser, structName string, opts *Options) ([]byte, error) {
	var subStructMap map[string]string = nil
	if opts.SubStruct {
		subStructMap = make(map[string]string)
	}

	var result map[string]interface{}

	iresult, err := parser(input)
	if err != nil {
		return nil, err
	}

	pkgLine := ""
	if opts.PkgName != "" {
		pkgLine = fmt.Sprintf("package %s\n\n", opts.PkgName)
	}

	switch iresult := iresult.(type) {
	case map[interface{}]interface{}:
		result = convertKeysToStrings(iresult)
	case map[string]interface{}:
		result = iresult
	case []interface{}:
		src := fmt.Sprintf("%stype %s %s\n",
			pkgLine,
			structName,
			typeForValue(iresult, structName, opts.Tags, subStructMap, opts))
		formatted, err := format.Source([]byte(src))
		if err != nil {
			err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
		}
		return formatted, err
	default:
		return nil, fmt.Errorf("unexpected type: %T", iresult)
	}

	src := fmt.Sprintf("%stype %s %s}",
		pkgLine,
		structName,
		generateTypes(result, structName, opts.Tags, 0, subStructMap, opts))

	keys := make([]string, 0, len(subStructMap))
	for key := range subStructMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, k := range keys {
		src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap[k], k)
	}

	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}
	return formatted, err
}

func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range obj {
		res[fmt.Sprintf("%v", k)] = v
	}

	return res
}

// Generate go struct entries for a map[string]interface{} structure
func generateTypes(obj map[string]interface{}, structName string, tags []string, depth int, subStructMap map[string]string, opts *Options) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := obj[key]
		valueType := typeForValue(value, structName, tags, subStructMap, opts)

		//value = mergeElements(value)

		//If a nested value, recurse
		switch value := value.(type) {
		case []interface{}:
			if len(value) > 0 {
				sub := ""
				if v, ok := value[0].(map[interface{}]interface{}); ok {
					sub = generateTypes(convertKeysToStrings(v), structName, tags, depth+1, subStructMap, opts) + "}"
				} else if v, ok := value[0].(map[string]interface{}); ok {
					sub = generateTypes(v, structName, tags, depth+1, subStructMap, opts) + "}"
				}

				if sub != "" {
					subName := sub

					if subStructMap != nil {
						if val, ok := subStructMap[sub]; ok {
							subName = val
						} else {
							subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

							subStructMap[sub] = subName
						}
					}

					valueType = "[]" + subName
				}
			}
		case map[interface{}]interface{}:
			sub := generateTypes(convertKeysToStrings(value), structName, tags, depth+1, subStructMap, opts) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

					subStructMap[sub] = subName
				}
			}
			valueType = subName
		case map[string]interface{}:
			sub := generateTypes(value, structName, tags, depth+1, subStructMap, opts) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					subName = fmt.Sprintf("%v_sub%v", structName, len(subStructMap)+1)

					subStructMap[sub] = subName
				}
			}

			valueType = subName
		}

		fieldName := FmtFieldName(key)

		tagList := make([]string, 0)
		omitempty := ""
		for _, t := range tags {
			if opts.Omitempty {
				switch t {
				case "json", "xml":
					omitempty = ",omitempty"
				}
			}
			tagList = append(tagList, fmt.Sprintf("%s:\"%s%s\"", t, key, omitempty))
		}

		tagLine := strings.Join(tagList, " ")
		if tagLine != "" {
			tagLine = fmt.Sprintf("`%s`", tagLine)
		}

		structure += fmt.Sprintf("\n%s %s %s",
			fieldName,
			valueType,
			tagLine,
		)
	}
	return structure
}

// FmtFieldName formats a string as a struct key
//
// Example:
// 	FmtFieldName("foo_id")
// Output: FooID
func FmtFieldName(s string) string {
	runes := []rune(s)
	for len(runes) > 0 && !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) {
		runes = runes[1:]
	}
	if len(runes) == 0 {
		return "_"
	}

	s = stringifyFirstChar(string(runes))
	name := lintFieldName(s)
	runes = []rune(name)
	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}
	s = string(runes)
	s = strings.Trim(s, "_")
	if len(s) == 0 {
		return "_"
	}
	return s
}

func lintFieldName(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}

	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		runes := []rune(name)
		if u := strings.ToUpper(name); commonInitialisms[u] {
			copy(runes[0:], []rune(u))
		} else {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}

	allUpperWithUnderscore := true
	for _, r := range name {
		if !unicode.IsUpper(r) && r != '_' {
			allUpperWithUnderscore = false
			break
		}
	}
	if allUpperWithUnderscore {
		name = strings.ToLower(name)
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word

		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower->non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))

		} else if strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// generate an appropriate struct type entry
func typeForValue(value interface{}, structName string, tags []string, subStructMap map[string]string, opts *Options) string {
	//Check if this is an array
	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + typeForValue(mergeElements(objects).([]interface{})[0], structName, tags, subStructMap, opts)
		}
		return "[]interface{}"
	} else if object, ok := value.(map[interface{}]interface{}); ok {
		return generateTypes(convertKeysToStrings(object), structName, tags, 0, subStructMap, opts) + "}"
	} else if object, ok := value.(map[string]interface{}); ok {
		return generateTypes(object, structName, tags, 0, subStructMap, opts) + "}"
	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}
	v := reflect.TypeOf(value).Name()
	if v == "float64" {
		v = disambiguateFloatInt(value)
	}
	return v
}

// All numbers will initially be read as float64
// If the number appears to be an integer value, use int instead
func disambiguateFloatInt(value interface{}) string {
	const epsilon = .0001
	vfloat := value.(float64)
	if !ForceFloats && math.Abs(vfloat-math.Floor(vfloat+epsilon)) < epsilon {
		var tmp int64
		return reflect.TypeOf(tmp).Name()
	}
	return reflect.TypeOf(value).Name()
}

// convert first character ints to strings
func stringifyFirstChar(str string) string {
	first := str[:1]

	i, err := strconv.ParseInt(first, 10, 8)

	if err != nil {
		return str
	}

	return intToWordMap[i] + "_" + str[1:]
}

func mergeElements(i interface{}) interface{} {
	switch i := i.(type) {
	default:
		return i
	case []interface{}:
		l := len(i)
		if l == 0 {
			return i
		}
		for j := 1; j < l; j++ {
			i[0] = mergeObjects(i[0], i[j])
		}
		return i[0:1]
	}
}

func mergeObjects(o1, o2 interface{}) interface{} {
	if o1 == nil {
		return o2
	}

	if o2 == nil {
		return o1
	}

	if reflect.TypeOf(o1) != reflect.TypeOf(o2) {
		return nil
	}

	switch i := o1.(type) {
	default:
		return o1
	case []interface{}:
		if i2, ok := o2.([]interface{}); ok {
			i3 := append(i, i2...)
			return mergeElements(i3)
		}
		return mergeElements(i)
	case map[string]interface{}:
		if i2, ok := o2.(map[string]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	case map[interface{}]interface{}:
		if i2, ok := o2.(map[interface{}]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i
	}
}
