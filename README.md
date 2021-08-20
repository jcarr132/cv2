# CV Generator

## Dependencies
```
- Go 1.15
  - yaml v3.0.0
- Python 3.9 (required for PDF)
  - WeasyPrint version 53.0
  -
```

Data are all defined in `data.yml`.  Running `build.sh` should build the
HTML resume from the HTML templates in `./templates/` and the data
contained in `data.yml`. The rendered HTML and PDF go in `./public/`.
