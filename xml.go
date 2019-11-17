package main

import "encoding/xml"

type BlogExport struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
		Language    string `xml:"language"`
		BaseSiteURL string `xml:"base_site_url"`
		BaseBlogURL string `xml:"base_blog_url"`
		Authors     []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"author_id"`
			Login       string `xml:"author_login"`
			Email       string `xml:"author_email"`
			DisplayName string `xml:"author_display_name"`
			FirstName   string `xml:"author_first_name"`
			LastName    string `xml:"author_last_name"`
		} `xml:"author"`
		Categories []struct {
			Text        string `xml:",chardata"`
			TermID      string `xml:"term_id"`
			Nicename    string `xml:"category_nicename"`
			Parent      string `xml:"category_parent"`
			Name        string `xml:"cat_name"`
			Description string `xml:"category_description"`
		} `xml:"category"`
		Tags []struct {
			TermID string `xml:"term_id"`
			Slug   string `xml:"tag_slug"`
			Name   string `xml:"tag_name"`
		} `xml:"tag"`
		Terms []struct {
			ID          string `xml:"term_id"`
			Taxonomy    string `xml:"term_taxonomy"`
			Slug        string `xml:"term_slug"`
			Parent      string `xml:"term_parent"`
			Name        string `xml:"term_name"`
			Description string `xml:"term_description"`
		} `xml:"term"`
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Creator string `xml:"creator"`
	Guid    struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	Description   string           `xml:"description"`
	Contents      []EncodedContent `xml:"encoded"`
	PostID        string           `xml:"post_id"`
	PostDate      string           `xml:"post_date"`
	PostDateGmt   string           `xml:"post_date_gmt"`
	CommentStatus string           `xml:"comment_status"`
	PingStatus    string           `xml:"ping_status"`
	PostName      string           `xml:"post_name"`
	Status        string           `xml:"status"`
	PostParent    string           `xml:"post_parent"`
	MenuOrder     string           `xml:"menu_order"`
	PostType      string           `xml:"post_type"` // page - attachment
	PostPassword  string           `xml:"post_password"`
	IsSticky      string           `xml:"is_sticky"`
	AttachmentURL string           `xml:"attachment_url"`
	Postmeta      []struct {
		Text      string `xml:",chardata"`
		MetaKey   string `xml:"meta_key"`
		MetaValue string `xml:"meta_value"`
	} `xml:"postmeta"`
	Categories []struct {
		Text     string `xml:",chardata"`
		Domain   string `xml:"domain,attr"` // post_tag || category
		Nicename string `xml:"nicename,attr"`
	} `xml:"category"`
}

type EncodedContent struct {
	XMLName xml.Name
	Data    string `xml:",chardata"`
}

const (
	EXCERPT_NS string = "http://wordpress.org/export/1.2/excerpt/"
	CONTENT_NS string = "http://purl.org/rss/1.0/modules/content/"
)
