release_resource_manager:
	protoc -I=./proto_files/resource_manager  --go_out=plugins=grpc:./proto_files/resource_manager ./proto_files/resource_manager/*.proto
release_users_srv:
	protoc -I=./proto_files/users_srv --go_out=plugins=grpc:./proto_files/users_srv ./proto_files/users_srv/*.proto
release-tag:
	git tag `cat .version`
	git push --tag

