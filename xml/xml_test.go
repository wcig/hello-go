package xml

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"
)

type user struct {
	Id       int      `xml:"id"`
	Name     string   `xml:"name,attr"`
	Favorite []string `xml:"like"`
}

// 1.struct -> xml str (not support map)
func TestXml01(t *testing.T) {
	u := &user{
		Id:   100,
		Name: "Tom",
		Favorite: []string{
			"basketball",
			"football",
		},
	}

	bytes, err := xml.Marshal(u)
	// bytes, err := xml.MarshalIndent(&u, "", "    ")
	if err != nil {
		log.Fatal("xml marshal err:", err)
	}
	fmt.Println("xml str:", string(bytes))
}

// xml str: <user name="Tom"><id>100</id><like>basketball</like><like>football</like></user>

// 2.xml str -> struct
func TestXml02(t *testing.T) {
	xmlStr := `<user name="Tom"><id>100</id><like>basketball</like><like>football</like></user>`
	var u user
	if err := xml.Unmarshal([]byte(xmlStr), &u); err != nil {
		log.Fatal("xml unmarshal err:", err)
	}
	fmt.Println("user:", u)
}

// user: {100 Tom [basketball football]}

// 3.write xml file
func TestXml03(t *testing.T) {
	u := &user{
		Id:   100,
		Name: "Tom",
		Favorite: []string{
			"basketball",
			"football",
		},
	}

	xmlFile, err := os.Create("user.xml")
	if err != nil {
		log.Fatal("create xml file err:", err)
	}
	defer xmlFile.Close()

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "    ") // format
	if err := encoder.Encode(u); err != nil {
		log.Fatal("xml encoder err:", err)
	}
	fmt.Println("create xml file success!!!")
}

// 4.read xml file
func TestXml04(t *testing.T) {
	xmlFile, err := os.Open("user.xml")
	if err != nil {
		log.Fatal("open xml file err:", err)
	}
	defer xmlFile.Close()

	var u user
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&u); err != nil {
		log.Fatal("xml decoder err:", err)
	}
	fmt.Println("user:", u)
}

// user: {100 Tom [basketball football]}
