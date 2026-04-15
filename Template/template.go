package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type Employee struct {
	Subject string
	Body    string
	Name    string
}

func main() {
	fmt.Println("String template")

	temp := `
        {{.Subject}}
		{{.Body}}
		best Regards
			{{.Name}}
	`
	templ, _ := template.New("template").Parse(temp)
	emp := Employee{
		Subject: "Zafar OOO ",
		Body:    "i am taking leave due fevar i am not feeeling weel",
		Name:    "zafar Imam",
	}
	err := templ.Execute(os.Stdout, emp)
	if err != nil {
		fmt.Println(err)
	}
	htmlTemalte()
}

func htmlTemalte() {
	htmlTemp := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Subject}}</title>
</head>
<body>
    <h1>{{.Subject}}</h1>
    <p>{{.Body}}</p>
    <footer>
        <strong>Best Regards,</strong> </br> {{.Name}}
    </footer>
</body>
</html>
`

	// Parse the template
	templ, err := template.New("employee").Parse(htmlTemp)
	if err != nil {
		panic(err)
	}

	emp := Employee{
		Subject: "Welcome Zafar",
		Body:    "This is your HTML template example.",
		Name:    "Zafar",
	}

	var buf bytes.Buffer
	if err := templ.Execute(&buf, emp); err != nil {
		panic(err)
	}

	result := buf.String()
	fmt.Println(result)

	fileName := "output.html"
	if err := os.WriteFile(fileName, buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	exec.Command("rundll32", "url.dll,FileProtocolHandler", fileName).Start()
}
