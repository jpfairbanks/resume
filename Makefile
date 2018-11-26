SRC = $(wildcard *.md)
BIBLINE=--bibliography mypubs.bib
PDFS=$(SRC:.md=.pdf)
HTML=$(SRC:.md=.html)

all:    clean $(PDFS) $(HTML)

pdf:   clean $(PDFS)
html:  clean $(HTML)

# %.html: %.md
# 	python2 resume.py html < $< | pandoc --template=./pandoc-templates/default.html -t html -c resume.css -o $@

# %.pdf:  %.md header.tex
# 	python2 resume.py tex < $< | pandoc --template=./pandoc-templates/default.latex -H header.tex -o $@

# %.tex:  %.md header.tex
# 	python2 resume.py tex < $< | pandoc --template=./pandoc-templates/default.latex -H header.tex -o $@

cv.toml: genresume.go
	go run genresume.go

cv.tex: cv.toml templates/*.tex
	sdrender -t 'templates' -o cv.tex -i cv.toml

cv.pdf: cv.tex
	pdflatex cv.tex

cv.html: cv.toml templates/*.html
	sdrender -t 'templates' -o cv.html -i cv.toml

open: cv.pdf
	open cv.pdf

clean:
	rm -f *.html *.pdf cv.tex
deliver: cv.html cv.pdf
	scp cv.pdf jpfnet:public/cv.pdf
	scp cv.html jpfnet:public/cv.html
