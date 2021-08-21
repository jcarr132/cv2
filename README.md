# CV Generator

## Dependencies

```
- Go 1.15
  - yaml v3.0.0
- Python 3.9 (required to generate PDF)
  - WeasyPrint version 53.0
```

Data are all defined in `data.yml`.  Running `build.sh` should build the
HTML CV from the HTML templates in `templates/` and the data contained
in `data.yml` with Go, then create the PDF with Python/WeasyPrint. The
rendered HTML and PDF go in `public/`.
