# Grumpo - The Grumpy Static Site Generator For Lazy People

If you are sick of overly complicated static site generators and you have a
good working knowledge of publishing websites, then Grumpo might be for you!

We aim to support a very minimal feature set hopefully keeping the maintenance
of this project low, and keeping you in control of your site.

Please be careful as this is a beta release and there may be serious bugs.
Use at your own risk!

## Quick Start

These instructions are appropriate for most developers. If you are having
trouble, please see the Setup Details section below. Everything should be
done in a shell application (terminal).

### Install Grumpo

`go get -u github.com/wrunk/grumpo`

Provided your `$GOPATH/bin` is in your path, test that grumpo is installed and
working by typing `grumpo` to see the version and command usage.
If that doesn't work ensure `$GOPATH/bin` is in your path or file an issue.

### Setup a New Project

Please create a directory specifically for this new website (let's call it `mysite`)
for the sake of this guide somewhere near where you create new software projects.

Change directory to wherever you keep your projects
`cd projects`

Make a new mysite directory
`mkdir mysite`

Change into the new directory
`cd mysite`

Now create a new Grumpo project
`grumpo init`

Before you proceed, please use git to check all these files in.
You don't need to use github at this point, just `git init`, `git add -A`
and `git commit` to make a local commit in case you want to revert later.
Version control is highly recommended given the stateless and minimal
nature of Grumpo.

Open up your favorite IDE and explore what was created.

### Running locally

`grumpo local`

Will run a local development server that will show you both pages
and drafts via an index page.

### Project Structure

TODO finish this

## Setup Details

Grumpo only works on UNIX-like operating system (so not windows). It was developed
and tested on MacOS, so your milage may vary.

Make sure:

- Your GOPATH is setup correctly
- And `$GOPATH/bin` is in your path so the grumpo command will work
- You should really have git installed and use it constantly

### Prerequisites

- You have experience with html/css
- You have at least a basic command line working knowledge

## What Grumpo Has

- SSG
- Markdown support
- Html support
- Decent error checking for pages to make sure you don't totally screw up
- A basic app engine setup that should be helpful
- Unicode support via UTF8 markdown/html files

## What Grumpo Doesn't Have

- Many features that most SSG have like a manifest/config file
- Any support for user generated content. Grumpo assumes all content
  **is trusted**. I didn't bother using Go html templates for this reason
- A fancy local dev server. It is **not** designed to work with more than
  one user at a time
- Currently no support for dynamic content like "recent posts" but
  this likely will be required in the future. Since grumpo is so easy
  to use, it should be easy to create your own recent posts manually
  for now.
- Same with reusable page fragments (see below)

## What I'm going to try very hard not to change

- The project structure

## What might change

- The grumpo command. This shouldn't be a huge deal if project
  structure stays the same

## What Grumpo Might Support Later

- Page fragments: Things like 5 recent articles that might be injected in
  multiple locations will probably need to be supported eventually
- More errors/warnings about common publishing issues or things that could
  be improved
- Some sort of better image workflow. I don't really know what this means
  yet, but generally speaking dealing with images in the online publishing world
  SUCKS - Keeping track of which image files go with which blogs - Keeping track of where you got the image (source and license if not yours) - Just making it easier to get that image you have in finder nicely
  organized in your blog project area
- Link to resources about modern web publishing like robots, seo, GA,
  GTM, ads, SEM, social, etc

## Notes On Unsupported Features

### Different Base Templates

This is probably a reasonable request, but it does feel like
it should just be another site altogether. For now not supported

### Auto Authoring, Dating, Etc

Since things are stateless and no page fragments are supported,
these dynamic sorts of things are not supported. It could be useful
to support page fragments like an "author block" but modern IDE and
shell commands make doing a bulk find and replace very easy.

Also we can leverage git dates for auto dating which might be available
soon.

## Test cases to test later

- What if a page has css or js or html in it? will markdown lib just work and ignore that?
- Spaces in files names are going to not be bueno. Probably annoying but should only support
  a-z A-Z - 0-9 .

## TODO next

- Sitemap generator
- RSS generator
- Create a robots.txt
- Also check to make sure each sub dir has an index
- Local server access logging
- 404 page support
- gcloud ignore thing for drafts
- Deploy commands (make or shell script?)
- OSS license
- When HTML fails to validate, it breaks webserver.
  Also it has an unhelpful, crappy error message
- When rss, sitemaps, robots ready, include rules in app.yaml
- Add commands to this readme for running tests, deploying

## Notes on using go modules

- Setup outside gopath
- go mod init github.com/wrunk/grumpo
- go build to add deps
- How to pull down/install?

https://blog.golang.org/using-go-modules
