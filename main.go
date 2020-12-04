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
func proccessString(inputs []string, s string) {
  output := ""
  r := len(strings.Split(s, "::"))
  v := strings.Split(s, "::")
  for i := 0;i < r;i++ {
    output += v[0]
    switch v[i] {
    case "URL":
      output += inputs[1]
    case "URI":
      output += inputs[0]
    default:
      //none
    }
    output += v[2]
    }
  }


func index(c *gin.Context) {
  links := `<div class="container">
      <a href="+"::URL::"+">
      <div class="picture">
        <img class="picture" src="::URI::"></img>
      </a><div class="text">
        This website, it is written in Golang, serves http currently and this is
        a thumbnail embeddable link.
        </div>
    </div>
</div>`
  output := ""
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
