clean:
	rm -rf gen

swagger: clean
	mkdir gen
	swagger -q generate server -t gen -f swagger.yaml --exclude-main -A todo