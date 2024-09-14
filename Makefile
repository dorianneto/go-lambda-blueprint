## ----------------------------------------------
## Pick one of the following commands below:
## ----------------------------------------------

awsProfile = --profile go-lambda-blueprint

help: ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
build: ## build compiles the packages
	GOOS=linux GOARCH=amd64 go build -C ./lambda -ldflags="-s -w" -o bootstrap .
local: ## runs lambda locally
	@$(MAKE) build && $(MAKE) synth
	sam local start-api -t ./cdk.out/HelloWorldStack.template.json ${awsProfile}
bootstrap: ## provision stack
	cdk bootstrap ${awsProfile}
destroy: # destroy stack
	cdk destroy ${awsProfile}
synth: ## synth stack
	cdk synth ${awsProfile}
diff: ## diff stack
	cdk diff ${awsProfile}
deploy: ## deploy stack
	$(MAKE) build && $(MAKE) synth && cdk deploy ${awsProfile}
