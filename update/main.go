package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	blogURL  = "https://charly3pins.dev"
	filename = "../README.md"
)

type Readme struct {
	BlogURL string
	Posts   []Post
}

type Post struct {
	Title string
	Link  string
	Date  string
}

func main() {
	tpl := `
# Hey there! <img src="https://media.giphy.com/media/hvRJCLFzcasrR4ia7z/giphy.gif" width="25px">

I'm a Software Engineer at <a href="https://github.com/arexio">@arexio</a> <img src="https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif" width="30">

[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/intent/follow?screen_name=charly3pins)
[![RSS](https://img.shields.io/badge/RSS-FFA500?style=for-the-badge&logo=rss&logoColor=white)]({{.BlogURL}})
[![Linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/carlesfuste/)

## 👨‍💻 Blog

Besides writing code, I like to write articles about things that I find interesting. You can read the articles at **[charly3pins.dev]({{.BlogURL}})**

Latest posts:
{{range .Posts}}- **[{{.Title}}]({{.Link}})** ({{.Date}})
{{end}}

![](https://media.giphy.com/media/OPYnG3Xf8zLag/giphy.gif)

<sub>Last update on ` + time.Now().Format("02/01/2006") + `</sub>`

	p := gofeed.NewParser()
	feed, err := p.ParseURL(blogURL + "/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	var posts []Post
	for i := 0; i < 5; i++ {
		p := feed.Items[i]
		post := Post{
			Title: p.Title,
			Link:  p.Link,
			Date:  relativeDate(p.Published),
		}
		posts = append(posts, post)
	}

	readme := Readme{
		BlogURL: blogURL,
		Posts:   posts,
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	t := template.Must(template.New("readme").Parse(tpl))
	if err = t.Execute(file, readme); err != nil {
		log.Fatalf("error executing template: %v", err)
	}
}

func relativeDate(d string) string {
	dt, err := time.Parse("Mon, 02 Jan 2006", d)
	if err != nil {
		log.Fatalf("error parsing post date: %v", err)
	}
	now := time.Now().Unix()
	days := (now - dt.Unix()) / 86400
	months := (now - dt.Unix()) / 2592000

	if days == 0 { // Published today
		return d
	}

	date := ""
	if days < 31 { // Published in the last 31 days
		date = strconv.Itoa(int(days))
		if days == 1 {
			date += " day"
		} else {
			date += " days"
		}
	} else {
		date = strconv.Itoa(int(months))
		if months == 1 { // Published month(s) ago
			date += " month"
		} else {
			date += " months"
		}
	}
	return fmt.Sprintf("%s ago", date)
}
