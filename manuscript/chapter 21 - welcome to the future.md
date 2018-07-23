# Welcome to the Future

## Go Warts

- unchecked return values
- []T can't be converted directly to []interface{}
- possible silent overflow when converting between numeric types
  - https://github.com/golang/go/issues/19623
  - https://github.com/golang/go/issues/19624
- no exceptions
- no default values
- no overloaded methods
- no ternary operator
- poor visibility rules
- T and T* have different method sets
- [goroutines and closures](http://localhost:8080/doc/faq#closures_and_goroutines)
- i++ and i-- useless (just use for-range or i += 1) and confusing (it's a statement and not an expression like in C/C++)

## Go 2.0

## Plugins

## Stability
