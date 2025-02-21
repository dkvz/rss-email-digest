# RSS Feed to email alert / digest
I couldn't find a project I liked for this so here I am creating my own as always.

**Work in progress** There's currently no SMTP authentication options.

## Config
See `.env.example` for the available options.

- SLEEP_INTERVAL (in seconds) is optional and defaults to the const set in `config.go` which should be serveral minutes.
- SMTP_HOST is optional and defaults to "localhost".

## Fields to use in the RSS
We're mostly concerned with the <item> elements, which have these relevant fields:
* title
* link
* guid (usually = link but we can use it for hashing)
* pubDate

## Notifications
Using `gomail` to send email notifications, documentation here: https://pkg.go.dev/gopkg.in/gomail.v2#section-readme

## TODO
- [] Check for email validity