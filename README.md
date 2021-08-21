# HTML/PDF CV

A small program I wrote to generate a decent-looking (imo) CV/résumé in 
HTML using Go's templating engine. It also produces a PDF version from 
the HTML using the WeasyPrint Python library.

Data are all defined in `data.yml`.  Running `build.sh` should build the
HTML CV from the HTML templates in `./templates/` and the data contained
in `data.yml` with Go, then create the PDF with Python/WeasyPrint. The
rendered HTML and PDF go in `./public/`.

--- 

### Dependencies
- Go 1.15
  - yaml v3
    - `go get gopkg.in/yaml.v3`
- Python 3.9 (required to generate PDF)
  - WeasyPrint version 53.0
    - `python -m pip install weasyprint`

---


![Link to the PDF here.](./public/jcarr_cv.pdf)
