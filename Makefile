graph:
	go run *.go
	dot -Tpng output/graph.dot > output/graph.png

clean:
	rm -r output/