.PHONY: setup-pulumi
setup-pulumi:
	pulumi stack init $(PULUMI_STACK)
	cp ./Pulumi.config.yaml.template "./Pulumi.$(PULUMI_STACK).yaml"

.PHONY: deploy
deploy:
	pulumi up

.PHONY: destroy
destroy:
	pulumi destroy
