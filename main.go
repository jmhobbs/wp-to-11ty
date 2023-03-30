package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v2"
)

const (
	POST_DATE_FORMAT   string = "2006-01-02 15:04:05"
	OUTPUT_DATE_FORMAT string = "2006-01-02T15:04:05"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	var (
		outputDirectory *string = flag.String("output", "./site", "Directory to output 11ty site to.")
		downloadMedia   *bool   = flag.Bool("download-media", false, "Download media files to local filesystem.")
	)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [options] <wordpress-export.xml>\n\noptions:\n", os.Args[0])
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

	export, err := readXML(flag.Arg(0))
	if err != nil {
		log.Fatalf("error reading export xml: %v", err)
	}

	log.Println("Writing scaffolding files...")
	if err = writeStaticFiles(*outputDirectory, export); err != nil {
		log.Fatal(err)
	}
	if err = writeGlobalDataFiles(*outputDirectory, export); err != nil {
		log.Fatal(err)
	}

	log.Println("Writing out pages and posts...")
	writePagesAndPosts(*outputDirectory, export)

	if *downloadMedia {
		log.Println("Downloading attachments...")
		downloadAttachments(*outputDirectory, export)
	}

	log.Println("Installing npm packages...")
	if err = npmInstall(*outputDirectory); err != nil {
		log.Println(err)
	}

	log.Println("Done!")
	log.Printf(`You can now switch to the "%s" directory and run "%s"`, color.CyanString(*outputDirectory), color.CyanString("npm run serve"))
}

func writePagesAndPosts(base string, export *BlogExport) {
	for _, item := range export.Channel.Items {
		if item.PostType == "page" || item.PostType == "post" {
			if item.Status == "draft" {
				log.Println(color.YellowString("Skipping draft:"), item.Title)
				continue
			}
			if err := writeOutPage(base, item); err != nil {
				log.Println(err)
			}
		}
	}
}

func downloadAttachments(base string, export *BlogExport) {
	for _, item := range export.Channel.Items {
		if item.PostType == "attachment" {
			log.Printf("  - %s", item.AttachmentURL)
			// TODO: Generate pages for these?
			u, err := url.Parse(item.AttachmentURL)
			if err != nil {
				log.Printf("%s: %v", color.RedString("    unable to download"), err)
				continue
			}
			localPath := filepath.Join(base, u.Path)
			if _, err = os.Stat(localPath); err == nil {
				log.Println(color.BlueString("    File exists:"), localPath)
			} else {
				err = downloadFile(localPath, item.AttachmentURL)
				if err != nil {
					log.Printf("%s: %v", color.RedString("    unable to download"), err)
				}
			}
		}
	}
}

func downloadFile(dest, attachmentURL string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0700); err != nil {
		return err
	}
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	resp, err := http.Get(attachmentURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	return err
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

	tags := []string{item.PostType}
	categories := make(map[string][]string)
	for _, cat := range item.Categories {
		categories[cat.Domain] = append(categories[cat.Domain], cat.Text)
		tags = append(tags, fmt.Sprintf("%s-%s", cat.Domain, cat.Text))
	}

	frontMatter := map[string]interface{}{
		"title":   item.Title,
		"date":    postDate.Format(OUTPUT_DATE_FORMAT),
		"layout":  "layout",
		"tags":    tags,
		"creator": item.Creator,
	}
	for domain, values := range categories {
		frontMatter[domain] = values
	}
	for _, el := range item.Contents {
		if el.XMLName.Space == EXCERPT_NS && el.Data != "" {
			frontMatter["summary"] = el.Data
		}
	}

	fmBytes, err := yaml.Marshal(frontMatter)
	if err != nil {
		return err
	}

	f.WriteString("---\n")
	f.Write(fmBytes)
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

func npmInstall(siteDir string) error {
	cmd := exec.Command("npm", "install")
	cmd.Dir = siteDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
