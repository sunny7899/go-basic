package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// clearScreen removes all the existing artifacts in console screen
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func formatAnonymousObject(input, prefix, indent string) (string, error) {
	var buf bytes.Buffer
	decoder := xml.NewDecoder(bytes.NewReader([]byte(input)))
	encoder := xml.NewEncoder(&buf)
	
	encoder.Indent(prefix, indent)
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		if err := encoder.EncodeToken(token); err != nil {
			return "", err
		}
	}

	if err := encoder.Flush(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func formatKnownObject(input, prefix, indent string) (string, error) {
	type Person struct {
		XMLName 	xml.Name	`xml:"person"`
		FirstName	string		`xml:"name>firstName"`
		LastName	string		`xml:"name>lastName"`
		Age			int8		`xml:"age"`
	}
	var person Person
	if err := xml.Unmarshal([]byte(input), &person); err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		os.Exit(1)
	}
	output, err := xml.MarshalIndent(person, prefix, indent)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// formatXML formats the XML data provided
func formatXML(input string) (string, error) {
	//change accordingly to test different scenarios
	isXMLStructureKnown := true
	//string which gets placed before every line
	prefix := ""

	//string which also gets placed before every line
	//but with each successive nesting, its instances increase
	//ideally Unicode space characters (like Tab or Space)
	indent := "	"
	if !isXMLStructureKnown {
		return formatAnonymousObject(input, prefix, indent)
	}
	return formatKnownObject(input, prefix, indent)
}

func main() {
	fmt.Print("Enter XML data: ")

	// Read input as a string
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Format XML
	output, err := formatXML(input)
	if err != nil {
		fmt.Println("Error formatting XML:", err)
		os.Exit(1)
	}

	// Print formatted XML
	fmt.Println("Formatted XML:")
	fmt.Println(output)
}

// SAMPLE INPUT OUTPUT AND EXPLANATION
// ---
// XML Formatter Example
// Input:
// <person><name><firstName>Guru</firstName><lastName>Dutt</lastName></name><age>67</age></person>


// Output:

// <person>
//         <name>
//                 <firstName>Guru</firstName>
//                 <lastName>Dutt</lastName>
//         </name>
//         <age>67</age>
// </person>

// Explanation:
// -Input: A single-line XML string with no formatting.
// -Output: The XML formatted with indentation and line breaks, making it easy to read.
