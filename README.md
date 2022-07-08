# humannumbers
[![humannumbers](https://github.com/kmulvey/humannumbers/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/humannumbers/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/humannumbers)](https://goreportcard.com/report/github.com/kmulvey/humannumbers) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/humannumbers.svg)](https://pkg.go.dev/github.com/kmulvey/humannumbers)

Convert numbers in the form of words to float64

## Example:
```
var number, err = humannumbers.Parse("three million eight hundred and ninety four thousand seven hundred and sixty five")
// number == 3,894,765.0
```

## Limitations
- English Only

## How it works
- input: "three million eight hundred ninety four thousand seven hundred five"
- Remove unnecessary words (and)
- if there are decimals, just smash them together behind a '.' and add it to the result of below
- Parse each word into a number (int)
- []int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 5}
- Go though this array and apply rules for addition and multiplication
- result: float64(3_894_705)
