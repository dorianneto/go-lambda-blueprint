awsProfile = --profile go-lambda-blueprint

build:
	GOOS=linux GOARCH=amd64 go build -C ./lambda -ldflags="-s -w" -o bootstrap .
local:
	@$(MAKE) build && $(MAKE) synth
	sam local start-api -t ./cdk.out/HelloWorldStack.template.json ${awsProfile}
bootstrap:
	cdk bootstrap ${awsProfile}
destroy:
	cdk destroy ${awsProfile}
synth:
	cdk synth ${awsProfile}
diff:
	cdk diff ${awsProfile}
deploy:
	$(MAKE) build && $(MAKE) synth && cdk deploy ${awsProfile}
