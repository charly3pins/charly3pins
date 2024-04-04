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

	hoursDay  = 24
	daysWeek  = 7
	daysMonth = 31
	daysYear  = 365
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
## Hello 👋

Engineering at [Bankable.com](https://bnkbl.com/), a global architect of innovative payment solutions enabling Banking as a Service & [AREX.io](https://arex.io/), building a real-time exchange for short-term corporate credit.

## Blog ✍️

Besides writing code, I like to write articles about things that I find interesting. You can read the articles at **[charly3pins.dev]({{.BlogURL}})**

Latest posts:
{{range .Posts}}- **[{{.Title}}]({{.Link}})**
{{end}}

## My readings 📚

You can find  the list of books I am reading, have read or plan to read in [this repo](https://github.com/charly3pins/readings) or in [this section](https://charly3pins.dev/readings).

Currently I am reading:

* [Software Engineering at Google: Lessons Learned from Programming Over Time](https://amzn.to/3TgWORq) by _Titus Winters, Tom Manshreck, Hyrum Wright_
* [The Phoenix Project: A Novel about IT, DevOps, and Helping Your Business Win](https://amzn.to/3TD3jPG) by _Gene Kim, Kevin Behr, George Spafford_
* [Las 48 leyes del poder](https://amzn.to/3IEvibx) by _Robert Greene, Joost Elffers_
* [Un paso por delante de Wall Street: Cómo utilizar lo que ya sabes para ganar dinero en bolsa](https://amzn.to/3VLPIHj) by _Peter Lynch_

![](https://media.giphy.com/media/OPYnG3Xf8zLag/giphy.gif)

## Keep in touch 👨‍💻

[![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/intent/follow?screen_name=charly3pins)
[![RSS](https://img.shields.io/badge/RSS-FFA500?style=for-the-badge&logo=rss&logoColor=white)](https://charly3pins.dev)
[![DEV.to](https://img.shields.io/badge/dev.to-0A0A0A?style=for-the-badge&logo=dev.to&logoColor=white)](https://dev.to/charly3pins)
[![Linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/carlesfuste/)`

	p := gofeed.NewParser()
	feed, err := p.ParseURL(blogURL + "/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}

	var posts []Post
	for i := 0; i < 10; i++ {
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
	published, err := time.Parse("Mon, 02 Jan 2006", d)
	if err != nil {
		log.Fatalf("error parsing post date: %v", err)
	}
	now := time.Now()
	difference := now.UTC().Sub(published.UTC())

	if difference.Hours()/hoursDay == 0 { // Published today
		return d
	}

	days := int64(difference.Hours() / hoursDay)
	weeks := int64(difference.Hours() / hoursDay / daysWeek)
	months := int64(difference.Hours() / hoursDay / daysMonth)
	years := int64(difference.Hours() / hoursDay / daysYear)

	date := ""
	if days < 31 { // Published in the last 31 days
		date = strconv.Itoa(int(days))
		if days == 0 {
			return "Today"
		} else if days == 1 {
			date += " day"
		} else {
			date += " days"
		}
	} else if weeks < 4 { // Published in the last 4 weeks
		date = strconv.Itoa(int(weeks))
		if weeks == 1 {
			date += " week"
		} else {
			date += " weeks"
		}
	} else if months < 12 { // Published in the last 12 months
		date = strconv.Itoa(int(months))
		if months == 1 {
			date += " month"
		} else {
			date += " months"
		}
	} else { // Published in the last year(s)
		date = strconv.Itoa(int(years))
		if years == 1 {
			date += " year"
		} else {
			date += " years"
		}
	}
	return fmt.Sprintf("%s ago", date)
}
