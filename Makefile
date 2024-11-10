push-auth-server:
	@docker build -f ./auth-server/Dockerfile -t superm4n/auth-server .
	@docker push superm4n/auth-server

push-oms-server:
	@docker build -f ./oms/Dockerfile -t superm4n/order-server .
	@docker push superm4n/order-server

push-order-processor:
	@docker build -f ./order-processor/Dockerfile -t superm4n/order-processor .
	@docker push superm4n/order-processor

push-all: push-auth-server push-oms-server push-order-processor