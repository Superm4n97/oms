push:
	@docker build -f ./auth-server/Dockerfile -t superm4n/auth-server .
	@docker push superm4n/auth-server