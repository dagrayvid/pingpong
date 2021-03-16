cert:
	cd certs; ./gen.sh; cd ..

ca-secret:
	rm ocp_certs/*
	openssl genrsa -out ocp_certs/ca.key 2048
	openssl req -x509 -new -nodes -key ocp_certs/ca.key -subj "/CN=pingpong-ca" -days 10000 -out ocp_certs/ca.crt
	sed s"/tls.key:.*/tls.key: $$(cat ocp_certs/ca.key|base64 -w 0)/" -i manifests/01_secret.yaml
	sed s"/tls.crt:.*/tls.crt: $$(cat ocp_certs/ca.crt|base64 -w 0)/" -i manifests/01_secret.yaml

.PHONY: cert 
