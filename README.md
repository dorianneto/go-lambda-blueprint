# Go-Lambda Blueprint
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/dorianneto/go-lambda-blueprint/main)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/dorianneto/go-lambda-blueprint/golangci-lint.yaml)

Allows users to spin up a quick Go lambda project using CDK

## Requirements

- CDK
- SAM
- AWS Account
- GO

### AWS Account

Once you have created the account and its access keys, you will need to set up the proper permissions in order to deploy the services accordingly via CDK.

The following permissions should be enough to accomplish that:

- `AmazonEC2ContainerRegistryFullAccess`
- `AmazonS3FullAccess`
- `AmazonSSMFullAccess`
- `AWSCloudFormationFullAccess`
- `IAMFullAccess`

### CDK

Once the CDK is installed and available, run `$ make bootstrap` in order to provision CDK base resources in AWS.

After that, you must be able to deploy it `$ make deploy`

_Note: the Makefile is set to use the profile `go-lambda-blueprint`, so remember to change it in case you create an user with a different name_

## Local development

Running `$ make local` will create a local API via SAM CLI.

_Note: you have to run that command again after every change_

## Questions?

Check the Makefile `$ make help` or open an issue. I will be happy to help you!
