# GoLang JSONL Example

This GoLang example showcases how to read a JSONL (JSON Lines) file and print details about a person on each iteration.

## Installation

1. Clone the repository:

```
$ clone https://github.com/CamTrew/jsonl-example.git
```

2. Navigate to the project directory:

```
$ cd jsonl-example
```

3. Build the project:

```
$ go build
```

4. Run the executable:

```
$ ./jsonl-example
```

## Usage

This example reads a `people.jsonl` file located in the root directory of the project. The file contains a list of people, each represented by a JSON object on a separate line.

The program reads each line of the file, unmarshals the JSON object into a `Person` struct, and prints the details of the person. We are able to process an individual object at a time and do not need to read all of the people into RAM at once.