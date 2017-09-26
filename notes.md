## Access nested json object with type assertion

```javascript

{
    "name": "Karim",
    "address": {
        "street": "Rue Madame"
    }
}

```

Normally, to access the 'street' key we'd do something like this:

```javascript
let karim = JSON.Parse({object})

let streetName = karim["address"]["street"] // Rue Madame
```

In golang we can't do that unfortunately. You'd get the following error by
trying to compile:

```bash
invalid operation: karim["address"]["name"] (type interface {} does not support indexing)
```

To remedy this, we have to use what is called type assertion, which (based on my
understanding, works like casting but for maps ?)

```
karim["address"].(map([string]interface{})["name"]
```
