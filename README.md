# Fact Retriever

Fact Retriever is a Go program that scrapes rhino facts from the website [factretriever.com](https://www.factretriever.com) and saves them to a JSON file. It utilizes the `colly` library for web scraping and `encoding/json` for JSON serialization.

## How It Works

The program performs the following steps:

1. Imports necessary packages, including `colly` for web scraping and `encoding/json` for JSON encoding/decoding.
2. Defines a `Fact` struct to hold information about each rhino fact.
3. Creates a new collector from `colly`, setting the allowed domains to prevent visiting other websites.
4. Defines an HTML callback function to extract fact IDs and descriptions from the website.
5. Visits the rhino facts page on `factretriever.com` and collects the data.
6. Encodes the collected data to JSON and writes it to a file named `rhinofacts.json`.


## JSON Output

The program will scrape rhino facts from the website and save them as a JSON file named `rhinofacts.json` in the same directory as the executable.

## Contributing

Contributions to this project are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- The `colly` library for making web scraping in Go easy and efficient.
- [FactRetriever](https://www.factretriever.com) for providing the rhino facts used in this project.
