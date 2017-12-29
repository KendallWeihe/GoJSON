

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

  // jtype string, key string, json JSON

  // FIND KEY-VALUE IN JSON
  value := find("key_1", *json)
  fmt.Println(reflect.TypeOf(value))
  // fmt.Printf(value.(string))

  // FIND EMBEDDED KEY-VALUE IN JSON
  nested_value:= find("nested_json.nested_key_1", *json)
  fmt.Println(reflect.TypeOf(nested_value))

  // FIND KEY-LIST IN JSON
  list := find("list", *json)
  fmt.Println(reflect.TypeOf(list))

  // FIND EMBEDDED KEY-LIST IN JSON
  nested_list := find("nested_json.nested_list", *json)
  fmt.Println(reflect.TypeOf(nested_list))

  // FIND KEY-JSON IN JSON
  nested_json := find("nested_json", *json)
  fmt.Println(reflect.TypeOf(nested_json))

  // WRITE EXAMPLE TO NEW FILE
  output_file := "../../test/example_output.json"
  write(*json, output_file)
}
