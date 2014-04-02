SRC = $(wildcard *.md)
BIBLINE=--bibliography mypubs.bib
PDFS=$(SRC:.md=.pdf)
HTML=$(SRC:.md=.html)

all:    clean $(PDFS) $(HTML)

pdf:   clean $(PDFS)
html:  clean $(HTML)

%.html: %.md
	python resume.py html < $< | pandoc -t html -c resume.css -o $@

%.pdf:  %.md
	python resume.py tex < $< | pandoc --template=./pandoc-templates/default.latex -H header.tex -o $@

clean:
	rm -f *.html *.pdf
deliver: resume.pdf resume.html
	scp resume.pdf jpfnet:cv.pdf
	scp resume.html jpfnet:cv.html
	#scp resume.css jpfnet:resume.css
