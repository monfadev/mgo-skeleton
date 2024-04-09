# go test -cover
# go test -coverprofile=cover.txt
# go tool cover -func=cover.txt
# go tool cover -html=cover.txt


go test ./...
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

# sonar-scanner \
#   -Dsonar.projectKey=smart-poultry-service-6 \
#   -Dsonar.sources=. \
#   -Dsonar.host.url=http://localhost:9000 \
#   -Dsonar.token=sqp_c123aa1268d93888d9ec44f2254f2d67583bba5b
