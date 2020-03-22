meili [![CircleCI](https://circleci.com/gh/fortytw2/meili.svg?style=svg)](https://circleci.com/gh/fortytw2/meili) [![Documentation](https://godoc.org/github.com/fortytw2/meili?status.svg)](http://godoc.org/github.com/fortytw2/meili)
-------

`meili` is a go package for using [MeiliSearch](https://meilisearch.com)

It has no dependencies outside the go standard library and is designed to be lightweight and easy to use.

#### comparison to [meilisearch-go](https://github.com/meilisearch/meilisearch-go)

- zero external dependencies
- no 'builder' pattern-ish API.
- cleaner API through use of functional options instead of configuration structs
- use of `json.RawMessage` for user supplied documents instead of reflection/automatic JSON encoding
    - this means no `Hits []interface{}` where the interface is unable to be your actual struct
    - you can also use your own higher-performance json encoding library without issue.
