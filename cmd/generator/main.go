package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/golang-commonmark/markdown"

	web "github.com/iwittkau/podcast-web"
	"github.com/iwittkau/podcast-web/yaml"
)

func main() {
	site, err := yaml.ReadSiteConfig("example_content/site.yml")
	if err != nil {
		panic(err)
	}

	md := markdown.New(markdown.XHTMLOutput(true))
	indexMDData, err := ioutil.ReadFile("example_content/index.md")

	site.Content = template.HTML(md.RenderToString(indexMDData))

	t, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		panic(err)
	}

	index := t.Lookup("index.html")
	if index == nil {
		panic("index template not found")
	}

	data := []byte{}
	buf := bytes.NewBuffer(data)
	w := bufio.NewWriter(buf)

	err = index.ExecuteTemplate(w, "index.html", site)
	if err != nil {
		panic(err)
	}

	err = w.Flush()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("dist/index.html", buf.Bytes(), os.ModePerm)

	if err != nil {
		panic(err)
	}
	fmt.Println("dist/index.html", "wrote", len(buf.Bytes()), "bytes")

	for i := range site.Episodes {
		site.Episodes[i].Link = strings.ReplaceAll(strings.ToLower(fmt.Sprintf("%s-%d.html", site.Episodes[i].Name, site.Episodes[i].Number)), " ", "-")
	}

	for i := range site.Menu {

		data = []byte{}
		buf = bytes.NewBuffer(data)
		w = bufio.NewWriter(buf)

		mdFile := fmt.Sprintf("example_content/%s.md", strings.ToLower(site.Menu[i].Name))

		mdData, err := ioutil.ReadFile(mdFile)

		page := web.Site{
			Title:   site.Menu[i].Name,
			Menu:    site.Menu,
			Content: template.HTML(md.RenderToString(mdData)),
		}

		if page.Title == "Episodes" {
			page.Episodes = site.Episodes
			err = index.ExecuteTemplate(w, "episodes.html", page)
			if err != nil {
				panic(err)
			}
		} else {
			err = index.ExecuteTemplate(w, "site.html", page)
			if err != nil {
				panic(err)
			}
		}

		err = w.Flush()
		if err != nil {
			panic(err)
		}
		filename := fmt.Sprintf("dist/%s.html", strings.ToLower(site.Menu[i].Name))
		err = ioutil.WriteFile(filename, buf.Bytes(), os.ModePerm)

		if err != nil {
			panic(err)
		}
		fmt.Println(filename, "wrote", len(buf.Bytes()), "bytes")
	}
	for i := range site.Episodes {

		data = []byte{}
		buf = bytes.NewBuffer(data)
		w = bufio.NewWriter(buf)

		mdFile := fmt.Sprintf("example_content/episode-%d.md", site.Episodes[i].Number)

		mdData, err := ioutil.ReadFile(mdFile)
		if err != nil {
			panic(err)
		}

		page := web.Page{
			Title:   site.Episodes[i].Name,
			Menu:    site.Menu,
			Content: template.HTML(md.RenderToString(mdData)),
		}

		err = index.ExecuteTemplate(w, "site.html", page)
		if err != nil {
			panic(err)
		}

		err = w.Flush()
		if err != nil {
			panic(err)
		}
		filename := fmt.Sprintf("dist/%s-%d.html", strings.ReplaceAll(strings.ToLower(site.Episodes[i].Name), " ", "-"), site.Episodes[i].Number)
		err = ioutil.WriteFile(filename, buf.Bytes(), os.ModePerm)

		if err != nil {
			panic(err)
		}
		fmt.Println(filename, "wrote", len(buf.Bytes()), "bytes")
	}

}
