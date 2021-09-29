# Minecraft Server Stat Tracker
A golang-based web app that parses minecraft server logs and provides information relevant to the server owner

## Setup
Execution requires working installations of golang and python3

> 1. Place server logs into `logs/` directory.
> 2. Run `fix_filenames.py` to re-order the numbering on log-files. (For some reason its reverse-chronological by default)
> 3. Run `go run main.go` to start the web server.
> 4. Navigate to localhost:8080 to view the page.
