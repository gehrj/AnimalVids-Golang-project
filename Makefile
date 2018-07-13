GCC=go
GCMD=run
GPATH=main.go

run:
	$(GCC) $(GCMD) $(GPATH)
build:
	make build_db

build_db:
	rm pkg/db/db_structs.go
	go run pkg/db/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/

install:
	make install_routes
	make install_db

install_routes:
	go get -u github.com/gorilla/mux
install_db:
	go get -u github.com/go-xorm/xorm