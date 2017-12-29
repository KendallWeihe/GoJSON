
package main

// -----------------------------------------------

func (kv KeyValue) set(key string, value string) {
  // kv.key_value = make(map[string]string)
  // kv.key_value[key] = value
  kv.key = key
  kv.value = value
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

func (json JSON) get(key string) (bool, string, interface{}) {
  for _, kv := range json.key_value {
    if (kv.key == key) {
      return true, "key-value", kv.value
    }
  }

  for k, v := range json.json {
    if (k == key) {
      return true, "json", *v
    }
  }

  for k, v := range json.list {
    if (k == key) {
      return true, "list", *v
    }
  }

  return false, "", nil
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
