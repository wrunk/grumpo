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

If you don't know what GOPATH is, make sure you include `~/go/bin` in your
PATH

### Setup a New Project

Please create a directory specifically for this new website (let's call it `mysite`)
for the sake of this guide somewhere near where you create new software projects.
It might be better to do this outside of your gopaht

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
and drafts via an index page (load the page the terminal tells you to).

## Project Structure And Content Creation

### Where to create content?

If you notice on the `__index__` page there are drafts and pages.
The pages/ dir is for content that is finished and should go live.
The drafts/ dir is for stuff in progress and can viewed via the local
development server.

Hopefully examining the initial pages and drafts should give you a sense
of how this all works.

In the static site generator world, a local page such as `pages/about.md`
would become `/about/` on the server and should be linked to like that.
It is important to include the trailing `/` because that will tell the
web server to load `/about/index.html`

### Generating and Deploying

Once you have created some pages and are ready to push them to a server,
do the following:

Run the Grumpo generation process:

`grumpo gen`

This will create a build/ directory with your entire site contained within.
This could be uploaded to S3 or similar.

If you want to try using Google App Engine for free to host this, create a
google cloud project and install the SDK then run:

`gcloud app deploy --project=<project-name> ./app.yaml`

Where you are replacing `<project-name>` with your GCP project name.

During the deploy you should see a note about a target url:

`target url: [https://grumpo-dot-my-project.appspot.com]`

Which once the deploy finishes should host your app.

### Static Files

All static files like js, css, fonts, and images should go in the static directory.
During a `grumpo gen` they will get copied to `build/static/...` directory

Usually you want to keep the different types in their own directory like
`static/js/` `static/img/` etc.

### Still Confused?

Consider reading the code (there's not that much of it), or checking out Grumpo
from Github and making some changes. You can mess around then run `go install`
and it will overwrite your `grumpo` command.

Of course you can also file an issue on Github and I will try to help.

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

## What Grumpo Is and Has

- A simple static site generator
- With Markdown support
- And Html support (can use md or html files)
- Decent error checking for pages to make sure you don't totally screw up
- A minimal local development to make life easier
- A basic app engine setup that should let you deploy your app for free
- Unicode support via UTF8 markdown/html files

## What Grumpo Doesn't Have

- Many features that most SSG have like a manifest/config file
- Any support for user generated content. Grumpo assumes all content
  **is trusted**. I didn't bother using Go html templates for this reason
- A fancy local dev server. It is **not** designed to work with more than
  one user at a time and does not take any security precautions
- Currently no support for dynamic content like "recent posts" but
  this likely will be required in the future. Since grumpo is so easy
  to use, it should be easy to create your own recent posts manually
  for now.
- Same with reusable page fragments (see below)

## What I'm going to try very hard not to change

- The project structure

## What might change

- The grumpo command. This shouldn't be a huge deal if project
  structure stays the same. Likely grumpo command will need to
  improve to add needed features and fix bugs

## What Grumpo Might Support Later

- Page fragments: Things like 5 recent articles that might be injected in
  multiple locations will probably need to be supported eventually
- More errors/warnings about common publishing issues or things that could
  be improved
- Some sort of better image workflow. I don't really know what this means
  yet, but generally speaking dealing with images in the online publishing world
  SUCKS
  - Keeping track of which image files go with which blogs
  - Keeping track of where you got the image (source and license if not yours)
  - Just making it easier to get that image you have in finder nicely organized
    in your blog project area
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
- Link to a simple GOPATH explanation for setup

## Notes on using go modules

These are just my personal notes for now and should move elsewhere soon.

- Setup outside gopath
- go mod init github.com/wrunk/grumpo
- go build to add deps
- How to pull down/install?

https://blog.golang.org/using-go-modules
