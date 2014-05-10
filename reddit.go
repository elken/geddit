package reddit

import (
    "fmt"
    "net/http"
    "encoding/json"
    "errors"
)

// Describes the reddit JSON
type Item struct {
    Title       string
    URL         string
    Comments    int `json:"num_comments"`
    Ups         int
    Downs       int
}

type response struct {
    Data struct {
        Children []struct {
            Data Item
        }
    }
}

func (i Item) String() string {
    com := ""
    switch i.Comments {
        case 0:
        // Nothing
        case 1:
            com = " (1 comment)"
        default:
            com = fmt.Sprintf(" (%d comments)", i.Comments)
    }
    return fmt.Sprintf("%s%s(%d:%d)\n%s", i.Title, com, i.Ups, i.Downs, i.URL)
}

// Fetches most recent items in passed "reddit"
func Get(reddit string) ([]Item, error) {
    url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(resp.Status)
    }
    r := new(response)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return nil,err
    }
    items := make([]Item, len(r.Data.Children))
    for i,child := range r.Data.Children {
        items[i] = child.Data
    }
    return items, nil
}
