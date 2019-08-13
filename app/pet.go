package app

import (
	"encoding/xml"
)

/*StructToString to convert struct to string*/
type StructToString struct{}

/*Scan to convert struct to string*/
// func (s *StructToString) Scan(v interface{}) error {
// 	// Change v from struct to JSON then convert to String
// 	valueStruct, err := json.Marshal(v)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(string(valueStruct))
// 	value1 := string(valueStruct)
// 	*s = StructToString(value1)
// 	return nil
// }

// /*Pet defined the pet's varialbes ...*/
// type Pet struct {
// 	ID        int64     `json:"id"`
// 	Category  category  `json:"category"`
// 	Name      string    `json:"name"`
// 	PhotoUrls PhotoUrls `json:"photoUrls"`
// 	tags      Tags      `json:"tags"`
// 	Status    Status    `json:"status"`
// }

/*Pet defined the pet's varialbes ...*/
type Pet struct {
	ID        int64  `db:"id"`
	Category  string `db:"category"`
	Name      string `db:"name"`
	PhotoUrls string `db:"photoUrls"`
	tags      string `db:"tags"`
	Status    string `db:"status"`
}

/*Status use for Enum status variable ...*/
type Status int

const (
	available Status = iota
	pending
	sold
)

func (s Status) String() string {
	return [...]string{"available", "pending", "sold"}[s]
}

/**
*	Define the Category struct
 */
type category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/**
*	Define the OrderMap for photoUrls and tags
 */
type photourl struct {
	XMLName xml.Name `xml:"OrderedMap"`
	Name    string   `xml:"name"`
	Wrapped bool     `xml:"wrapped"`
}

/*PhotoUrls XMLName, OrderedMaps ...*/
type PhotoUrls struct {
	XMLName xml.Name   `xml:"photoUrls"`
	PUrls   []photourl `xml:"photourl"`
}

/*Tag use for tags ...*/
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

/**
*	Define tag for tags
 */
type tag struct {
	XMLName xml.Name `xml:"OrderedMap"`
	Name    Tag      `xml:"name"`
	Wrapped bool     `xml:"wrapped"`
}

/*Tags use for all tags ...*/
type Tags struct {
	XMLName xml.Name `xml:"tags"`
	tags    []tag    `xml:tag"`
}
