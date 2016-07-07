net.a:
	go build -buildmode=c-archive -o net.a net.go

http-server:net.a
	gcc -pthread -o http-server http_server.c net.a


run: http-server
	./http-server

clean:
	rm http-server net.a
