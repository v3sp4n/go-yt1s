<h1>installing</h1>
>go get github.com/v3sp4n/go-yt1s
<h1>usage</h1>
```go
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
        k := yt1s.Download(urls[0],"720p","/home/vespan/Desktop/")
        fmt.Println(k)
    }
}
```
