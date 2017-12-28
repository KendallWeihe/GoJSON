

package main

import (
  "fmt"
  "reflect"
)

func main() {
  // READ EXAMPLE FILE ------------------------------------
  config_file := "../../test/example.json"
  json := new(JSON) // TODO: use create() to create JSON
  read(config_file, json)

  // FIND KEY-VALUE IN JSON
  value, _, _ := find(json, "key_1")
  fmt.Println(reflect.TypeOf(value))

  // FIND EMBEDDED KEY-VALUE IN JSON
  nested_value, _, _ := find(json, "nested_json.nested_key_1")
  fmt.Println(reflect.TypeOf(nested_value))

  // FIND KEY-LIST IN JSON
  _, _, list := find(json, "list")
  fmt.Println(reflect.TypeOf(list))

  // FIND EMBEDDED KEY-LIST IN JSON
  _, _, nested_list := find(json, "nested_json.nested_list")
  fmt.Println(reflect.TypeOf(nested_list))

  // FIND KEY-JSON IN JSON
  _, nested_json, _ := find(json, "nested_json")
  fmt.Println(reflect.TypeOf(nested_json))

  // WRITE EXAMPLE TO NEW FILE
  output_file := "../test/example_output.json"
  write(*json, output_file)
}
