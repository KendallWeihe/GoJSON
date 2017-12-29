
package main

import (
  "fmt"
  "reflect"
)

// -----------------------------------------------
// ADD
// -----------------------------------------------

func (json *JSON) add(jtype string, key string, value interface{}) {
  if (jtype == "key-value") {
    if (json.key_values == nil) {
      json.key_values = make(map[string]string)
    }
    json.key_values[key] = value.(string)
  } else if (jtype == "key-json") {
    if (json.json == nil) {
      json.json = make(map[string]*JSON)
    }
    json.json[key] = value.(*JSON)
  } else if (jtype == "key-list") {
    if (json.list == nil) {
      json.list = make(map[string]*JSONList)
    }
    json.list[key] = value.(*JSONList)
  } else {
    fmt.Println("oops...")
  }
}

func (list *JSONList) add(jltype string, value interface{}) {
  if (jltype == "value") {
    list.values = append(list.values, value.(string))
  } else if (jltype == "json") {
    list.json_objs = append(list.json_objs, value.(*JSON))
  } else {
    fmt.Println("hmm...")
  }
}

// -----------------------------------------------
// GET
// -----------------------------------------------

func (json *JSON) get(key string) interface{} {
  for k, v := range json.key_values {
    if (k == key) {
      return v
    }
  }

  for k, v := range json.json {
    if (k == key) {
      return v
    }
  }

  for k, v := range json.list {
    if (k == key) {
      return v
    }
  }

  return nil
}

// -----------------------------------------------
// SET
// -----------------------------------------------

func (json *JSON) set(key string, new_value interface{}) bool {
  for k, v := range json.key_values {
    if (k == key) {
      fmt.Println(reflect.TypeOf(v))
      return true
    }
  }

  for k, _ := range json.json {
    if (k == key) {
      return true
    }
  }

  for k, _ := range json.list {
    if (k == key) {
      return true
    }
  }

  return false
}

func (json_list JSONList) set(existing_json *JSON) {

}

// -----------------------------------------------
// DELETE
// -----------------------------------------------

func (json JSON) delete(existing_json *JSON) {

}

func (json_list JSONList) delete(existing_json *JSON) {

}







// -----------------------------------------------
// EDIT
// -----------------------------------------------

func (json JSON) edit(existing_json *JSON) {

}

func (json_list JSONList) edit(existing_json *JSON) {

}
