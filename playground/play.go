package main

import (
    "fmt"
    "os"
)


func main() {
  
  // READ EXAMPLE FILE ------------------------------------
  config_file := "example.json"
  example := new(JSON)
  input(config_file, example)

  // CREATE NEW KEY-VALUE
  // ...
  //
  // CREATE NEW KEY-LIST
  // ...
  //
  // CREATE NEW KEY-JSON
  // ...
  //
  // GET VALUE
  // ...
  //
  // GET LIST
  // ...
  //
  // GET JSON
  // ...
  //
  // UPDATE VALUE
  // ...
  //
  // UPDATE LIST
  // ...
  //
  // UPDATE JSON
  // ...
  //
  // DELETE KEY-VALUE
  // ...
  //
  // DELETE KEY-LIST
  // ...
  //
  // DELETE KEY-JSON
  // ...

  // WRITE EXAMPLE TO NEW FILE
  output_file := "playground/example_output.json"
  output(output_file, example)

}
