# Design

## Logic:
  - instantiate json object
  - then (maybe):
    - read from a file
    - CRUD json object
    - write to a file

## Files & functions:

### `json.go`:

##### Types:

```[golang]
type JSONList struct {
  values []string // list of values
  json_objs []*JSON // list of json objects
}

type JSON struct {
  key_value map[string]string // key/value map
  json_nested map[string]*JSON // key/json object map
  json_list map[string]*JSONList // key/json list map
}
```
##### Functions:

`init()`

Parameters:
  - ...

Returns:
  - ...



...
