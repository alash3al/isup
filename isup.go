
package main

import ("log"; "flag"; "time"; "net/http"; "os/exec")

var (
        interval        =       flag.Int("intval", 1, "the check interval in seconds")
        url             =       flag.String("url", "http://google.com", "the url to check")
        media           =       flag.String("media", "none", "the media file to play when the link become available")
)

func main() {
        flag.Parse()
        if *url == "" {
                log.Fatal("You must enter the website url")
        }
        for {
                select {
                        case <- time.After(time.Duration(*interval) * time.Second):
                                resp, _ := http.Get(*url)
                                if resp != nil {
                                        resp.Body.Close()
                                }
                                if resp != nil && resp.StatusCode >= 400 {
                                        log.Println("Still Closed ...")
                                } else if resp != nil && resp.StatusCode >= 200 {
                                        log.Println("Available now ...")
                                        exec.Command("cvlc", *media).Run()
                                        break
                                }
                }
        }
}
