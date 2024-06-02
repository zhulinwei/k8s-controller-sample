
crd:
	controller-gen crd paths=./... output:crd:dir=config/crd

obj:
	controller-gen object paths=./api/rollout_sample/v1

install:
	kubectl apply -f config/crd
