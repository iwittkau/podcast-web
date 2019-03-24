package web

import (
	"html/template"
)

// Page is a single page
type Page struct {
	Title   string
	Index   int
	Menu    []MenuItem
	Content template.HTML `yaml:"-"`
}

// MenuItem is an item in the main menu
type MenuItem struct {
	Name string
	Link string
}

// EpisodeLink is a link to an episode
type EpisodeLink struct {
	Name   string
	Link   string
	Number int
}

// Site is the main struct that holds all the site's data
type Site struct {
	Title    string
	Author   string
	Menu     []MenuItem
	Content  template.HTML `yaml:"-"`
	Episodes []EpisodeLink
}
