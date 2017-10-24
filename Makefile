default:
	go build -o tdmg main.go

fmt:
	find main.go src/tiny-dmg -type f -exec gofmt -w {} \;

test: default
	./tdmg
