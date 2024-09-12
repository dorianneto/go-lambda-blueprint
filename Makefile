cdkProfile = --profile media-metadata

build:
	@cd lambda && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bootstrap . && mv bootstrap ./bootstrap
local:
	@$(MAKE) build && $(MAKE) synth
	sam local start-api -t cdk.out/MediaMetadataStack.template.json ${cdkProfile}
bootstrap:
	cdk bootstrap ${cdkProfile}
destroy:
	cdk destroy ${cdkProfile}
synth:
	cdk synth ${cdkProfile}
diff:
	cdk diff ${cdkProfile}
deploy:
	cdk deploy ${cdkProfile}
