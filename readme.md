# Public API

Provides a public API for the game platform and also related public docs.

## How to generate for Go

```bash
make gen-api
```

## How to version

We use semver versioning for our API `major.minor.patch`.  
When we make breaking changes, we increase the major version.  
When we add new features in a backwards-compatible manner, we increase the minor version.  
When we make backwards-compatible bug fixes, we increase the patch version.

When we crete a new version we need to create new folder and copy the content of the previous version to the new one.
For both API and docs.  
Then we need to add generator for the version in the `Makefile`.

Once creating a new version, we need to create it first on readme web portal, then use this section id in the `Makefile`.
https://dash.readme.com/project/playhubapi/v1.1/versions

## How to reupload (or change categories)

First go to the web portal and delete existing categories, then change `Makefile` and remove categori ID and `--update` flag.
Once uploaded (choose create new category), get the category ID and put it back to the `Makefile` along with `--update` flag.
You need to change all docs accordingly to the new category ID.

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

`api/vx.y/<feature>` contains OpenAPI specifications for our API.  
To generate docs from OpenAPI specifications, we use readme CLI as well.

## Docs folder

Docs folder contains markdown files for readme.io. It's structured based on the following structure:

`/docs/vx.y/<section>/<page>.md`

where v1.0 is the version of API docs (1.0 for the moment),
`<section>` is a section of docs (e.g. general, games, etc.) and `<page>` is a page name (e.g. authentication, games, etc.).

_Notice:_ Every doc page should have unque slug, this is how readme.io identifies pages.

## How to version

We use semver versioning for our API `major.minor.patch`.  
When we make breaking changes, we increase the major version.  
When we add new features in a backwards-compatible manner, we increase the minor version.  
When we make backwards-compatible bug fixes, we increase the patch version.

When we crete a new version we need to create new folder and copy the content of the previous version to the new one.
Then we need to add generator for the version in the `Makefile`.

all docs should be fixed according to the new version (change all IDs).

```base
rdme categories --version x.y
```

## How to add a new section

To add a new section you need to create it first on readme web portal
https://dash.readme.com/project/playhubapi/v1.0

then you need to call readme.cli to figure out the section id:

execute

```base
rdme categories --version x.y
```

and get new IDs from there.

then add this information to your pages, so they will be published to the correct section.

## Where to find docs online

https://playhubapi.readme.io/reference/general-getting-started
