# wp-to-11ty

`wp-to-11ty` is a tool that will take a WordPress XML Export and (attempt to) convert it into a new [Eleventy](https://www.11ty.io/) static site.

This is an early tool, written for my use case, and will possibly (probably) not work perfectly for your use case, but that's ok!

I'd love to make it better, and make it more flexible, but I only have the one blog to export and test it with, so please [file an issue](https://github.com/jmhobbs/wp-to-11ty/issues/new) if something doesn't work for you.

There's a little bit more details on the overall process available [on my blog](https://velvetcache.org/2023/04/05/moving-from-wordpress-to-11ty/)

# Installation

## Go Install

You can install `wp-to-11ty` with go.  It currently requires version >= 1.20

```
$ go install github.com/jmhobbs/wp-to-11ty@latest
go: downloading github.com/jmhobbs/wp-to-11ty v0.0.1
```

## Release Binaries

Additionally, you can use the [Releases](https://github.com/jmhobbs/wp-to-11ty/releases) binaries if they are available for your OS and architecture.

# Usage

To use `wp-to-11ty` you will need to [export your WordPress site](https://wordpress.com/support/export/#export-content-to-another-word-press-site) and unzip that to get the XML file.  This file will be provided to `wp-to-11ty` for conversion.

Additionally, be aware that `wp-to-11ty` current does nothing with draft posts, it will simply skip over them.

```
$ wp-to-11ty -h
usage: wp-to-11ty [options] <wordpress-export.xml>

options:
  -download-media
    	Download media files to local filesystem. (default false)
  -output string
    	Directory to output 11ty site to. (default "./site")
```

## Flags

### `-download-media`

By default, `wp-to-11ty` will _not_ download your uploads from the export. If you wish to download all your uploads (i.e. `wp-content`), pass this flag.

### `-output <string>`

By default `wp-to-11ty` will put all your new 11ty files into a new directory called `site`.  If you need this to be somewhere else, you can specify that path here.

## Example

```
$ wp-to-11ty -download-media blog-2019-11-16.xml
21:54:06 Writing scaffolding files...
21:54:06 Writing out pages and posts...
21:54:06 Skipping draft: Never Luke / Nikolai Archives
21:54:06 Skipping draft: Backup your WordPress to S3 with backwpup
21:54:06 Skipping draft: Method to select an interesting single frame from a gif
21:54:06 Skipping draft: Parking Assistant
21:54:06 Skipping draft: Netlify + Cloudflare = Crazy Delicious
21:54:07 Downloading attachments...
21:54:07   - http://www.velvetcache.org/wp-content/uploads/baby_hobbs.jpg
21:54:07     File exists: site/wp-content/uploads/baby_hobbs.jpg
21:54:07 Installing npm packages...

added 210 packages, and audited 211 packages in 721ms

38 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
21:54:08 Done!
21:54:08 You can now switch to the "../site" directory and run "npm run serve"

```

# Features

## Completed

- Post and Page export
- Optional media download
- Very basic starter layout
- Tag listings
- Category listings
- Year archives
- Comment export to JSON

# Anti-features

- Custom ontolgy support
- Nested categories

# Output

WordPress and 11ty have different opinions about how to manage content, and this tool attempts to squeeze the WordPress model into something that 11ty can consume in a useful way.  The primary means of doing this is through the use of front matter and custom collections.

The guiding principle of all of this is attempting to keep your URL structure as close to what WordPress had as possible, because [cool URI's don't change](https://www.w3.org/Provider/Style/URI).

If you'd like to see the output, after my own extended mangling, you can view the git repository for my blog: [jmhobbs/velvetcache.org](https://github.com/jmhobbs/velvetcache.org)

## Custom Front Matter

There are a few custom front matter fields we use.

- `type` is used to indicate the PostType, generally these are `post` or `page`
- `wp_id` is the original WordPress ID for the content, useful to tie back comments
- `creator` is the username of the WordPress author that created the entry
- `category` is the mapping of WordPress categories. It was kept separate from tags intentionally.

## Tags & Categories

Tags and Categories are [very similar tools](https://wordpress.com/support/posts/categories-vs-tags/).  I've chosen to embrace tags as the first class organizational tool, and have assigned them to the 11ty `tags` field for use in [collections](https://www.11ty.dev/docs/collections/)

Categories are still tracked using the `category` front matter.

Neither tags nor categories are normalized in any way.  There are two data files exported, `categories.json` and `post_tags.json`, which help map them cleanly from their raw form to the correct nice format for URL's.  The filters `wp_tag_slug` and `wp_category_slug` can be used to safely and consistently slugify these.

## Examples

Here are some example front matter taken from exporting my blog.

### Page

```
---
creator: admin
date: 2012-02-05T20:17:47
layout: page.njk
permalink: /contact/
tags: []
title: Contact
type: page
wp_id: "2164"
---
```

### Post

```
---
category:
- Geek
creator: admin
date: 2019-01-09T23:36:01
layout: layout.njk
permalink: /2019/01/09/easy-visual-identification-of-git-short-shas/
tags:
- git
- JavaScript
title: Easy visual identification of git short sha's.
type: post
wp_id: "2845"
---
```

## Comments

Comments (and pingbacks) are exported to a data file, `comments.json`, for your use (or not).  They are tied to the `wp_id` value in the front matter of each entry.

# Alternatives

There are alternatives to this project, which may be better suited for your use case.

- [wordpress-export-to-markdown](https://github.com/lonekorean/wordpress-export-to-markdown)
- [wordpress-to-jekyll-exporter](https://github.com/benbalter/wordpress-to-jekyll-exporter)
