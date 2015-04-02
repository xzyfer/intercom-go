package intercom

// import (
//     "fmt"
// )

// Tag represents a tag.
//
// See https://doc.intercom.io/api/#tags for more information.
type Tag struct {
    ID      string
    Name    string
}

// ListTags returns all of the tags that belong to a client.
//
// See https://doc.intercom.io/api/#tags for more information.
func (c *APIClient) ListTags() ([]*Tag, error) {
    req, err := c.NewRequest("GET", "tags", nil)
    if err != nil {
        return nil, err
    }

    var v map[string]interface{}
    err = c.Do(req, &v)
    if err != nil {
        return nil, err
    }

    vv := v["tags"].([]interface{})

    var tags = make([]*Tag, len(vv));
    for i, vvv := range vv {
        // fmt.Printf("%d - %+v\n", i, vvv)
        // fmt.Printf("%d - %+v\n", i, vvv)
        vvvv := vvv.(map[string]interface{})
        // fmt.Printf("%d - %+v\n", i, vvvv)
        // fmt.Printf("%s\n", vvvv["name"])
        tag := Tag{vvvv["id"].(string), vvvv["name"].(string)}
        tags[i] = &tag
    }

    return tags, err
}
