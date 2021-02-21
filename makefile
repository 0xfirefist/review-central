graphql-gen:
	cd graphql-server; go run github.com/99designs/gqlgen generate; cd ..;

run-local-server:
	go run main.go;


## UI
build-ui: 
	cd ui;npm run build;cd ..;

dev-ui:
	cd ui;npm run start;cd ..;