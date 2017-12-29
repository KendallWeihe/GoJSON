
package main

import (
    "strings"
)

type KeyValue struct {
  key string
  value string
}

type JSONList struct {
  values []string // list of values
  json_objs []*JSON // list of json objects
}

type JSON struct {
  key_value []KeyValue // key/value map
  json map[string]*JSON // key/json object map
  list map[string]*JSONList // key/json list map
}

type JSONCRUD interface {
  set()
  add()
  get()
  edit()
  delete()
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func find_recursive(keys []string, index int, json JSON) interface{} {
  // check if the key even exists
  key := keys[index]
  if found, jtype, value := json.get(key); found {
    if (jtype == "key-value") {
      return value
    } else if (jtype == "json") {
      if (index == (len(keys)-1)) {
        return value
      } else {
        return find_recursive(keys, index+1, value.(JSON))
      }
    } else if (jtype == "list") {
      return value // NOTE: for now, if the key includes a list type, it should be the last key element
    } else {
      return nil
    }
  } else {
    return nil
  }
}

func find(key string, json JSON) interface{} {
  keys := strings.Split(key, ".")
  value := find_recursive(keys, 0, json)
  return value
}










// ...
