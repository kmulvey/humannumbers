# humannumbers
[![humannumbers](https://github.com/kmulvey/humannumbers/actions/workflows/release_build.yml/badge.svg)](https://github.com/kmulvey/humannumbers/actions/workflows/release_build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/kmulvey/humannumbers)](https://goreportcard.com/report/github.com/kmulvey/humannumbers) [![Go Reference](https://pkg.go.dev/badge/github.com/kmulvey/humannumbers.svg)](https://pkg.go.dev/github.com/kmulvey/humannumbers)

Convert numbers in the form of words to ints

## Example:
```
var number, err = humannumbers.Parse("three million eight hundred and ninety four thousand seven hundred and sixty five")
// number == 3,894,765
```

## Limitations
- English Only, interested in i18n [help wanted]
- Does not support floats (yet)

## How it works
- input: "three million eight hundred ninety four thousand seven hundred five"
- Remove unnecessary words (and)
- Parse each word into a number (int)
- []int{3, 1e6, 8, 100, 90, 4, 1000, 7, 100, 5}
- Go though this array and apply rules for addition and multiplication
- result: int(3_894_705)
