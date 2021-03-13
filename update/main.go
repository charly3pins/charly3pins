package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	blogURL  = "https://charly3pins.dev"
	filename = "../README.md"
)

func main() {
	p := gofeed.NewParser()
	feed, err := p.ParseURL(blogURL + "/index.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	newestItem := feed.Items[0]

	date := time.Now().Format("02/01/2006")

	hi := "# Hey there! <img src=\"https://media.giphy.com/media/hvRJCLFzcasrR4ia7z/giphy.gif\" width=\"25px\">\n\nI'm a Software Engineer at <a href=\"https://github.com/arexio\">@arexio</a><img src=\"https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif\" width=\"30\">\n\n[![Twitter](https://img.shields.io/twitter/follow/charly3pins?label=%40charly3pins&style=social)](https://twitter.com/intent/follow?screen_name=charly3pins)\n![GitHub](https://img.shields.io/github/followers/charly3pins?label=%40charly3pins&style=social)\n[![Linkedin](https://img.shields.io/badge/Linkedin-Carles%20Fuste-blue?style=social&logo=Linkedin)](https://www.linkedin.com/in/carlesfuste/)"
	blog := "## &#x270d; Blog & Writing\n\nApart from coding you can find my articles on my website at [charly3pins.dev](https://charly3pins.dev/).\n\nMy latest blog post is: **[" + newestItem.Title + "](" + newestItem.Link + ")**. If you liked, you can subscribe to my [**blog RSS**](" + blogURL + "/index.xml) feed."
	tech := "## 🔧 Tools & Technologies\n\nThe main tools and technologies that I use everyday are:\n\n ![](https://img.shields.io/badge/Golang-informational?style=flat&logo=go&logoColor=white&color=29BEB0) ![](https://img.shields.io/badge/Docker-informational?style=flat&logo=docker&logoColor=white&color=049CEC) ![](https://img.shields.io/badge/Kubernetes-informational?style=flat&logo=kubernetes&logoColor=white&color=047ADC)\n\n ![](https://img.shields.io/badge/Git-informational?style=flat&logo=git&logoColor=white&color=F1502F) ![](https://img.shields.io/badge/PostgreSQL-informational?style=flat&logo=postgresql&logoColor=white&color=blue) ![](https://img.shields.io/badge/Jenkins-informational?style=flat&logo=jenkins&logoColor=white&color=D33834)\n\n ![](https://img.shields.io/badge/Linux-informational?style=flat&logo=linux&logoColor=white&color=orange) ![](https://img.shields.io/badge/ZSH-informational?style=flat&logo=gnu-bash&logoColor=white&color=brightgreen)\n\n ![](https://media.giphy.com/media/OPYnG3Xf8zLag/giphy.gif)"
	updated := "<sub>Last update on " + date + ".</sub>"

	data := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n", hi, blog, tech, updated)

	file, err := os.Create(filename)
	if err != nil {
		log.Println("error creating file", err)
		return
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		log.Println("error writing content to file", err)
		return
	}
	file.Sync()
}
