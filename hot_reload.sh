#!/bin/bash

inotifywait -q -m -e close_write --format %e $1 |
while read events; do
	pdflatex -interaction=nonstopmode $1
done

