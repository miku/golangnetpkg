all: server.key server.crt

server.key:
	openssl genrsa -out server.key 3072

server.crt:
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

clean:
	rm -f server.key server.crt

mkcert:
	mkcert example.com "*.example.com" example.test localhost 127.0.0.1 ::1
