# Web Scraping Tool

This is a command-line tool written in Go for web scraping. It uses the "github.com/gocolly/colly" package to extract HTML content from a specified URL based on provided HTML tags.

## Installation

1. Install Go on your machine if you haven't already. Refer to the official Go documentation for installation instructions: [https://golang.org/doc/install](https://golang.org/doc/install)

2. Clone the repository or download the source code files.

3. Navigate to the directory containing the source code files.

4. Build the executable by running the following command:
   ```
   go build
   ```

## Usage

To use the web scraping tool, follow these steps:

1. Open a terminal or command prompt.

2. Navigate to the directory containing the built executable.

3. Run the executable with the following command:
   ```
   ./www2txt -url [URL] -tag [HTML tag 1] -tag [HTML tag 2] ...
   ```

   Replace `[URL]` with the target website URL you want to scrape. Specify one or more HTML tags to extract content from by using the `-tag` flag followed by the HTML tag. For example:
   ```
   ./www2txt -url https://example.com -tag h1 -tag p
   ```

   The tool will visit the specified URL and extract the text content from the provided HTML tags.

4. The extracted text will be saved in a file named `[hostname].txt`, where `[hostname]` is the hostname of the target website.

   For example, if the target URL is "https://example.com", the extracted text will be saved in a file named "example.com.txt".

## Example

Here's an example usage of the web scraping tool:

```
./www2txt -url https://example.com -tag h1 -tag p
```

This command will visit "https://example.com" and extract the text content from all `<h1>` and `<p>` tags. The extracted text will be saved in a file named "example.txt".

## Notes

- Ensure that you have proper permissions to write files in the directory where the executable is located.

- Make sure to provide valid HTML selectors as arguments for the `-tag` flag.

- The tool only extracts text content from HTML tags and does not handle other types of content or complex web scraping scenarios.

- It's important to be aware of and follow the legal and ethical guidelines regarding web scraping. Make sure you have the necessary permissions and comply with the website's terms of service before scraping its content.
