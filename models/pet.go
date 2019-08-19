package models

/*Pet defined the pet's varialbes ...*/
type Pet struct {
	ID        int64    `json:"id"`
	Category  Category `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    Status   `json:"status"`
}

/*Status use for Enum status variable ...*/
type Status string

const (
	available Status = "available"
	pending   Status = "pending"
	sold      Status = "sold"
)

// Returning the status value for checking Status ISValid
func (st Status) String() string {
	switch st {
	case available:
		return "available"
	case pending:
		return "pending"
	case sold:
		return "sold"
	default:
		return "INVALID"
	}
}

// IsValid checking the input Status value is valid or not.
func (st Status) IsValid() bool {
	return st.String() != "INVALID"
}

/*Category Define the Category struct*/
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*Tag use for tags ...*/
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/*PetNotFound type for Pet Not Found error*/
type PetNotFound struct {
	Code    int64  `json:"code"`
	Type    string `json:"error"`
	Message string `json:"message"`
}

/*ResponseBody type for Reponse Body Message*/
type ResponseBody struct {
	Code    int64  `json:"code"`
	Type    string `json:"error"`
	Message string `json:"message"`
}
