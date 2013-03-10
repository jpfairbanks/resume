SRC = $(wildcard *.md)

PDFS=$(SRC:.md=.pdf)
HTML=$(SRC:.md=.html)

all:    clean $(PDFS) $(HTML)

pdf:   clean $(PDFS)
html:  clean $(HTML)

%.html: %.md
	python resume.py html < $< | pandoc -t html --bibliography mypubs.bib -c resume.css -o $@

%.pdf:  %.md
	python resume.py tex < $< | pandoc --template=./pandoc-templates/default.latex -H header.tex --bibliography mypubs.bib -o $@

clean:
	rm -f *.html *.pdf
