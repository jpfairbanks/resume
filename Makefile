SRC = $(wildcard *.md)
BIBLINE=--bibliography mypubs.bib
PDFS=$(SRC:.md=.pdf)
HTML=$(SRC:.md=.html)

all:    clean $(PDFS) $(HTML)

pdf:   clean $(PDFS)
html:  clean $(HTML)

%.html: %.md
	python2 resume.py html < $< | pandoc -t html -c resume.css -o $@

%.pdf:  %.md
	# python2 resume.py tex < $< | pandoc --template=./pandoc-templates/default.latex -H header.tex -o $@
	python2 resume.py tex < $< | pandoc -H header.tex -o $@

clean:
	rm -f *.html *.pdf
deliver: resume.pdf resume.html
	scp resume.pdf jpfnet:public/cv.pdf
	scp resume.html jpfnet:public/cv.html
	#scp resume.css jpfnet:resume.css
