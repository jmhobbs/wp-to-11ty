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
$
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

# Alternatives

There are alternatives to this project, which may be better suited for your use case.

- [wordpress-export-to-markdown](https://github.com/lonekorean/wordpress-export-to-markdown)
- [wordpress-to-jekyll-exporter](https://github.com/benbalter/wordpress-to-jekyll-exporter)
