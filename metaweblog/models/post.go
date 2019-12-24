package models

type Base64 struct {
}

type BlogInfo struct {
	BlogId   string `xml:"blogid"`
	URL      string `xml:"url"`
	BlogName string `xml:"blogName"`
}

type Enclosure struct {
	Length int    `xml:"length"`
	Type   string `xml:"type"`
	URL    string `xml:"url"`
}

type Source struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type Post struct {
	//DateCreated time.Time `xml:"dateCreated"`
	Description string    `xml:"description"`
	Title       string    `xml:"title"`
	Categories  []string  `xml:"categories"`
	Enclosure   Enclosure `xml:"enclosure"`
	Source      Source    `xml:"source"`
	//Link        string    `xml:"link"`
	WpSlug     string `xml:"wp_slug"`
	Permalink  string `xml:"permalink"`
	PostId     string `xml:"postid"`
	MtKeywords string `xml:"mt_keywords"`
}

type CategoryInfo struct {
	Description string `xml:"description"`
	HtmlURL     string `xml:"htmlUrl"`
	RssURL      string `xml:"rssUrl"`
	Title       string `xml:"title"`
	CategoryId  string `xml:"categoryid"`
}

type MediaObject struct {
	Name string `xml:"name"`
	Type string `xml:"type"`
	Bits Base64 `xml:"bits"`
}

type UrlData struct {
	Url string `xml:"url"`
}

type WpCategory struct {
	Name        string `xml:"name"`
	Slug        string `xml:"slug"`
	ParentId    int    `xml:"parent_id"`
	Description string `xml:"description"`
}

type GetPostResp struct {
	DateCreated string `json:"dateCreated"`
	Description string `json:"description"`
	Enclosure   struct {
		Length    int    `json:"length"`
		Link      string `json:"link"`
		Permalink string `json:"permalink"`
		Postid    int    `json:"postid"`
		Source    struct {
			MtKeywords string `json:"mt_keywords"`
		} `json:"source"`
	} `json:"enclosure"`
	Title string `json:"title"`
}

type PostLink struct {
	Link   string `json:"link"`
	Postid string `json:"postid"`
	Title  string `json:"title"`
}
