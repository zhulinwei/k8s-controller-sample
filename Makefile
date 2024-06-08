
crd:
	controller-gen crd paths=./... output:crd:dir=config/crd

obj:
	controller-gen object paths=./...

install:
	kubectl apply -f config/crd
