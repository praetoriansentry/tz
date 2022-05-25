build: tz
tz:
	go build -o tz main.go

clean:
	$(RM) tz

install: tz
	cp tz ${HOME}/.bin
