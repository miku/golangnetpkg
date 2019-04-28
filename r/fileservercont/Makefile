.PHONY: clean install image run

webshare: webshare.go
	CGO_ENABLED=0 go build -o $@ $<

clean:
	rm -f webshare
	rm -f webshare.1
	docker rmi -f gosp/webshare

webshare.1: webshare.md
	pandoc $< -s -t man > $@

PREFIX = $(HOME)
BINDIR = $(PREFIX)/bin
MANDIR = $(PREFIX)/man/man1

install: webshare webshare.1
	install -d $(BINDIR) $(MANDIR)
	install webshare $(BINDIR)
	install webshare.1 $(MANDIR)

image:
	# About 7.53MB.
	docker build -t gosp/webshare .

run:
	DOCKER_CONTENT_TRUST=0 docker run -v `pwd`:/share --rm -it -p 3000:3000 gosp/webshare

