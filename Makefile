#defaul variables
COMPOSE_PROJECT_NAME := $(shell basename "$$(pwd)")
MONGO_DATABASE ?= app

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

.PHONY: setting-up-mongodb

mongodb-seeders:
	@echo "============================================================"
	@echo "Setting up mongodb seeders on database \"${MONGO_DATABASE}\"... \n"
	@docker cp ${COMPOSE_PROJECT_NAME}-app:/src/internal/application/repository/mongo/seeders/examples.json examples.json > /dev/null 2>&1
	@docker cp examples.json ${COMPOSE_PROJECT_NAME}-mongodb:/examples.json > /dev/null 2>&1
	@echo "Cleaning previous seeds... \n"
	@docker exec -d ${COMPOSE_PROJECT_NAME}-mongodb mongosh --eval "use admin && db.auth('root', 'secret') && use ${MONGO_DATABASE} && db.examples.deleteMany({})"
	@echo "Seeding database... \n"
	@docker exec -d ${COMPOSE_PROJECT_NAME}-mongodb mongoimport \
		--uri "mongodb://root:secret@${COMPOSE_PROJECT_NAME}-mongodb:27017/${MONGO_DATABASE}" \
		--authenticationDatabase=admin \
		--collection "examples" \
		--file "examples.json" \
		--jsonArray
	@rm -rf examples.json
	@echo "Seeders done successfully.!"
	@echo "============================================================"

testing-example-service:
	@echo "============================================================"
	@echo "Creating new order... \n"
	@curl --location --request POST '127.0.0.1:8080/api/v1/order/create'
	@echo "\n============================================================"
	@echo "Searching order from seeder \n"
	@curl --location '127.0.0.1:8080/api/v1/order/656045095ff16ef1a00fd4ef'
	@echo "\n============================================================"