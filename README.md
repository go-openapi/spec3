# OAI object model [![Build Status](https://travis-ci.org/go-openapi/spec3.svg?branch=master)](https://travis-ci.org/go-openapi/spec3) [![codecov](https://codecov.io/gh/go-openapi/spec3/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/spec3) [![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/spec3/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/spec3?status.svg)](http://godoc.org/github.com/go-openapi/spec3)

***This repository is not usable at this moment, the implementation is incomplete***

The object model for OpenAPI specification v3 documents.

It aims to fix some mistakes that were made in the spec for v2. Top-level maps are now sorted by default so you can rely on their ordering.

## Schemas

| Schema                 | `struct{}` | Unit tests |
| ---------------------- | ---------- | ---------- |
| OpenAPI                | [x]        | [ ]        |
| Info                   | [x]        | [ ]        |
| Contact                | [x]        | [ ]        |
| License                | [x]        | [ ]        |
| Server                 | [x]        | [ ]        |
| Server Variable        | [x]        | [ ]        |
| Components             | [x]        | [ ]        |
| Paths                  | [x]        | [ ]        |
| Path Item              | [x]        | [ ]        |
| Operation              | [x]        | [ ]        |
| External Documentation | [x]        | [ ]        |
| Parameter              | [x]        | [ ]        |
| Request Body           | [x]        | [ ]        |
| Media Type             | [x]        | [ ]        |
| Encoding               | [x]        | [ ]        |
| Responses              | [x]        | [ ]        |
| Response               | [x]        | [ ]        |
| Callback               | [x]        | [ ]        |
| Example                | [x]        | [ ]        |
| Link                   | [x]        | [ ]        |
| Header                 | [x]        | [ ]        |
| Tag                    | [x]        | [ ]        |
| Reference              | [x]        | [ ]        |
| Schema                 | [ ]        | [ ]        |
| Discriminator          | [x]        | [ ]        |
| XML                    | [x]        | [ ]        |
| Security Scheme        | [x]        | [ ]        |
| OAuth Flows            | [x]        | [ ]        |
| OAuth Flow             | [x]        | [ ]        |
| Security Requirement   | [x]        | [ ]        |
