# Simple Template

A lazy man's templating thingy

`go get github.com/romainmenke/simple-template`

---

I needed a tool to render template html files into static files.

- it recursively parses all files in the source directory and sub-directories as templates.
- it renders the files in the source directory as pages.
- it happens to remove html comments. (did not know golang templating did that, but happy about it)

I use it with `//go:generate simple-template` and [modd](https://github.com/cortesi/modd).

---

### Options

- `-h`            : help
- `-source`       : source directory
- `-out`          : output directory
- `trailing args` : exclusion -> simple `must not contain` logic

---

Uses golang templating.

---

### Simple

- [simple-mini](https://github.com/romainmenke/simple-mini)
- [simple-bundle](https://github.com/romainmenke/simple-bundle)
- [simple-gzip](https://github.com/romainmenke/simple-gzip)
- [simple-template](https://github.com/romainmenke/simple-template)
