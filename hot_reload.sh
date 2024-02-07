#!/bin/bash

inotifywait -q -m -e close_write --format %e cgw.tex |
while read events; do
	pdflatex -interaction=nonstopmode cgw
done

