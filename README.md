# Form3 Take Home Exercise

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## Questions

- German contract?
- Insurance (Life, Accident)?
- Confirm to be able to use own hardware?
- RSU, Stock Options or other short/long term incentives?
- Working hours? Time Tracker?
- Job Description that I'm hired for

go test -tags=integration

## Documentation

To generate the documentation first install `gomarkdoc` via:

```bash
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
```

Then execute

go test -v
go test -v -tags=int

!!!!! ENTRYPOINT vs CMD = https://stackoverflow.com/questions/21553353/what-is-the-difference-between-cmd-and-entrypoint-in-a-dockerfile
ENTRYPOINT go test -v -tags=int [... hostname]

docker run -t swaggerapi/swagger-ui
docker run -i -t golang:1.18beta1 bash

docker build . -t form3:test
docker run -t form3:test

```bash
gomarkdoc --output doc/f3.md ./
gomarkdoc --output doc/iso.md ./src/iso/...
gomarkdoc --output doc/countryCode.md ./src/iso/countryCode/...
gomarkdoc --output doc/currencyCode.md ./src/iso/currencyCode/...

client := f3.NewClient()
account := f3.NewAccount(...)
account, e := client.CreateAccount(account)

```

To read the documentation, just look [here](doc/README.md).

-- test coverage
go get golang.org/x/tools/cmd/cover
go tool cover -help

go test f3 -coverprofile=f3.out

## Errors

//
// Why do I need to URI encode reserved characters that are interpreted? https://datatracker.ietf.org/doc/html/rfc3986#section-2.2
//
// curl -v 'http://localhost:8080/v1/organisation/accounts?filter%5Bcustomer_id%5D=test' -H 'accept: application/vnd.api+json'
// curl -v 'http://localhost:8080/v1/organisation/accounts?filter[customer_id]=test' -H 'accept: application/vnd.api+json'
// curl -v 'http://localhost:8080/v1/organisation/accounts?filter[customer_id]=test%2Chello' -H 'accept: application/vnd.api+json'
// curl -v 'http://localhost:8080/v1/organisation/accounts?filter[customer_id]=test,hello' -H 'accept: application/vnd.api+json'
//



## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**
- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

### Should Nots

The finished solution **should not:**
- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected.
- Use a library for your client (e.g: go-resty). Anything from the standard library (such as `net/http`) is allowed. Libraries to support testing or types like UUID are also fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.
> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
