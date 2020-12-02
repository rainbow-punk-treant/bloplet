package main

import (
  "fmt"
  "os"

  "github.com/gin-gonic/gin"
  "io/ioutil"
)

const PATH = "super/"

func main() {
  fmt.Println("Doot")

  server := gin.Default()
  server.GET("/home", index)
  server.NoRoute(index)
  server.Static("/img", "super/img")

  server.Run(":80")
}


func index(c *gin.Context) {
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
      pictureDir, err := ioutil.ReadDir(PATH+"img")
      if err != nil {
        panic(err)
      }

      for i, pic := range pictureDir {
        if len(pic.Name()) > 2 {
          //do thing!
        }else {
          continue
        }
        item := `<a href="img/`+pic.Name()+`" data-caption="testing lightbox">

            <img src="img/`+pic.Name()+`" width="250px" height="300px" alt="First image">

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

      c.Writer.Write([]byte(output))
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
