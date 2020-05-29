# Grumpo - The Grumpy Static Site Generator For Lazy People

If you are sick of overly complicated static site generators and you have a
good working knowledge of publishing websites, then Grumpo might be for you!

We aim to support a very minimal feature set hopefully keeping the maintenance
of this project low, and keeping you in control of your site.

Please be careful as this is a beta release and there may be serious bugs.
Use at your own risk!

Grumpo is free and open source software and is published with no
warranties of any kind! Please see [License](/LICENSE) for details.

## Quick Start

These instructions are appropriate for most developers. If you are having
trouble, please see the Setup Details section below. Everything should be
done in a shell application (terminal).

### Install Grumpo

`go get -u github.com/wrunk/grumpo`

Provided your `$GOPATH/bin` (`~/go/bin` with Go modules) is in your path, test that
grumpo is installed and
working by typing `grumpo` to see the version and command usage.
If that doesn't work ensure `$GOPATH/bin` is in your path or file an issue.

If you don't know what GOPATH is, make sure you include `~/go/bin` in your
PATH

### Setup a New Project

Please create a directory specifically for this new website (let's call it `mysite`)
for the sake of this guide somewhere near where you create new software projects.
It might be better to do this outside of your gopath

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
Constant use of version control is highly recommended given the stateless
and minimal nature of Grumpo.

Open up your favorite IDE and explore what was created.

### Running locally

`grumpo local`

Will run a local development server to test your pages. This command
will give you a special index to view.

## Project Structure And Content Creation

### Where to create content?

If you notice on the `__index__` page there are a few pages with their
metadata at the top (see metadata details below).

Hopefully examining the initial pages should give you a sense
of how this all works.

In the static site generator world, a local page such as `pages/about.md`
would become `/about/` on the server and should be linked to like that.
It is important to include the trailing `/` because that will tell the
web server to load `/about/index.html`

## Metadata

Originally I was trying to avoid having this section, but for even a
semi-complicated site it's very necessary.

### Format

The metadata section is just a simple json object `{...}`. You can add single
line, c-style comments `//` so long as they are on their own line and within
the object.

### Full Example

Most fields besides live and title can just be skipped if not desired/needed.

```json
{
  // You can use c-style comments as long as they are on their own
  // line and use the // format (not /**/)

  // Title is ALWAYS required. You can control how this is used in
  // base.html
  "title": "Great Blog Post Aboot Cats!",

  // Used in OG tags and can be used in recent posts promo
  "desc": "Here is a short(ish) description about this article",

  // Headlines can be used in various ways. More on this later
  "hl1": "Headlines could be variations on title for",
  "hl2": "a/b testing purposes or",
  "hl3": "auto tweeting so you post your content multiple times spread out",

  // Specify a relative or full URL to the article's canonical image.
  // We'll set article's og:image tag to this
  "image": "/static/img/hello.jpg",
  "image_alt": "A hello face",

  // If set to true, this page will go live with grumpo gen
  // However it wont show up in recent posts until you set
  // a publish date
  "live": false,

  // Both support the format 2020-01-01 OR 2020-01-01:03:04:05
  // No other formats are supported and no timezones can be passed in
  // grumpo commands use your machine's local time zone
  "publish_date": "",
  "updated_date": "",

  // Defaults to false. Set to true to treat this page as a full html page
  "skip_base_template": false
}
```

## Generating and Deploying

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

## Static Files

All static files like js, css, fonts, and images should go in the static directory.
During a `grumpo gen` they will get copied to `build/static/...` directory

Usually you want to keep the different types in their own directory like
`static/js/` `static/img/` etc.

## Still Confused?

Consider reading the code (there's not that much of it), or checking out Grumpo
from Github and making some changes. You can mess around then run `go install`
and it will overwrite your `grumpo` command.

Of course you can also file an issue on Github and I will try to help.

## Setup Details

Grumpo only works on UNIX-like operating system (so not windows). It was developed
and tested on MacOS, so your milage may vary.

Make sure (assuming Go Modules):

- `~/go/bin` is in your path (so `go get` or `go install` will make the
  grumpo command available)
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

- Any support for user generated content. Grumpo assumes all content
  **is trusted**. I didn't bother using Go html templates for this reason
- A fancy local dev server. It is **not** designed to work with more than
  one user at a time and does not take any security precautions

## What I'm going to try very hard not to change

- The project structure

## What might change

- The grumpo command. This shouldn't be a huge deal if project
  structure stays the same. Likely grumpo command will need to
  improve to add needed features and fix bugs

## What Grumpo Might Support Later

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
  GTM, ads, SEM, social, OG tags, etc

## Running Software/Automated Tests

To run what few tests there are, simply `go test -v`

## Notes On Unsupported Features

### Different Base Templates

This is probably a reasonable request, but it does feel like
it should just be another site altogether. For now not supported
although I am working on a setting to skip the base template
so you could make full pages within

## Test cases to test later

- What if a page has css or js or html in it? will markdown lib just work and ignore that?
- Spaces in files names are going to not be bueno. Probably annoying but should only support
  a-z A-Z - 0-9 .

## TODO next

- Need to update files in init.go to reflect new json config process
- Sitemap generator
- RSS generator
- Create a robots.txt
- Also check to make sure each sub dir has an index
- Local server access logging
- 404 page support
- When HTML fails to validate, it breaks webserver.
  Also it has an unhelpful, crappy error message
- When rss, sitemaps, robots ready, include rules in app.yaml
- Home page should list out all pages/ to show how to use dynamic content
- Consider adding or at least commenting on anchorjs or how this
  could be done with markdown generator. https://github.com/bryanbraun/anchorjs
- Same with highlight.js
- Or perhaps use the server side highlighting
- Print an error if port 9876 is in use

## Notes on using go modules

These are just my personal notes for now and should move elsewhere soon.

- Setup outside gopath
- go mod init github.com/wrunk/grumpo
- go build to add deps
- How to pull down/install?

https://blog.golang.org/using-go-modules
