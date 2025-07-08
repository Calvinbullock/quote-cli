![Workflow Status](https://github.com/Calvinbullock/quote-cli/actions/workflows/go-ci.yml/badge.svg)
# Quote CLI

## Default quote file location and quote format:
Here's where you can typically find it:
- Linux: `$XDG_CONFIG_HOME/quote-cli/default.json` or, if `XDG_CONFIG_HOME` is not set, `~/.config/quote-cli/default.json`
- macOS: `~/Library/Application Support/quote-cli/default.json`
- Windows: `%APPDATA%\quote-cli\default.json` (e.g., `C:\Users\<YourUsername>\AppData\Roaming\quote-cli\default.json`)

```
[
  {
    "author": "John F. Kennedy",
    "text": "A nation that is afraid to let its people judge the truth and falsehood in an open market is a nation that is afraid of its people.",
    "tags": ["February 26, 1962"]
  },
  {
    "author": null,
    "text": "To change yourself you must first change your surroundings",
    "tags": []
  }
]
```

## RUNING
- `go run ./cmd/quote-cli` - for production

#### Manuel Build / run
- `go build -o quote-cli ./cmd/quote-cli` - build the binary
- `./quote-cli`  - run the binary

## plans -- Stories
 - TODO: 
    - [ ] search by keyword
    - [ ] search by author
        - [x] search by author basic
        - [ ] search by partial author basic
    - [ ] Combine Filters (EX: use both --tag and --author search)
    - [x] add single letter flags (-a = --author, -t = --tag, etc)
    - [ ] limit the total print count (`--limit <number>`)
    - [ ] add / delete a quote
        - [ ] print all quotes with an ID?
    - [ ] different outputs to terminal (basic, json, csv, etc)

- TDD - TEST DRIVEN!! FROM THE START
- a program that display's a quote every time I start a terminal or with a cmd and can add new ones with cli tool
    - nice formatting (when printing)
    - a help doc / tldr
    - store in a csv -- for now
    - add a flag for a category -- "motivational", "emotinal", etc (be able to add and list flags for a qoute)

## Stories
- "sometimes while programming you need a quick burst of motivation, one that does not take you out of the flow, one that does not distract you, here it is qoutes-cli."
