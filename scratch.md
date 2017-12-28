# Scratch work


## User functionality:
  - read from a file
  - create a new JSON object
  - add keys to JSON object
  - get values from JSON object
  - edit values from JSON object
  - delete from object
  - write to file

## Note on keys:
  - keys are deliminated by "." (but this should be easily defined by the user)

## Functions:
  - read() from file
  - add() to object
  - get() from object
  - edit() object
  - delete() object
  - write() to file

## Implementation:
  - interface for each of the CRUD functions
    - types:
      - JSON
      - JSONList
      - key-value

## Interfaces:
  - the way it would work...
    ```[golang]
      type JSONCRUD interface {
        add()
        get()
        edit()
        delete()
      }

      new_key_value := make(map[string]string)
      new_key_value["key1"] = "value1"
      new_key_value.add(existing_json_object)

      new_key_json := make(map[string]JSON)
      new_json_value := new(JSON)
      new_key_json["key2"] = new_json_value
      new_key_json.add(existing_json_object)

      ...same for JSONList
    ```
