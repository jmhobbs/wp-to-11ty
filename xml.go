package main

import (
	"encoding/xml"
	"net/url"
	"strings"
)

type BlogExport struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string     `xml:"title"`
		Link        string     `xml:"link"`
		Description string     `xml:"description"`
		PubDate     string     `xml:"pubDate"`
		Language    string     `xml:"language"`
		BaseSiteURL string     `xml:"base_site_url"`
		BaseBlogURL string     `xml:"base_blog_url"`
		Authors     []Author   `xml:"author"`
		Categories  []Category `xml:"category"`
		Tags        []struct {
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
	Comments []Comment `xml:"comment"`
}

func (i Item) PermalinkPath() (string, error) {
	u, err := url.Parse(i.Link)
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(u.Path, "/") {
		return u.Path, nil
	}
	return u.Path + "/", nil
}

type EncodedContent struct {
	XMLName xml.Name
	Data    string `xml:",chardata"`
}

type Author struct {
	Text        string `xml:",chardata" json:"-"`
	ID          string `xml:"author_id" json:"-"`
	Login       string `xml:"author_login" json:"-"`
	Email       string `xml:"author_email" json:"email"`
	DisplayName string `xml:"author_display_name" json:"display_name"`
	FirstName   string `xml:"author_first_name" json:"first_name"`
	LastName    string `xml:"author_last_name" json:"last_name"`
}

type Category struct {
	Text        string `xml:",chardata" json:"-"`
	TermID      string `xml:"term_id" json:"-"`
	Nicename    string `xml:"category_nicename" json:"nice_name"`
	Parent      string `xml:"category_parent" json:"parent"`
	Name        string `xml:"cat_name" json:"name"`
	Description string `xml:"category_description" json:"description"`
}

type Comment struct {
	ID          int    `xml:"comment_id"`
	Author      string `xml:"comment_author"`
	AuthorEmail string `xml:"comment_author_email"`
	AuthorURL   string `xml:"comment_author_url"`
	AuthorIP    string `xml:"comment_author_IP"`
	Date        string `xml:"comment_date"`
	DateGMT     string `xml:"comment_date_gmt"`
	Content     string `xml:"comment_content"`
	Approved    int    `xml:"comment_approved"`
	Type        string `xml:"comment_type"`
	ParentID    int    `xml:"comment_parent"`
	UserID      int    `xml:"comment_user_id"`
}

func (c Comment) toOutputComment() OutputComment {
	return OutputComment{
		ID: c.ID,
		Author: OutputCommentAuthor{
			Name:  c.Author,
			Email: c.AuthorEmail,
			URL:   c.AuthorURL,
			IP:    c.AuthorIP,
		},
		Date:     c.DateGMT,
		Content:  c.Content,
		Approved: c.Approved == 1,
		Type:     c.Type,
		ParentID: c.ParentID,
	}
}

type OutputCommentAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
	IP    string `json:"ip"`
}

type OutputComment struct {
	ID       int                 `json:"id"`
	Author   OutputCommentAuthor `json:"author"`
	Date     string              `json:"date"`
	Content  string              `json:"content"`
	Approved bool                `json:"approved"`
	Type     string              `json"type"`
	ParentID int                 `json:"parent_id"`
}

const (
	EXCERPT_NS string = "http://wordpress.org/export/1.2/excerpt/"
	CONTENT_NS string = "http://purl.org/rss/1.0/modules/content/"
)
