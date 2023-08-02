.PHONY: shipments-snapshots-service
ss:
	@echo "Compiling shipments-snapshots-service... \c"
	@protoc -Iinternal/adapters/driven/grpc/protos/core/services/ \
		-Iinternal/adapters/driven/grpc/protos/core/models/ \
		--go_out=internal/adapters/driven/grpc/ss/ \
	 	--go_opt=paths=source_relative \
	 	--go-grpc_out=internal/adapters/driven/grpc/ss/ \
	 	--go-grpc_opt=paths=source_relative \
		--experimental_allow_proto3_optional \
	 	internal/adapters/driven/grpc/protos/core/services/shipments-snapshots.proto
	@echo "Done."