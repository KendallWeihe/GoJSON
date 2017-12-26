# GoJSON
Custom JSON package

# The idea:
I didn't like [this](https://golang.org/pkg/encoding/json/), I came from Python and I primarily wanted to manipulate JSON data like from [here](https://docs.python.org/2/library/json.html), so I built `*this`.

In other words, I wanted to read a JSON file, manipulate data, and then write to an output file.

##### Current caveats: every value is stores as a string & performance optimization is not currently taken into account

There are three types of JSON key-value stores:
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

# Current functionality:
  - read from a JSON file
  - find elements within the JSON file
  - write to a JSON file

# Test:
  - `cd src`
  - `go build`
  - `./src`

# TODO:
  - restructure repo for GoLang GitHub tree structure standards
  - add "CRUD-like" functionality for JSON elements
  - support types other than strings wihtin the JSON object
  - `find()` should support embedded list elements
  - build automation
