 package main

 import (
     "bufio"
     "flag"
     "io"
     "os"
     "strings"
     "time"

     "github.com/gin-gonic/gin"
 )

 func main() {
     r := gin.Default()
     gin.DisableConsoleColor()
     f, err := os.Create(time.Now().String() + ".log")
     if err != nil {
         panic(err)
     }
     defer f.Close()
     gin.DefaultWriter = io.MultiWriter(f)
     r.GET("/stz/:service/*host", func(c *gin.Context) {
         dirEntry, err := os.ReadDir(".")
         if err != nil {
             c.JSON(500, nil)
             return
         }
         service := c.Param("service")
         var filename string
         for _, v := range dirEntry {
             if service+".stz" == v.Name() {
                 filename = v.Name()
             }
         }
         if filename == "" {
             c.JSON(200, nil)
             return
         }
         file, err := os.Open(filename)
         result := make(map[string]map[string]string)
         if err != nil {
             c.JSON(500, nil)
             return
         }
         defer file.Close()
         scanner := bufio.NewScanner(file)
         var subject string
         for scanner.Scan() {
             text := scanner.Text()
             trimedtext := strings.TrimSpace(text)
             if strings.HasPrefix(trimedtext, "#") {
                 continue
             }
             if strings.HasSuffix(trimedtext, ":") {
                 subject = trimedtext[:len(trimedtext)-1]
                 result[subject] = make(map[string]string)
             } else if !strings.Contains(trimedtext, "=") {
                 continue
             } else {
                 split := strings.Split(trimedtext, "=")
                 result[subject][strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
             }
         }
         subject = strings.TrimPrefix(c.Param("subject"), "/")
         if subject == "" {
             c.JSON(200, result)
         } else {
             c.JSON(200, result[subject])
         }
     })
     port := flag.String("port", "9867", "the port of server")
     flag.Parse()
     r.Run(":" + *port)
 }
