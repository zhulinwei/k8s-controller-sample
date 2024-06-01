
generate:
	controller-gen paths=./apis/...

crd:
	controller-gen crd paths=./... output:crd:dir=config/crd
	# controller-gen crd paths=./... output:crd:dir=config/crd output:none

crd1:
	controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd

obj:
	controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."