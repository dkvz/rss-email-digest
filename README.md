# RSS Feed to email alert / digest
I couldn't find a project I liked for this so here I am creating my own as always.

**Work in progress** There's currently no SMTP authentication options.

## Config
See `.env.example` for the available options.

- SLEEP_INTERVAL (in seconds) is optional and defaults to the const set in `config.go` which should be serveral minutes.
- SMTP_HOST is optional and defaults to "localhost".

## TODO
- [] Check for email validity