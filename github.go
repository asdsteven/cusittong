package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

const (
	owner       = "asdsteven"
	repo        = "cuaddropStorage"
	githubURL   = "https://raw.githubusercontent.com/asdsteven/cuaddropStorage/master/"
	githubToken = "df49515198af3bf33f558e7c160d121a7a077d81"
)

var githubClient *github.Client

var contentSHA string

var githubPush chan struct{}

func githubUpdate(d map[string][6]interface{}) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	msg := "update"
	fileOpt := github.RepositoryContentFileOptions{
		Message: &msg,
		Content: b,
		SHA:     &contentSHA,
	}
	res, _, err := githubClient.Repositories.UpdateFile(owner, repo, thisTerm[1], &fileOpt)
	if err != nil {
		return err
	}
	contentSHA = *res.Content.SHA
	return fmt.Errorf("finished")
}

func githubUpdater() {
	for range githubPush {
		d := make(map[string][6]interface{})
		for i, v := range db {
			reserves := [][3]interface{}{}
			changes := [][3]interface{}{}
			for _, w := range v.reserves {
				reserves = append(reserves, [3]interface{}{
					w.major,
					w.quota,
					w.enroll,
				})
			}
			for _, w := range v.changes {
				changes = append(changes, [3]interface{}{
					w.stamp,
					w.index,
					w.value,
				})
			}
			d[i] = [6]interface{}{
				v.group,
				v.title,
				v.quota,
				v.enroll,
				reserves,
				changes,
			}
		}
		if err := githubUpdate(d); err != nil {
			writeLogs <- fmt.Sprintf("githubUpdater: %v", err)
		}
	}
}

func githubFetchDB() error {
	msg := "create"
	fileOpt := github.RepositoryContentFileOptions{
		Message: &msg,
		Content: []byte("{}"),
	}
	_, _, err := githubClient.Repositories.CreateFile(owner, repo, thisTerm[1], &fileOpt)
	if err == nil {
		writeLogs <- fmt.Sprintf("github created %v", thisTerm[1])
	}
	tree, _, err := githubClient.Git.GetTree(owner, repo, "master", false)
	if err != nil {
		return err
	}
	for _, v := range tree.Entries {
		if *v.Path == thisTerm[1] {
			contentSHA = *v.SHA
		}
	}
	res, err := http.Get(githubURL + thisTerm[1])
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var d map[string][6]interface{}
	if err := json.Unmarshal(content, &d); err != nil {
		return err
	}
	db = make(map[string]*historyS)
	for i, v := range d {
		var reserves []reserveS
		var changes []changeS
		for _, w := range v[4].([]interface{}) {
			ws := w.([]interface{})
			reserves = append(reserves, reserveS{
				ws[0].(string),
				int(ws[1].(float64)),
				int(ws[2].(float64)),
			})
		}
		for _, w := range v[5].([]interface{}) {
			ws := w.([]interface{})
			switch wss := ws[2].(type) {
			case float64:
				changes = append(changes, changeS{
					int64(ws[0].(float64)),
					int(ws[1].(float64)),
					int(wss),
				})
			case string:
				changes = append(changes, changeS{
					int64(ws[0].(float64)),
					int(ws[1].(float64)),
					wss,
				})
			}
		}
		db[i] = &historyS{
			v[0].(string),
			v[1].(string),
			int(v[2].(float64)),
			int(v[3].(float64)),
			reserves,
			changes,
		}
	}
	return fmt.Errorf("finished")
}

func githubInit() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	githubClient = github.NewClient(tc)
	githubPush = make(chan struct{})
}
