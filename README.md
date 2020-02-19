# OAI object model [![Build Status](https://travis-ci.org/go-openapi/spec3.svg?branch=master)](https://travis-ci.org/go-openapi/spec3) [![codecov](https://codecov.io/gh/go-openapi/spec3/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/spec3) [![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/spec3/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/spec3?status.svg)](http://godoc.org/github.com/go-openapi/spec3)

***This repository is not usable at this moment, the implementation is incomplete***

The object model for OpenAPI specification v3 documents.

It aims to fix some mistakes that were made in the spec for v2. Top-level maps are now sorted by default so you can rely on their ordering.

## Schemas

| Schema                 | `struct{}` + `map`                              | Unit tests (`struct{}` + `map`)               |
| ---------------------- | ----------------------------------------------- | --------------------------------------------- |
| OpenAPI                | :ballot_box_with_check:                         | :black_square_button:                         |
| Info                   | :ballot_box_with_check:                         | :black_square_button:                         |
| Contact                | :ballot_box_with_check:                         | :black_square_button:                         |
| License                | :ballot_box_with_check:                         | :black_square_button:                         |
| Server                 | :ballot_box_with_check:                         | :black_square_button:                         |
| Server Variable        | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Components             | :ballot_box_with_check:                         | :black_square_button:                         |
| Paths                  | :ballot_box_with_check:                         | :black_square_button:                         |
| Path Item              | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Operation              | :ballot_box_with_check:                         | :black_square_button:                         |
| External Documentation | :ballot_box_with_check:                         | :black_square_button:                         |
| Parameter              | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Request Body           | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Media Type             | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Encoding               | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Responses              | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Response               | :ballot_box_with_check:                         | :black_square_button:                         |
| Callback               | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Example                | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Link                   | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Header                 | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| Tag                    | :ballot_box_with_check:                         | :black_square_button:                         |
| Reference              | :ballot_box_with_check:                         | :black_square_button:                         |
| Schema                 | :black_square_button: + :black_square_button:   | :black_square_button: + :black_square_button: |
| Discriminator          | :ballot_box_with_check:                         | :black_square_button:                         |
| XML                    | :ballot_box_with_check:                         | :black_square_button:                         |
| Security Scheme        | :ballot_box_with_check: + :black_square_button: | :black_square_button: + :black_square_button: |
| OAuth Flows            | :ballot_box_with_check:                         | :black_square_button:                         |
| OAuth Flow             | :ballot_box_with_check:                         | :black_square_button:                         |
| Security Requirement   | :ballot_box_with_check:                         | :black_square_button:                         |

## TODO

- [ ] Update OrderedMap to use ART under the hood instead of the golang's map
- [ ] Use GoJay for decoding/encoding the JSON
- [ ] Retire EasyJson dependency from the project
