# 📚 DocScan

🎉 Welcome to the comprehensive guide for the DocScan tool. DocScan is used to parse filenames from scanned documents, specifically those starting with 'invoice' or 'rechnung'. It extracts data from the filenames, outputs tables in the command line for you to copy and paste into your spreadsheets. This tool is typically used after another tool has scanned documents and saved them with relevant data as their filenames. Please note that the 'invoice' and 'rechnung' fields are not case sensitive.

## 🗂️ Project Structure

The project is organized into a simple file structure for easy navigation. Here's a brief overview of the key files and directories:

- [`docscan.go`](docscan.go): This is the main Go file where the logic of the parser is implemented.
- [`Makefile`](Makefile): This file is used for building and testing the project.

## 🚀 Getting Started

### How to Build and Run the Tool

1. **Build the project**: You can build the project using the Makefile by running `make build` in the root directory of the project.

2. **Run the tool**: After building, you can run the tool with `./docscan -path=<directory>` where `<directory>` should be replaced with the path of the directory containing the files you want to process.

### File Naming Convention

The tool expects files to be named according to the following pattern: 

`{invoice}-{group}-{establishement}-{category}-{amount}-{yyyy}-{mm}-{dd}.pdf`

The `{}` placeholders should be replaced with the actual data. Here is an example of a correctly formatted filename:

`invoice-IT-amazon-electronics-200,00-2023-06-29.pdf`

The tool will parse filenames according to this pattern and output the parsed data as a table in the command line.

## 🔧 Customizing the Setup

You can customize the tool's data parsing capabilities by modifying the code in `docscan.go`. Here are some steps to follow if you want to add more data parsing capabilities:

1. **Add a new field to the struct**: In the `Invoice` struct, add a new field for the data you want to parse. For example, if you want to parse a 'payment method' from the filename, you might add a `PaymentMethod string` field to the struct.

2. **Modify the regular expression**: Update the regular expression in the `parseInvoice` function to capture the new data from the filename. 

3. **Update the parsing logic**: In the `parseInvoice` function, extract the new data from the matched regular expression and assign it to the new field in the `Invoice` struct.

4. **Update the output format**: Finally, update the `String` method of the `Invoice` struct to include the new field in the output format.

Keep in mind that any changes to the parsing logic should also be reflected in the expected file naming convention.

## 📚 Additional Resources

For more information about Go, you can refer to the following resources:

- [Go Documentation](https://golang.org/doc/)
- [Go Playground](https://play.golang.org/)
- [A Tour of Go](https://tour.golang.org/welcome/1)

🌟 Enjoy parsing your invoice data with DocScan!
