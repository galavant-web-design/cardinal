# Cardinal

![cardinal](sample/images/cardinal.png)
A simple site builder

## Install

```shell
go install github.com/tygern/cardinal@v0.0.7
```

## Start the generator

Cardinal builds a site from the files in the current directory and starts a webserver to serve the files locally.
Cardinal rebuilds the site whenever the source files change.

```shell
cardinal
```

To build the files once without starting the server run

```shell
cardinal -build
```

## Project layout

```
.
├── template.html
├── index.html
├── about.html
├── style
│   └── app.css
└── images
    └── cardinal.png
```

- The `template.html` file defines your site template.
  Cardinal replaces the `<#content/>` tag with site content.
- Any other HTML files define site content.
  Cardinal applies the template to each html file and creates a file with the same name in the `build` directory.
- Cardinal also copies all non-HTML files and directories to the `build` directory. 

## Local development

### Build

```shell
go build
```

### Run

```shell
cd sample
go run ..
```
