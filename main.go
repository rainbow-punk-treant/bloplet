package main

import (
  "fmt"
  "os"
  "strings"
  "github.com/microcosm-cc/bluemonday"
  "github.com/gin-gonic/gin"
  "io/ioutil"
)

const PATH = "/home/weasel/go/src/gitlab.com/entro-pi/supercut/super/"

func main() {
  fmt.Println("Doot")

  server := gin.Default()
  server.GET("/", index)
  server.GET("/home", index)
  server.NoRoute(index)
  server.Static("/assets", "/home/weasel/go/src/gitlab.com/entro-pi/supercut/super/")

  server.Run(":80")
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


func index(c *gin.Context) {
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
  output := ""
  payloadAlpha := make(map[string]string, 35)
  payloadBeta := make(map[string]string, 35)

  payloadAlpha["URL"] = "https://gitlab.com/entro-pi/supercut"
  payloadAlpha["URI"] = "assets/img/cc.png"
  payloadBeta["URL"] = "https://gitlab.com/entro-pi/snowfone"
  payloadBeta["URI"] = "assets/img/knoife.png"

  stringPayload := processString(payloadAlpha, links)
  stringPayload += processString(payloadBeta, links)


  dir, err := ioutil.ReadDir(PATH)
  if err != nil {
    panic(err)
  }
  for _, val := range dir {
    if val.Name() == "img"{
      file, err := os.Open(PATH+"/index_feader.html")
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

      pictureDir, err := ioutil.ReadDir(PATH+"/img")
      if err != nil {
        panic(err)
      }

      for i, pic := range pictureDir {
        if len(pic.Name()) > 2 {
          //do thing!
        }else {
          continue
        }
        item := `<a href="assets/img/`+pic.Name()+`" data-caption="testing lightbox">

            <img src="assets/img/`+pic.Name()+`" width="39px" height="39px" style="border-radius: 50%;" alt="First image">

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
      file, err = os.Open(PATH+"/index_hooter.html")
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
