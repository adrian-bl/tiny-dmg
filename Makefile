default:
	go build -o tiny-dmg cmd/tiny-dmg.go

test: default
	./tdmg
