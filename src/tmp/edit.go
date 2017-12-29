
package main

import (
  "fmt"
  "reflect"
)

// -----------------------------------------------
// ADD
// -----------------------------------------------

func (json *JSON) add(jtype string, key string, value interface{}) {
  // if (jtype == "key-value") {
  //
  // }


  fmt.Println("Value type: %s", reflect.TypeOf(value))
  switch v := value.(type) {
    case string:
      if (json.key_values == nil) {
        json.key_values = make(map[string]string)
      }
      json.key_values[key] = v
      fmt.Println("Case string")
      fmt.Println(v)
    case JSON:
      if (json.json == nil) {
        json.json = make(map[string]*JSON)
      }
      json.json[key] = &v
      fmt.Println("Case JSON")
      fmt.Println(v)
    case JSONList:
      if (json.list == nil) {
        json.list = make(map[string]*JSONList)
      }
      json.list[key] = &v
      fmt.Println("Case JSONList")
      fmt.Println(v)
    default:
      fmt.Println("hmm...")
      fmt.Println(v)
  }
}

func (list *JSONList) add(value interface{}) {
  switch v := value.(type) {
    case string:
      list.values = append(list.values, v)
    case JSON:
      list.json_objs = append(list.json_objs, &v)
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
      fmt.Println(reflect.TypeOf(new_value))
      // json.key_value[idx] = new_value.value
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
