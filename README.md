# wp-to-11ty

`wp-to-11ty` is a tool that will take a WordPress XML Export and (attempt to) convert it into a new [Eleventy](https://www.11ty.io/) static site.

This is an early tool, written for my use case, and will possibly (probably) not work for your use case.

However!  I'd love to make it better, so please [file an issue](https://github.com/jmhobbs/wp-to-11ty/issues/new) if something doesn't work for you.

# Usage

```
$ wp-to-11ty -h
usage: ./wp-to-11ty [options] <wordpress-export.xml>

options:
  -download-media
    	Download media files to local filesystem.
  -output string
    	Directory to output 11ty site to. (default "./site")
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
npm WARN blog@1.0.0 No repository field.
npm WARN blog@1.0.0 No license field.

audited 1 package in 0.765s
found 0 vulnerabilities

21:54:08 Done!
21:54:08 You can now switch to the "site" directory and run "eleventy --serve"
```

# Features

## Completed

- Post and Page export
- Optional media download
- Very basic starter layout
- Tag stubs
- Category stubs

## Planned

- Shortcode detection and warning

## Unlikely

- Custom ontolgy support
