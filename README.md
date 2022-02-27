# Toga

A JSON-based rules engine in Go.

## Why?

This can be used to store configuration data when you require something more dynamic than a static key/value mapping. Feature flags are one use case for something like this.

For example, let's say you just added a new feature to your application. This may have been a very complex feature and you would like to only show it to a subset of your users at first to make sure that everything is working correctly. For such a situation, you may create a rule like this:
```javascript
{
  "eq": [
    { "context": "beta" },
    true
  ]
}
```

This rule evaluates to true if the context passed in has a `beta` value of true.

```bash
$ ./toga eval -rule '{"eq": [{"context": "beta"}, true]}' -context '{"beta": true}'
true

$ ./toga eval -rule '{"eq": [{"context": "beta"}, true]}' -context '{"beta": false}'
false
```

### More complex rules

You can also use if/then logic in your rules and return things other than booleans.

```bash
$ ./toga eval -rule '{"if": {"condition": {"gt": { "first": {"context": "accountAge"}, "second": 365.0 }}, "then": "Hello, old friend!", "else": "Welcome, new friend!" } }' -context '{"accountAge": 400.0}'
"Hello, old friend!"
```

```bash
$ ./toga eval -rule '{"if": {"condition": {"gt": { "first": {"context": "accountAge"}, "second": 365.0 }}, "then": "Hello, old friend!", "else": "Welcome, new friend!" } }' -context '{"accountAge": 12.0}' 
"Welcome, new friend!"
```

### Rule files

Sometimes, you may want to read a rule from a file instead of specifying it directly as a string. To do so, you can use the `rule-file` parameter like so:

```bash
$ cat << EOF >> rule.json
{
  "if": {
    "condition": {
      "gt": {
        "first": {"context": "accountAge"},
        "second": 365.0 
      }
    },
    "then": "Hello, old friend!",
    "else": "Welcome, new friend!"
  }
}
EOF
$ ./toga eval -rule-file rule.json -context '{"accountAge": 12.0}'
"Welcome, new friend!"
$ ./toga eval -rule-file rule.json -context '{"accountAge": 999.0}'
"Hello, old friend!"
```

## Go API

To use this library within your Go code, the main interface is in the sdk package. The functionality is exposed through the `EvalRuleFile` function. See the [todo-list application](./examples/todo-app/) for a full working example.

## Examples
See the [examples](./examples) folder for some usage examples. There is a small [todo-list application](./examples/todo-app/) which shows how you can use this in a Go application.