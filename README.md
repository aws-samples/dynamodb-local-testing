# Testing webservice using DynamoDB Local

This repository contains sample webservice written in Go and deployed with CDK.

It demonstrates:
* how to build a simple webservice using Go
* how to build the webservice Docker container using multistage build process
* how to deploy the webservice to AWS using CDK as a Fargate service
* how to test the webservice using DynamoDB Local

### Prerequisites
The following components need to be installed in order to build and run the code:
* Node.JS
* AWS CDK
* Go
* Docker

The code was tested using Node.JS 12.18.2, CDK 1.110.0, Go 1.16.5, Docker 20.10.7


### Webservice endpoints
* Add post `POST /post`

    Request format:
    ```json
    {
        "id": "1",
        "title": "my post #1",
        "content": "here is the post content",
        "status": "posted"
    }
    ````
* Get post by post number `GET /post/<post_number>`

## Commands

You need to run the command to install NodeJS modules: `cd cdk; npm install`

### Deploy and run the webservice on AWS

1. Deploy resources to AWS running `make deploy`
2. Get DNS name of ALB deployed and call the service using Postman or other HTTP client of your choice.
3. Remove resources running `make destroy`

### Run as a local application

Run `make run` command to run the service as a local application connected to DynamoDB in AWS (the table should be created before).

### Build Docker container

Run `make build_docker` to build Docker container.

### Run in Docker container

Run `make run_docker` to build Docker container and run locally

### Run unit test locally using DynamoDB Local

Run `make test` to run unit tests of the service using DynamoDB Local database.

### Clean resources

Run `make destroy` to undeploy AWS resources.

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This code is licensed under the MIT-0 License. See the LICENSE file.