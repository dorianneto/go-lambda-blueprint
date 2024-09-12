package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MediaMetadataStackProps struct {
	awscdk.StackProps
}

func MediaMetadataStack(scope constructs.Construct, id string, props *MediaMetadataStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaFn := awslambda.NewFunction(stack, jsii.String("mediaMetadataLambdaFn"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.Code_FromAsset(jsii.String("./lambda"), nil),
		Handler: jsii.String("bootstrap"),
	})

	apiGateway := awsapigateway.NewRestApi(stack, jsii.String("mediaMetadataApiGateway"), nil)
	integration := awsapigateway.NewLambdaIntegration(lambdaFn, nil)

	metadata := apiGateway.Root().AddResource(jsii.String("metadata"), nil)
	metadata.AddMethod(jsii.String("POST"), integration, nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	MediaMetadataStack(app, "MediaMetadataStack", &MediaMetadataStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
