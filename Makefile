deploy:
	kubectl apply -f manifests/nginx-load-balancer.yaml
	kubectl apply -f manifests/ns.yaml
	kubectl apply -f manifests/serviceaccount.yaml
	kubectl apply -f manifests/globalizer-cr.yaml
	kubectl apply -f manifests/globalizer-crb.yaml
	kubectl apply -f manifests/deployment.yaml

delete:
	kubectl delete -f manifests/ns.yaml
	kubectl delete -f manifests/globalizer-cr.yaml
	kubectl delete -f manifests/globalizer-crb.yaml