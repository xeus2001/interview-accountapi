# Form3 Take Home Exercise

The original task can be found in the [Form3 Interview GitHub repository](https://github.com/form3tech-oss/interview-accountapi).

## Preamble

**This is my first go project ever.**

In the last 10+ years I only wrote Java applications and services. The last C application that I wrote was at an 80486,
so this was a great learning experience. No matter if the application is accepted or not; it was a great experience 
for me and nice to have some dedicated learning goal!

## Documentation

The client library documentation can be found [here](doc/README.md).

## Go 1.18

The code does uses generics, which are part of Go 1.18, therefore to build the code Go 1.18 is required. In most 
64-bit x86 Linux systems the installation can be done like:

```bash
cd ~/Downloads/
curl -LO https://go.dev/dl/go1.18beta1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.18beta1.linux-amd64.tar.gz
mkdir ~/go
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

## Docker

The library comes with a docker stack that can be started using `docker-compose up` (as requested in the _Shoulds_).

## Make

The following options are available:

- `make`: Create the library and demo binaries for the local docker stack (endpoint: `http://localhost:8080/v1`).
- `make test`: Run standard tests.
- `make test-int`: Run integration tests against docker stack.
- `make test-int-result`: Shows the test coverage as HTML file in the browser.
- `make release`: Create the library and binaries for production (endpoint: `https://api.f3.tech/v1`).
- `make clean`: Remove all build artifacts.
- `make doc`: Generate the documentation in the [doc/](doc/README.md) folder.
- `make fmt`: Apply auto-formatting.
- `make check`: Check files for correct formatting.

Some special commands are:

- `make test-docker`: Run integration tests from inside docker.
- `make docker`: Create the library and binaries for execution inside the docker (endpoint: `http://accountapi:8080/v1`).
- `make swagger-ui`: Requires a running docker stack and installed chromium, if available, runs chromium against a
  swagger-ui to test the local docker account API.

---

## Instructions

The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a
Docker container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with
the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the
file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first, before
reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**

- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against
  the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests
  pass.

### Should Nots

The finished solution **should not:**

- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in
  the test being rejected.
- Use a library for your client (e.g: go-resty). Anything from the standard library (such as `net/http`) is allowed.
  Libraries to support testing or types like UUID are also fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`
  , `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be
  useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.

> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider
  this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem
  necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1
  to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "
AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.
