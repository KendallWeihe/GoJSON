

package main

import (
  "fmt"
  "reflect"
)

func find_functions(json *JSON) {
  // -----------------------------------------------
  // FIND
  // -----------------------------------------------

  // FIND KEY-VALUE IN JSON
  value := find("key_1", json)
  fmt.Println(reflect.TypeOf(value))
  fmt.Println(value.(string))

  // FIND NESTED KEY-VALUE IN JSON
  nested_value := find("nested_json.nested_key_1", json)
  fmt.Println(reflect.TypeOf(nested_value))
  fmt.Println(nested_value.(string))

  // FIND KEY-LIST IN JSON
  list := find("list", json)
  fmt.Println(reflect.TypeOf(list))
  fmt.Println(list.(*JSONList))

  // FIND NESTED KEY-LIST IN JSON
  nested_list := find("nested_json.nested_list", json)
  fmt.Println(reflect.TypeOf(nested_list))

  // FIND KEY-JSON IN JSON
  nested_json := find("nested_json", json)
  fmt.Println(reflect.TypeOf(nested_json))

  // FIND NESTED KEY-JSON IN JSON
}

func set_functions(json *JSON) {
  // -----------------------------------------------
  // SET
  // -----------------------------------------------

  // set_key := "set_key"
  // set_value := "set_value"

  // SET A KEY-VALUE VALUE (ROOT LEVEL)
  // _ = set("key_2", json, kv)

  // SET A NESTED KEY-VALUE VALUE

  // SET A KEY-LIST VALUE

  // SET A NESTED KEY-LIST VALUE

  // SET A KEY-JSON VALUE

  // SET A NESTED KEY-JSON VALUE
}

func main() {
  // READ EXAMPLE FILE ------------------------------------
  config_file := "../../test/example.json"
  json := new(JSON) // TODO: use create() to create JSON
  read(config_file, json)

  find_functions(json)
  // set_functions(json)

  // WRITE EXAMPLE TO NEW FILE
  output_file := "../../test/example_output.json"
  write(*json, output_file)
}
