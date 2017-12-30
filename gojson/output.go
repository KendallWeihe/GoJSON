
package gojson

import (
    "fmt"
    "os"
    "strings"
)

func get_indent(indent_count int) string {
  indent := ""
  i := 0
  for i < indent_count {
    indent += "\t"
    i += 1
  }
  return indent
}

func write_json_list(list JSONList, indent_count int) string {
  indent := get_indent(indent_count)
  output_str := fmt.Sprintf("%s[\n", indent)
  indent_count += 1
  indent = get_indent(indent_count)

  for _, v := range list.values {
    output_str += fmt.Sprintf("%s\"%s\",\n", indent, v)
  }

  for _, json := range list.json_objs {
    output_str += fmt.Sprintf("%s%s,\n", indent, write_json(*json, indent_count+1))
  }

  indent_count -= 1
  indent = get_indent(indent_count)
  output_str = strings.TrimSuffix(output_str, ",\n")
  output_str += fmt.Sprintf("\n%s]", indent)
  return output_str
}

func write_json(custom_json JSON, indent_count int) string {

  indent := get_indent(indent_count)
  output_str := fmt.Sprintf("%s{\n", indent)
  indent_count += 1
  indent = get_indent(indent_count)

  for k, v := range custom_json.key_values {
    output_str += fmt.Sprintf("%s\"%s\": \"%s\",\n", indent, k, v)
  }

  for k, json := range custom_json.json {
    output_str += fmt.Sprintf("%s\"%s\":\n%s,\n", indent, k, write_json(*json, indent_count+1))
  }

  for k, list := range custom_json.list {
    output_str += fmt.Sprintf("%s\"%s\":\n%s,\n", indent, k, write_json_list(*list, indent_count+1))
  }

  indent_count -= 1
  indent = get_indent(indent_count) // TODO: remove the last comma
  output_str = strings.TrimSuffix(output_str, ",\n")
  output_str += fmt.Sprintf("\n%s}", indent)
  return output_str
}

func write(custom_json JSON, path string) bool {

  output_str := write_json(custom_json, 0)
  file, err := os.Create(path)
  check(err)
  defer file.Close()
  _, err = file.WriteString(output_str)
  check(err)
  fmt.Printf("Saved to file!\n")
  return true
}
