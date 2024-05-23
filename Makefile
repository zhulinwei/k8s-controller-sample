
generate:
	controller-gen paths=./apis/...

crd:
	controller-gen crd paths=./... output:crd:dir=config/crd output:stdout

crd1:
	controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd

obj:
	controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."