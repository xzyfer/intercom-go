package intercom

import (
    "encoding/json"
)

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

    var tags []*Tag
    for _, vvv := range vv {
        var tag Tag;
        j, err := json.Marshal(&vvv)
        if err != nil {
            return nil, err
        }
        err = json.Unmarshal(j, &tag)
        if err != nil {
            return nil, err
        }
        tags = append(tags, &tag)
    }

    return tags, err
}
