package app

/*StructToString to convert struct to string*/
type StructToString struct{}

/*Pet defined the pet's varialbes ...*/
type Pet struct {
	ID        int64      `json:"id"`
	Category  category   `json:"category"`
	Name      string     `json:"name"`
	PhotoUrls []photourl `json:"photoUrls"`
	Tags      []tags     `json:"tags"`
	Status    Status     `json:"status"`
}

/*Status use for Enum status variable ...*/
type Status string

const (
	available Status = "available"
	pending   Status = "pending"
	sold      Status = "sold"
)

/**
*	Define the Category struct
 */
type category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/**
*	Define the photourl for photoUrls and tags
 */
type photourl struct {
	// XMLName xml.Name `xml:"photourl"`
	Name    string `json:"name"`
	Wrapped bool   `json:"wrapped"`
}

/*Tag use for tags ...*/
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*Tags use for all tags ...*/
type tags struct {
	// XMLName xml.Name `xml:"tags"`
	Name    int64 `json:"name"`
	Wrapped bool  `json:"wrapped"`
}
