#!/usr/bin/env sh

go run main.go
weasyprint ./public/index.html ./public/jcarr_cv.pdf
