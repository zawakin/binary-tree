dotfiles = $(wildcard output/*.dot)
pngfiles = $(dotfiles:%.dot=%.png)

gif: image
	convert -delay 50 -gravity center -background white -extent 1500x1500 output/*.png graph.gif

image: $(pngfiles)

run:
	go run *.go

graph: run
	make image

%.png: %.dot
	dot -Tpng -Gsize=15,15\! -Gdpi=100 -Gratio=fill $< > $@

clean:
	rm output/*.png

clean-all: clean
	rm output/*.dot