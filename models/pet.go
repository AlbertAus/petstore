package models

/*StructToString to convert struct to string*/
type StructToString struct{}

/*Pet defined the pet's varialbes ...*/
type Pet struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category"`
	Name      string     `json:"name"`
	PhotoUrls []Photourl `json:"photoUrls"`
	Tags      []Tags     `json:"tags"`
	Status    Status     `json:"status"`
}

/*Status use for Enum status variable ...*/
type Status string

const (
	available Status = "available"
	pending   Status = "pending"
	sold      Status = "sold"
)

/*Category Define the Category struct*/
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*Photourl Define the photourl for photoUrls and tags*/
type Photourl struct {
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
type Tags struct {
	// XMLName xml.Name `xml:"tags"`
	Name    int64 `json:"name"`
	Wrapped bool  `json:"wrapped"`
}
