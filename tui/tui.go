package tui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// User data
type Data struct {
	Token   string `json:"token"`
	Posting string `json:"posting"`
	Name    string `json:"name"`
	Email   string `json:"email"`

	Resume   string `json:"resume,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Employer string `json:"employer,omitempty"`
	Source   string `json:"source,omitempty"`
	Comment  string `json:"comments,omitempty"`
}

// Print welcome message
func WelcomeMessage() {
	fmt.Println("Welcome to ApplyByAPI client (built on ApplyByAPI's API xD)")
}

// Print notification about new token
func TokenNotification(token string) {
	fmt.Println("Ok, we got token for posting:", token, "you have 5 minutes to send your info")
}

// Run survey to collect data
func RunSurvey() *Data {
	data := &Data{}
	fmt.Println("So, let's start the survey. That data will be sent to ApplyByAPI posting (vacation)")
	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)

	data.Name = readString("Name [required]: ", reader)
	data.Email = readString("Email [required]: ", reader)

	data.Resume = readString("Resume .pdf file path [required]: ", reader)
	data.Phone = readString("Phone [optional]: ", reader)
	data.Employer = readString("Employer [optional]: ", reader)
	data.Source = readString("Source [optional]: ", reader)
	data.Comment = readString("Comments [optional]: ", reader)

	return data
}

// Last message
func Done(id int) {
	fmt.Println("---------------------------")
	fmt.Println("Well done! Your application ID:", id)
	fmt.Println("Check your email for confirmation")
}

// Read string input from STDIN
func readString(prefix string, reader *bufio.Reader) string {
	fmt.Print(prefix)
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}
