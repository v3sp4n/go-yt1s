package yt1s

import (
    // "fmt"
    // "log"
    "errors"
    "io"
    "os"
    "net/http"
    "net/url"
    "encoding/json"
)

func checkStatusInResponse(json_string string) (*bool, *string, error) {
    type json__struct struct {
        Status string
        C_Status string
        Mess string
    }
    j := json__struct{}
    err := json.Unmarshal([]byte(json_string),&j)
    if err != nil {
        return nil,nil,err
    }
    if j.C_Status == "FAILED" {
        r := false
        return &r,&j.Mess,nil
    }
    r := true
    return &r,&j.Mess,nil
}

func getDlink(vid,k string) (*string, error) {
    type getDlink__struct struct {
        _ string
        _ string
        _ string
        _ string
        _ string
        _ string
        _ string
        Dlink string
    }
    resp, err := http.PostForm("https://yt1s.com/api/ajaxConvert/convert", url.Values{
        "vid":{vid},
        "k": {k},
    })
    if resp.StatusCode != 200 {
        return nil,errors.New("StatusCode != 200!")
    }
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil,err
    }
    status,mess,err := checkStatusInResponse(string(body))//check Body.Status "ok"
    if err != nil || *status == false {
        if err != nil {
            return nil,err
        }
        return nil,errors.New("body.Status != ok! " + *mess)
    }
    j := getDlink__struct{}
    err = json.Unmarshal([]byte(string(body)), &j)
    if err != nil {
        return nil,err
    }
    return &j.Dlink,nil
}
func getK(video, resolution, format string) (*map[string]string, error) {
    type getK__struct struct {
        Status string
        Mess string
        P string
        Vid string
        Title string
        T int
        A string
        Links map[string]map[string]map[string]string
    }
    formats := map[string]string {
        "1080p": "137",
        "720p": "22",
        "480p": "135",
        "360p": "18",
        "240p": "133",
        "144p": "160",
        "mp3": "mp3128",
    }
    formatv,formatb := formats[resolution]
    if formatb == false {
        return nil,errors.New("undefined format(mp3/mp4(1080p,720p,480p,360p,240p,144p))")
    }
    resp, err := http.PostForm("https://yt1s.com/api/ajaxSearch/index", url.Values{
        "q": {video},
        "vt": {formatv},
    })
    if resp.StatusCode != 200 {
        return nil,errors.New("StatusCode != 200!")
    }
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    status,mess,err := checkStatusInResponse(string(body))//check Body.Status "ok"
    if err != nil || *status == false {
        if err != nil {
            return nil,err
        }
        return nil,errors.New("body.Status != ok! " + *mess)
    }
    j := getK__struct{}
    err = json.Unmarshal([]byte(string(body)), &j)
    if err != nil {
        return nil, err
    }
    _,Links__err := j.Links[format][formatv]
    if Links__err == false {
        return nil,errors.New("undefined format/resolution")
    }
    res := map[string]string {
        "videoname": j.Title,
        "videoid": j.Vid,
        "key": j.Links[format][formatv]["k"],
    }
    return &res, nil
}

func downloadFromDlink(dlink,filename,format string) (error) {
    res,err := http.Get(dlink)
    if err != nil {
        return err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }
    if format != "mp3" {
        format = "mp4"
    }
    err = os.WriteFile(filename+"."+format, body, 0644)
    //+
    return nil
}


func download(url, resolution, format string) (error) {
    if format == "mp3" || format == "avi" {
        resolution = "mp3"
    }
    vid_,err := getK(url, resolution, format)
    if err != nil {
        return err
    }
    vid := *vid_
    filename := vid["videoname"]
    dlink,err := getDlink(vid["videoid"],vid["key"])
    if err != nil {
        return err
    }
    downloadFromDlink(*dlink,filename,format)
    return nil
}
