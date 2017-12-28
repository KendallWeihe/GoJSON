# GoJSON
Custom JSON package

# The idea:
Wasn't a fan of [this](https://golang.org/pkg/encoding/json/), I came from Python and I primarily wanted to handle JSON data in an easy-to-use way, like from [here](https://docs.python.org/2/library/json.html), so I built `*this` (plus, this was part of my "learning Go" project).

That is to say, with a simple set of function calls, I wanted to do any of the following...
  - read from a JSON file
  - create a JSON object from scratch
  - manipulate JSON elements
  - write JSON data to a file

##### Please refer to **Current functionality** -- this project is still under construction

There are three types of JSON key-value stores:
  1. key-value
  2. key-list
  3. key-JSON (nested)

There are two types of JSON list elements:
  1. value
  2. JSON

# Current functionality:
  - read from a JSON file
  - find elements within the JSON file
  - write to a JSON file

# Test:
  - `git clone ...`
  - `cd GoJSON`
  - `cd src`
  - `go build`
  - `./src`

# TODO:
  - add functionality
  - add "How to use" to README
  - restructure repo for GoLang GitHub tree structure standards
  - add "CRUD-like" functionality for JSON elements (including creating an object from scratch)
  - support types other than strings within the JSON object
  - `find()` should support embedded list elements
  - build automation
