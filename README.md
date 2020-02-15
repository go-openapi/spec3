# OAI object model [![Build Status](https://travis-ci.org/go-openapi/spec3.svg?branch=master)](https://travis-ci.org/go-openapi/spec3) [![codecov](https://codecov.io/gh/go-openapi/spec3/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/spec3) [![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/spec3/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/spec3?status.svg)](http://godoc.org/github.com/go-openapi/spec3)

***This repository is not usable at this moment, the implementation is incomplete***

The object model for OpenAPI specification v3 documents.

It aims to fix some mistakes that were made in the spec for v2. Top-level maps are now sorted by default so you can rely on their ordering.

## Schemas

| Schema                 | `struct{}`         | Unit tests         |
| ---------------------- | ------------------ | ------------------ |
| OpenAPI                | :heavy_check_mark: | :white_check_mark: |
| Info                   | :heavy_check_mark: | :white_check_mark: |
| Contact                | :heavy_check_mark: | :white_check_mark: |
| License                | :heavy_check_mark: | :white_check_mark: |
| Server                 | :heavy_check_mark: | :white_check_mark: |
| Server Variable        | :heavy_check_mark: | :white_check_mark: |
| Components             | :heavy_check_mark: | :white_check_mark: |
| Paths                  | :heavy_check_mark: | :white_check_mark: |
| Path Item              | :heavy_check_mark: | :white_check_mark: |
| Operation              | :heavy_check_mark: | :white_check_mark: |
| External Documentation | :heavy_check_mark: | :white_check_mark: |
| Parameter              | :heavy_check_mark: | :white_check_mark: |
| Request Body           | :heavy_check_mark: | :white_check_mark: |
| Media Type             | :heavy_check_mark: | :white_check_mark: |
| Encoding               | :heavy_check_mark: | :white_check_mark: |
| Responses              | :heavy_check_mark: | :white_check_mark: |
| Response               | :heavy_check_mark: | :white_check_mark: |
| Callback               | :heavy_check_mark: | :white_check_mark: |
| Example                | :heavy_check_mark: | :white_check_mark: |
| Link                   | :heavy_check_mark: | :white_check_mark: |
| Header                 | :heavy_check_mark: | :white_check_mark: |
| Tag                    | :heavy_check_mark: | :white_check_mark: |
| Reference              | :heavy_check_mark: | :white_check_mark: |
| Schema                 | :white_check_mark: | :white_check_mark: |
| Discriminator          | :heavy_check_mark: | :white_check_mark: |
| XML                    | :heavy_check_mark: | :white_check_mark: |
| Security Scheme        | :heavy_check_mark: | :white_check_mark: |
| OAuth Flows            | :heavy_check_mark: | :white_check_mark: |
| OAuth Flow             | :heavy_check_mark: | :white_check_mark: |
| Security Requirement   | :heavy_check_mark: | :white_check_mark: |
