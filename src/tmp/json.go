
package main

import (
    "strings"
)

type KeyValue struct {
  key_value map[string]string
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

func find_json_recursive(custom_json *JSON, keys []string, key_index int) (*string, *JSON, *JSONList) {

  empty_str := new(string)
  empty_json := new(JSON)
  empty_json_list := new(JSONList)

  // value found
  key := keys[key_index]
  if found := custom_json.exists(key); found {
    value := custom_json.get(key)
    return &value, empty_json, empty_json_list
  }

  // NOTE: HERE!
  // nested json found
  // if found := custom_json.exists()

  if nested_json, ok := custom_json.json_nested[keys[key_index]]; ok {
    // if the key requested is in fact a JSON object
    if key_index == (len(keys)-1) {
      return empty_str, nested_json, empty_json_list
    }
    // keep searching in nested JSON
    return find_json_recursive(nested_json, keys, key_index+1)
  }

  // json list
  if json_list, ok := custom_json.json_list[keys[key_index]]; ok {
    // if the key is the last key in the key_string
    if key_index == (len(keys)-1) {
      return empty_str, empty_json, json_list
    }
    // if there is more to find ... TODO this is shit?
    json_objs := custom_json.json_list[keys[key_index]].json_objs
    for _, json_obj := range json_objs {
      return find_json_recursive(json_obj, keys, key_index+1)
    }
  }

  return empty_str, empty_json, empty_json_list
}

func find(custom_json *JSON, key string) (*string, *JSON, *JSONList) {
  keys := strings.Split(key, ".")
  value, json, list := find_json_recursive(custom_json, keys, 0)
  return value, json, list
}











// ...
