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
I didn't like [this](https://), I came from [this](https://), and I primarily wanted to manipulate JSON data like from [here](https://), so I built `*this`.

In other words, read a JSON file, manipulate data, write to an output file.

There are three types of JSON "keys":
  1. key-value
  2. key-list
  3. key-JSON (nested)

There are two types of JSON list elements:
  1. value
  2. JSON

##### Current caveat: every value is stores as a string

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
