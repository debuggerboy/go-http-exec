/******************************************************

Execute local commands from a unsecure HTTP Interface

This is for special cases, when you run your target
application in a fully internal environment.

Not adviced to be used on systems with direct internet
access.

******************************************************/

package main

import (
    "fmt"
    "os/exec"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the desert!",
        })
    })

    r.GET("/launcher/:cmd", func(c *gin.Context) {
        launchcmd := c.Params.ByName("cmd")
        err := exec.Command(launchcmd).Start()
        if err != nil {
            fmt.Printf("%s", err)
            c.Error(err)
            c.Abort()
        }
        fmt.Sprintf("%s Successfully Executed", launchcmd)
        output := fmt.Sprintf("%s Launched", launchcmd)
        c.String(http.StatusOK, output)
    })

    r.Run()

}
