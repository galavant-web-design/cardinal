# Cardinal

![cardinal](sample/images/cardinal.png)
A simple site builder

## Install

```shell
go install github.com/tygern/cardinal
```

## Build files

```shell
cardinal
```

## Build and serve files

```shell
cardinal -serve
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
  When Cardinal is run it applies the template to each html file and creates a file with the same name in the `build`
  directory.
- Cardinal copies all other files and directories to the `build` directory. 

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
