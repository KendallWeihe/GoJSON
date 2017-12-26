package main

import (
  "fmt"
  "reflect"
)


func main() {

  // READ EXAMPLE FILE ------------------------------------
  config_file := "../test/example.json"
  json := new(JSON) // TODO: use create() to create JSON
  read(config_file, json)

  // FIND KEY-VALUE IN JSON
  key_value, _, _ := find(json, "key_value_1")
  fmt.Println(reflect.TypeOf(key_value))

  // FIND EMBEDDED KEY-VALUE IN JSON
  embedded_key_value, _, _ := find(json, "key_json.key_value_1")
  fmt.Println(reflect.TypeOf(embedded_key_value))

  // FIND KEY-LIST IN JSON
  _, _, key_list := find(json, "key_list")
  fmt.Println(reflect.TypeOf(key_list))

  // FIND EMBEDDED KEY-LIST IN JSON
  _, _, embedded_key_list := find(json, "key_json.key_list")
  fmt.Println(reflect.TypeOf(embedded_key_list))

  // FIND KEY-JSON IN JSON
  _, key_json, _ := find(json, "key_json")
  fmt.Println(reflect.TypeOf(key_json))

  // WRITE EXAMPLE TO NEW FILE
  output_file := "../test/example_output.json"
  write(*json, output_file)

  // TODO: "CRUD-like" actions for JSON values

}
