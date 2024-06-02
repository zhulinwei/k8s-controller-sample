
crd:
	controller-gen crd paths=./... output:crd:dir=config/crd

install:
	kubectl apply -f config/crd
