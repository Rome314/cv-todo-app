
test:
	./scripts/test ${ARGS}

twirp:
	./scripts/twirp.sh


migrate-up:
	migrate -path migrations -database "postgresql://postgres:simplepassword@localhost:5432/postgres?sslmode=disable&search_path=public" -verbose  up ${v}

migrate-down:
	migrate -path migrations -database "postgresql://postgres:simplepassword@localhost:5432/postgres?sslmode=disable&search_path=public" -verbose  down ${v}
migrate-force:
	migrate -path migrations -database "postgresql://postgres:simplepassword@localhost:5432/postgres?sslmode=disable&search_path=public" -verbose  force ${v}

migrate-drop:
	migrate -path migrations -database "postgresql://postgres:simplepassword@localhost:5432/postgres?sslmode=disable&search_path=public" -verbose  drop