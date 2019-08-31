
generate:
	protoc api/register.proto --go_out=plugins=grpc:.
	protoc api/blog.proto --go_out=plugins=grpc:.
clean:
		rm -rf vendor/
