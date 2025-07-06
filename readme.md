![Workflow Status](https://github.com/Calvinbullock/quote-cli/actions/workflows/go-ci.yml/badge.svg)
# Quote CLI

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
    - [ ] add single letter flags (-a = --author, -t = --tag, etc)
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
