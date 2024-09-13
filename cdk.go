package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HelloWorldStackProps struct {
	awscdk.StackProps
}

func HelloWorldStack(scope constructs.Construct, id string, props *HelloWorldStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaFn := awslambda.NewFunction(stack, jsii.String("helloWorldLambdaFn"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.Code_FromAsset(jsii.String("./lambda"), nil),
		Handler: jsii.String("bootstrap"),
	})

	apiGateway := awsapigateway.NewRestApi(stack, jsii.String("helloWorldApiGateway"), nil)
	integration := awsapigateway.NewLambdaIntegration(lambdaFn, nil)

	helloWorldResource := apiGateway.Root().AddResource(jsii.String("hello-world"), nil)
	helloWorldResource.AddMethod(jsii.String("GET"), integration, nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	HelloWorldStack(app, "HelloWorldStack", &HelloWorldStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
