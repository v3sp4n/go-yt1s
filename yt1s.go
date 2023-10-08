//shit code..damn!

/*
error := yt1s.download(url, quality, path string)
&map[string][]map[string]string,error := yt1s.GetAvalibleQualites(url string)
*/
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

func getK__req(video string) (*string,error) {
    resp, err := http.PostForm("https://yt1s.com/api/ajaxSearch/index", url.Values{
        "q": {video},
        // "vt": {formatv},
    })
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != 200 {
        return nil,errors.New("StatusCode != 200!")
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
    sbody := string(body)
    return &sbody,nil
}

func getAvalibleQualites(links map[string]map[string]map[string]string) (*map[string][]map[string]string) {
    res := map[string][]map[string]string{}
/*
{
    "mp4": {
        {
            "quality":"720p",
            "key":"qwerty123....",
        },
    },
}
*/
    for f := range links {
        res[f] = []map[string]string{}
        for k := range links[f] {
            res[f] = append(res[f],map[string]string{
                "quality": links[f][k]["q"],
                "key_quality": k,
                "key": links[f][k]["k"],
                "format": links[f][k]["f"],
            })
        }
    }

    return &res
}
func GetAvalibleQualites(url string) (*map[string][]map[string]string,error) {
    body,err := getK__req(url)
    if err != nil {
        return nil,err
    }
    j := getK__struct{}
    err = json.Unmarshal([]byte(*body), &j)
    if err != nil {
        return nil, err
    }
    res := getAvalibleQualites(j.Links)
    return res,nil
}


func getK(video, quality/*mp4:140p,240p,...mp3:mp3*/ string) (*map[string]string, error) {
    resp, err := http.PostForm("https://yt1s.com/api/ajaxSearch/index", url.Values{
        "q": {video},
        // "vt": {formatv},
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
    undefined_format := true
    undefined_format__avalible_qualites := ""
    qq := getAvalibleQualites(j.Links)
    q := *qq
    for f := range q {
        for i := range q[f] {
            undefined_format__avalible_qualites = undefined_format__avalible_qualites + q[f][i]["quality"] + " "
        }
    }
    key := ""
    format := ""
    if quality == "mp3" {
        for k := range q["mp3"] {
            key = q["mp3"][k]["key"]
            format = q["mp3"][k]["format"]
            undefined_format = false
            break
        }
    } else {
        for f := range q {
            for i := range q[f] {
                if quality == q[f][i]["quality"] {
                    key = q[f][i]["key"]
                    format = q[f][i]["format"]
                    undefined_format = false
                }
            }
        }
    }
    if undefined_format {
        return nil,errors.New("undefined format/quality. Aavalible format&quality:"+undefined_format__avalible_qualites)
    }
    res := map[string]string {
        "videoname": j.Title,
        "videoid": j.Vid,
        "key": key,
        "format": format, 
    }
    return &res, nil
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
    err = os.WriteFile(filename+"."+format, body, 0644)
    //+
    return nil
}


func Download(url, resolution, path string) (error) {
    aVid,err := getK(url, resolution)
    if err != nil {
        return err
    }
    vid := *aVid
    filename := vid["videoname"]
    dlink,err := getDlink(vid["videoid"],vid["key"])
    if err != nil {
        return err
    }
    downloadFromDlink(*dlink,path+filename,vid["format"])
    return nil
}
