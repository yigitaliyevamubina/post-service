CURRENT_DIR =$(shell pwd)

proto-gen:
	chmod +x ./scripts/gen-proto.sh
	./scripts/gen-proto.sh