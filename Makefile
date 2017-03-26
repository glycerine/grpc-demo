all:
	cd streambigfile && make
	cd client && go build
	cd server && go build

run:
	./server/server &
	sleep 1 # let server start
	./client/client

clean:
	rm -f server/server
	rm -f client/client
	find . -name '*~' | xargs rm -f
