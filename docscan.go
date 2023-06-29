package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Record represents a single invoice record.
type Record struct {
	Group         string
	Date          time.Time
	Price         string
	Establishment string
	Category      string
}

// extractFields extracts the fields from the filename of an invoice.
// It returns an error if the filename does not contain enough parts.
func extractFields(filename string) (string, string, string, string, string, error) {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename)) // Strip file extension
	parts := strings.Split(filename, "-")
	if len(parts) < 6 { // At least group, establishment, category, price, date (5 parts) should be present
		return "", "", "", "", "", fmt.Errorf("file %s does not have enough parts", filename)
	}
	group := parts[1]
	price := parts[len(parts)-4]                                                        // Price is the fourth-last part
	establishment := parts[2]                                                           // Establishment is the third part
	category := strings.Join(parts[3:len(parts)-4], "-")                                // Category is all the parts between establishment and price
	date := parts[len(parts)-3] + "-" + parts[len(parts)-2] + "-" + parts[len(parts)-1] // Date is the last part
	return group, date, price, establishment, category, nil
}

// parseDate parses a date string in the format "2006-01-02".
// It returns an error if the date string is not in the correct format.
func parseDate(date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date: %v", err)
	}
	return parsedDate, nil
}

// getDirectoryFromUser prompts the user to enter a directory and returns it.
func getDirectoryFromUser() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please provide a directory. Using '.' for current directory: ")
	dir, _ := reader.ReadString('\n')

	return strings.TrimSpace(dir)
}

// scanDirectory scans a directory for invoice files and returns a slice of Records.
// It skips directories and irrelevant files.
// It logs an error and continues if it encounters a problem processing a file.
func scanDirectory(dir string) ([]Record, error) {
	var records []Record

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !isRelevantFile(info.Name()) {
			return nil
		}

		group, date, price, establishment, category, err := extractFields(info.Name())
		if err != nil {
			color.Red("Error processing file %s: %v", path, err)
			return nil
		}

		parsedDate, err := parseDate(date) // Parse the date
		if err != nil {
			color.Red("Error processing date for file %s: %v", path, err)
			return nil
		}

		records = append(records, Record{
			Group:         group,
			Date:          parsedDate,
			Price:         price,
			Establishment: establishment,
			Category:      category,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return records, nil
}

// isRelevantFile checks if a file name is relevant to the invoice scanning.
// A file is relevant if its name contains "rechnung" or "invoice".
func isRelevantFile(fileName string) bool {
	lowerFileName := strings.ToLower(fileName)
	return strings.Contains(lowerFileName, "rechnung") || strings.Contains(lowerFileName, "invoice")
}

// generateReport generates a report from a slice of Records.
// The report is grouped by the Group field of the Records.
// Each line of the report contains the Date, Price, Establishment, and Category fields of a Record.
func generateReport(records []Record) string {
	groups := make(map[string][]Record)
	for _, record := range records {
		groups[record.Group] = append(groups[record.Group], record)
	}

	report := ""
	for group, records := range groups {
		report += group + "\n"
		for _, record := range records {
			report += fmt.Sprintf("%s\t%s\t%s\t%s\n",
				record.Date.Format("2006-01-02"),
				record.Price,
				record.Establishment,
				record.Category,
			)
		}
		report += "\n"
	}

	return report
}

// main is the entry point of the program.
// It scans a directory for invoice files and prints a report.
// The directory is obtained from a command-line argument, if provided.
// Otherwise, the user is prompted to enter a directory.
func main() {
	var dir string
	if len(os.Args) > 1 { // If a command-line argument is provided, use it as the directory
		dir = os.Args[1]
	} else { // Otherwise, prompt the user to enter a directory
		dir = getDirectoryFromUser()
	}

	documents, err := scanDirectory(dir)
	if err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		return
	}

	fmt.Println("Scanning complete. Preparing report...")
	report := generateReport(documents)
	fmt.Println(report)
}
