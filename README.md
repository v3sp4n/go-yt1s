
download video/music of service `yt1s`

<h1>installing</h1>

`go get github.com/v3sp4n/go-yt1s`

<h1>yt1s.</h1> 

```go
error := yt1s.download(url, quality, path string)
&map[string][]map[string]string,error := yt1s.GetAvalibleQualites(url string)
```
<h1>usage</h1>

```go
package main
import (
    "fmt"

    "github.com/v3sp4n/go-yt1s"
)
func main() {
    urls := []string{
        "https://www.youtube.com/watch?v=dQw4w9WgXcQ&pp=ygUJcmljayByb2xs",
        "https://www.youtube.com/watch?v=OedY-kNAxD4&pp=ygULZ29vZnkgc291bmQ%3D",
    }
    for i := range urls{
        q,err := yt1s.GetAvalibleQualites(urls[i])
        fmt.Println("avalible qualites for",urls[i],q,err)
        err2 := yt1s.Download(urls[0],"720p/*audio?put"mp3"*/","/home/vespan/Desktop/")
        fmt.Println(err2)
    }
}
```
