
package main

import (
    "strings"
    "fmt"
    "reflect"
)

type JSONList struct {
  values []string // list of values
  json_objs []*JSON // list of json objects
}

type JSON struct {
  key_values map[string]string
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

func find_recursive(keys []string, index int, json *JSON) interface{} {
  // check if the key even exists
  key := keys[index]
  value := json.get(key)
  vtype := fmt.Sprintf("%s", reflect.TypeOf(value))

  if (vtype == "string") {
    return value
  } else if (vtype == "*main.JSON") {
    if (index == (len(keys)-1)) {
      return value
    } else {
      return find_recursive(keys, index+1, value.(*JSON))
    }
  } else if (vtype == "*main.JSONList") {
    return value
  } else {
    fmt.Println("hmmmmm")
    return nil
  }
}

func find(key string, json *JSON) interface{} {
  keys := strings.Split(key, ".")
  value := find_recursive(keys, 0, json)
  return value
}

func set_recursive(keys []string, index int, json *JSON, new_value interface{}) bool {
  key := keys[index]
  if found := json.set(key, new_value); found {
    fmt.Println(reflect.TypeOf(found))
  } else {
    if ((len(keys)-1) == index) {
      return false
    } else {
      return set_recursive(keys, index+1, json, new_value)
    }
  }

  return false
}

func set(key string, json *JSON, value interface{}) bool {
  keys := strings.Split(key, ".")
  set := set_recursive(keys, 0, json, value)
  return set
}











// ...
