#!/bin/bash -e

generate:
	protoc registerpb/register.proto --go_out=plugins=grpc:.