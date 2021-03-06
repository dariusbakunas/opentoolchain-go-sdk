[![Build Status](https://travis-ci.com/dbakuna/opentoolchain-go-sdk.svg?token=eW5FVD71iyte6tTby8gr&branch=main)](https://travis.ibm.com/dbakuna/opentoolchain-go-sdk.svg?token=eW5FVD71iyte6tTby8gr&branch=main)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Open Toolchain Go SDK 0.0.1
Go client library to interact with the various [IBM Cloud Open Toolchain Services APIs](https://cloud.ibm.com/apidocs?category=devops).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [IBM Cloud Open Toolchain Go SDK 0.0.1](#ibm-cloud-open-toolchain-go-sdk-001)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
    - [Go modules](#go-modules)
    - [`go get` command](#go-get-command)
  - [Using the SDK](#using-the-sdk)
  - [Questions](#questions)
  - [Issues](#issues)
  - [Open source @ IBM](#open-source--ibm)
  - [Known issues](#known-issues)
  - [Contributing](#contributing)
  - [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Open Toolchain Services Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
<!-- [Example Service](https://cloud.ibm.com/apidocs/example-service) | exampleservicev1 -->

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.14 or above.

## Installation
The current version of this SDK: 0.0.1

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/dariusbakunas/opentoolchain-go-sdk/opentoolchainv1"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `exampleservicev1` part of the import path is the package name
associated with the Example Service service.
See the service table above to find the approprate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/dariusbakunas/opentoolchain-go-sdk/opentoolchainv1
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Generating SDK

```bash
$ openapi-sdkgen.sh generate -g ibm-go -i docs/openapi.yaml -o .
$ go fmt ./...
$ patch opentoolchainv1/open_toolchain_v1.go < opentoolchainv1/patch
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](github.com/dariusbakunas/opentoolchain-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)


## Known issues

* Current openapi generator does not yet support `application/x-www-form-urlencoded` body type, fix by applying `patch` file in opentoolchainv1 folder:

```bash
patch open_toolchain_v1.go < patch
```

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
