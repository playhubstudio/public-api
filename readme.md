# Public API

Provides a public API for the game platform and also related public docs.

## How to generate for Go

```bash
make gen-api
```

# Docs

## Readme CLI

We use readme CLI to publish docs and API specifications to readme.io
You can find and download it here: https://github.com/readmeio/rdme#readme

Register readme API key and put it into the env variable `PLAYHUB_README_API_KEY`.

Execute the following command to publish docs and API specifications to readme.io:
```bash
make publish-docs
```

## OpenAPI specifications

`api/v1/<feature>` contains OpenAPI specifications for our API.  
To generate docs from OpenAPI specifications, we use readme CLI as well.

## Docs folder

Docs folder contains markdown files for readme.io. It's structured based on the following structure:  

`/docs/v1.0/<section>/<page>.md`

where v1.0 is the version of API docs (1.0 for the moment), 
`<section>` is a section of docs (e.g. general, games, etc.) and `<page>` is a page name (e.g. authentication, games, etc.).  

*Notice:* Every doc page should have unque slug, this is how readme.io identifies pages.

## How to add a new section

To add a new section you need to create it first on readme web portal 
https://dash.readme.com/project/playhubapi/v1.0

then you need to call readme.cli to figure out the section id:
```bash
rdme sections
```

then add this information to your pages, so they will be published to the correct section.

## Where to find docs online

https://playhubapi.readme.io/reference/general-getting-started
