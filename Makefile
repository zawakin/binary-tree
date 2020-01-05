dotfiles = $(wildcard output/*.dot)
pngfiles = $(dotfiles:%.dot=%.png)

image: $(pngfiles)

graph:
	go run *.go
	make image

%.png: %.dot
	dot -Tpng $< > $@

clean:
	rm output/*.png