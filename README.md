<h1>installing</h1>
>>>go get github.com/v3sp4n/go-yt1s
<h1>usage</h1>
<br>


```~golang
package main
import (
    "fmt"

    "github.com/v3sp4n/go-yt1s"
)
func main() {
    urls := []string{
        "https://www.youtube.com/watch?v=F5RnAFl_gz0",
        "https://www.youtube.com/watch?v=OpNNmGs8Xao",
    }
    for i := range urls{
        q,err := yt1s.GetAvalibleQualites(urls[i])
        fmt.Println("avalible qualites for",urls[i],q,err)
        err2 := yt1s.Download(urls[0],"720p","/home/vespan/Desktop/")
        fmt.Println(err2)
    }
}
```
s
