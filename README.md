# RSS Feed to email alert / digest
I couldn't find a project I liked for this so here I am creating my own as always.

**Work in progress** There's currently no SMTP authentication options.

## Config
See `.env.example` for the available options.

- SLEEP_INTERVAL (in seconds) is optional and defaults to the const set in `config.go` which should be serveral minutes.
- SMTP_HOST is optional and defaults to "localhost".

## Building the project
I run that line manually:
```
go build -o target/rssemaildigest cmd/rssemaildigest/main.go
```
And I often need "CGO_ENABLED=0" because of my target systems but performance isn't really a priority for this project.

## Fields to use in the RSS
We're mostly concerned with the <item> elements, which have these relevant fields:
* title
* link
* guid (usually = link but we can use it for hashing)
* pubDate

## Notifications
Using `gomail` to send email notifications, documentation here: https://pkg.go.dev/gopkg.in/gomail.v2#section-readme

## Status
For now I'll just save the state in the current directory the script is running in.

Exit with error if the directory is not writable.

## Systemd script

Put it in /etc/systemd/system/rss-email-digest.service:
```
[Unit]
Description=RSS feed alerts daemon
After=network.target network-online.target
Requires=network-online.target

[Service]
Type=simple
WorkingDirectory=/opt/rss-email-digest
User=www-data
Group=www-data
ExecStart=/opt/rss-email-digest/rssemaildigest
StandardError=journal
StandardOutput=journal

[Install]
WantedBy=multi-user.target
```
Make sure the binary name is correct in ExecStart.

Make sure the user or group is allowed to write to the working directory.

## TODO
- [x] If the status file wasn't there, we should just report the first item and that's it
- [ ] Feed content is usually HTML, we should either render it or make it go through some sort of [HTML to text thingy](https://github.com/jaytaylor/html2text)
- [ ] Check if current directory is writable or exit immediately
- [ ] Check for email validity
- [ ] We have error notifications by email, should probably send those from errors processing the feeds
- [ ] The goal was to have a single email for all the feeds, I'm sending multiple ones for now
