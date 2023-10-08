
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
        "https://youtu.be/OjNpRbNdR7E?si=YPqMT79475UfM4PM",
    }
    for i := 0; i < len(urls); i++ {
        q,err := yt1s.GetAvalibleQualites(urls[i])
        fmt.Println("avalible qualites for",urls[i],q,err)
        err2 := yt1s.Download(urls[0],"720p","/home/vespan/Desktop/")
        fmt.Println(err2)
    }
}
```
