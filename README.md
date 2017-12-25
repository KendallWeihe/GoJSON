# GoJSON
Custom JSON package

# Files:
  - `package/`
    - `input.go`
    - `manipulate.go`
    - `output.go`
  - `playground/`
    - `example.json`
    - `play.go`
  - `README.md`

# The idea:
I didn't like [this](https://golang.org/pkg/encoding/json/), I came from Python and I primarily wanted to manipulate JSON data like from [here](https://docs.python.org/2/library/json.html), so I built `*this`.

In other words, I wanted to read a JSON file, manipulate data, and then write to an output file.

##### Current caveat's: every value is stores as a string & performance optimization is not currently taken into account

There are three types of JSON "keys":
  1. key-value
  2. key-list
  3. key-JSON (nested)

There are two types of JSON list elements:
  1. value
  2. JSON

# Setup:
  - `git clone ...`
  - copy files from `package/` to your project
  - include files in your build

# Test build:
  - ``

# Use:

### Read:
  - ...

### Manipulate:
  - ...

## Write:
  - ...

# TODO:
  - build automation
  - type support (other than strings)
