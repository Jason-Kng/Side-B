# Side-B

This will be a TUI for a vinyl record collection. This is very very early development

I am using this project to learn how to use golang and how to interface with an API

What am I using to make this?
- I would like to connect to the Discogs api in order to let you interact with your collection through this app
- I plan to use Bubble Tea and Lip Gloss for the TUI
- I am using a sqlite local database for storing the data

## Installation Instructions
** WARNING: This project is very early days and its very bad right now lol **
#### Build From Source
```sh
git clone https://github.com/Jason-Kng/Side-B.git && cd Side-B && go build -o Side-B ./cmd/Side-B/main.go
```
Make sure to make a .env file and store your discogs api key
touch .env && echo "DISCOGS_TOKEN=enter_your_discogs_token">> .env
