package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	fmt.Println("bloplet v1.0 Author rainbow_punk_treant")
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("Usage is \033[38;2;0;200;0mbloplet \033[38;2;150;0;150musername \033[38;2;50;100;100mURL\033[0m")
		os.Exit(1)
	}

	USER := os.Args[1]
	URL := os.Args[2]

	server := gin.Default()
	server.GET("/", index)
	server.GET("/home", index)
	server.NoRoute(index)
	server.Static("/assets", "/home/"+USER+"/go/src/github.com/rainbow-punk-treant/bloplet/super/")
	autotls.Run(server, URL)
}

func getIt(c *gin.Context) {
	policy := bluemonday.UGCPolicy()
	policy.Sanitize("no")
}
func processString(inputs map[string]string, s string) string {
	output := ""
	//r := len(strings.Split(s, "::"))
	v := strings.Split(s, "::")
	output += v[0]

	output += inputs["URL"]
	output += v[2]
	output += inputs["URI"]
	output += v[4]

	return output
}

func populate(uri string, url string) string {
	links := `<div class="container">
      <a href="::URL::">
      <div class="picture">
        <img class="picture" src="::URI::"></img>
      </a><div class="text">
        This website, it is written in Golang, serves http currently and this is
        a thumbnail embeddable link. And now it is a template action to be applied.
        </div>
    </div>
</div>`

	payloadAlpha := make(map[string]string, 35)

	payloadAlpha["URL"] = url
	payloadAlpha["URI"] = uri

	stringPayload := processString(payloadAlpha, links)

	return stringPayload

}

func index(c *gin.Context) {
	output := ""
	PATH := "/home/" + os.Args[1] + "/go/src/github.com/rainbow-punk-treant/bloplet/super/"

	//Define your payloads
	pl := "https://gitlab.com/entro-pi/supercut"
	pi := "assets/img/cc.png"
	pbl := "https://gitlab.com/entro-pi/snowfone"
	pbi := "assets/img/knoife.png"

	//Format  your payloads
	stringPayload := populate(pi, pl)
	stringPayload += populate(pbi, pbl)

	dir, err := ioutil.ReadDir(PATH)
	if err != nil {
		panic(err)
	}
	for _, val := range dir {
		if val.Name() == "img" {
			file, err := os.Open(PATH + "/index_feader.html")
			if err != nil {
				panic(err)
			}

			payload, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			c.Writer.Write(payload)
			c.Writer.Write([]byte(stringPayload))
			//c.File(PATH+"/index_feader.html")

			pictureDir, err := ioutil.ReadDir(PATH + "/img")
			if err != nil {
				panic(err)
			}

			for i, pic := range pictureDir {
				if len(pic.Name()) > 2 {
					//do thing!
				} else {
					continue
				}
				item := `<a href="assets/img/` + pic.Name() + `" data-caption="testing lightbox">

            <img src="assets/img/` + pic.Name() + `" width="39px" height="39px" style="border-radius: 50%;" alt="First image">

            </a>`
				item0 := `<div class="gallery" style="margin: 0 auto;">`
				end := false
				if i == 0 {
					output += item0
					end = true
				}
				output += item
				if end {
					output += `</div>`
				}
			}

			//c.Writer.Write([]byte(output))
			file, err = os.Open(PATH + "/index_hooter.html")
			if err != nil {
				panic(err)
			}

			payload, err = ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			c.Writer.Write(payload)
		}
	}
}
