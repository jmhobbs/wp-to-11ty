package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var (
		outputDirectory *string = flag.String("output", "./site", "Directory to output 11ty site to.")
	)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] <wordpress-export.xml>", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if "" == flag.Arg(0) {
		fmt.Fprintln(flag.CommandLine.Output(), "error: WordPress export XML is required.")
		flag.Usage()
		os.Exit(1)
	}

	*outputDirectory = filepath.Clean(*outputDirectory)

	if err := os.MkdirAll(*outputDirectory, 0700); err != nil {
		log.Fatalf("unable to create output directory %q: %v", *outputDirectory, err)
	}

	writeConfigFile(*outputDirectory)

	export, err := readXML(flag.Arg(0))
	if err != nil {
		log.Fatalf("error reading export xml: %v", err)
	}

	fmt.Printf("Blog Name: %s\n", export.Channel.Title)
	fmt.Printf(" Base URL: %s\n", export.Channel.BaseSiteURL)

	fmt.Println("== Writing Pages")
	for _, item := range export.Channel.Items {
		if item.PostType == "page" {
			fmt.Println(item.Title)
			err = writeOutPage(*outputDirectory, item)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func writeConfigFile(base string) {
	f, err := os.Create(filepath.Join(base, ".eleventy.js"))
	if err != nil {
		log.Fatalf("unable to create eleventy config file: %v", err)
	}
	defer f.Close()
	f.WriteString(configJs)
}

func writeOutPage(base string, item Item) error {
	u, err := url.Parse(item.Link)
	if err != nil {
		return err
	}

	dirs, file := pathToFileSystem(u.Path)
	dirs = append([]string{base}, dirs...)
	dirPath := filepath.Join(dirs...)

	if err = os.MkdirAll(dirPath, 0700); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dirPath, fmt.Sprintf("%s.html", file)))
	if err != nil {
		return err
	}
	defer f.Close()

	postDate, err := time.Parse(POST_DATE_FORMAT, item.PostDateGmt)
	if err != nil {
		return err
	}

	// front matter
	f.WriteString("---\n")
	fmt.Fprintf(f, "title: %s\n", item.Title)
	fmt.Fprintf(f, "permalink: %s\n", u.Path) // TODO: ?
	fmt.Fprintf(f, "date: %s\n", postDate.Format(OUTPUT_DATE_FORMAT))
	fmt.Fprintf(f, "tags: \n") // TODO
	f.WriteString("---\n")

	for _, el := range item.Contents {
		if el.XMLName.Space == CONTENT_NS {
			f.WriteString(el.Data)
		}
	}

	return err
}

func pathToFileSystem(urlPath string) ([]string, string) {
	split := strings.Split(urlPath, "/")
	for i, v := range split {
		if v == "" {
			split = append(split[:i], split[i+1:]...)
		}
	}

	return split[:len(split)-1], split[len(split)-1]
}

func readXML(path string) (*BlogExport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var export BlogExport

	dec := xml.NewDecoder(f)

	err = dec.Decode(&export)

	return &export, err
}

const (
	POST_DATE_FORMAT   string = "2006-01-02 15:04:05"
	OUTPUT_DATE_FORMAT string = "2006-01-02T15:04:05"
)
