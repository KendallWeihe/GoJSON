
package main

// -----------------------------------------------

func (kv KeyValue) set(key string, value string) {
  kv.key_value = make(map[string]string)
  kv.key_value["key"] = "value"
}

// -----------------------------------------------

func (kv KeyValue) add(parent_json *JSON) {
  parent_json.key_value = append(parent_json.key_value, kv)
}

func (json JSON) add(key string, parent_json *JSON) {
  if (parent_json.json == nil) {
      parent_json.json = make(map[string]*JSON)
  }
  parent_json.json[key] = &json
}

func (list JSONList) add(key string, parent_json *JSON) {
  if (parent_json.list == nil) {
      parent_json.list = make(map[string]*JSONList)
  }
  parent_json.list[key] = &list
}

// -----------------------------------------------

func (json JSON) exists(key string) bool {
  for _, kv := range json.key_value {
    if val, ok := kv.key_value[key]; ok {
      return true
    }
  }

  for _, json := range json.json {
    if val, ok := json[key]; ok {
      return true
    }
  }

  for _, list := range json.list {
    if val, ok := list.list[key]; ok {
      return true
    }
  }

  return false
}

// -----------------------------------------------

func (kv KeyValue) get(parent_json *JSON) {

}

func (json JSON) get(key string) string {
  for _, kv := range json.key_value {
    if val, ok := kv.key_value[key]; ok {
      return val
    }
  }
  return ""
}

func (json_list JSONList) get(parent_json *JSON) {

}

func (json JSON) get(key string) interface, string {

  // RET VALUE: 
  //   - typeof
  //   - value
}

// -----------------------------------------------

func (kv KeyValue) edit(existing_json *JSON) {

}

func (json JSON) edit(existing_json *JSON) {

}

func (json_list JSONList) edit(existing_json *JSON) {

}

// -----------------------------------------------

func (kv KeyValue) delete(existing_json *JSON) {

}

func (json JSON) delete(existing_json *JSON) {

}

func (json_list JSONList) delete(existing_json *JSON) {

}
