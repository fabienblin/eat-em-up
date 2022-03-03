
test: ./dealer ./eagles.gang ./cobras.gang
	./dealer -map "neighbourhoods/00.hood" -g1="eagles.gang" -g2="cobras.gang"

all:
	go build -o "dealer" ./hood/src/main.go
	go build -o "cobras.gang" ./cobra\ gang/src/main.go
	go build -o "eagles.gang" ./eagle\ gang/src/main.go

clean: ./dealer ./eagles.gang ./cobras.gang
	rm ./dealer
	rm ./cobras.gang
	rm ./eagles.gang