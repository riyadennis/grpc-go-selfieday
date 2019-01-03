
generate:
	protoc api/register.proto --go_out=plugins=grpc:.
clean:
		rm -rf vendor/
