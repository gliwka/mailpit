# Mailpit - email testing for developers

![Tests](https://github.com/axllent/mailpit/actions/workflows/tests.yml/badge.svg)
![Build status](https://github.com/axllent/mailpit/actions/workflows/release-build.yml/badge.svg)
![Docker builds](https://github.com/axllent/mailpit/actions/workflows/build-docker.yml/badge.svg)
![CodeQL](https://github.com/axllent/mailpit/actions/workflows/codeql-analysis.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/axllent/mailpit)](https://goreportcard.com/report/github.com/axllent/mailpit)

Mailpit is a multi-platform email testing tool & API for developers.

It acts as both an SMTP server, and provides a web interface to view all captured emails.

Mailpit is inspired by [MailHog](#why-rewrite-mailhog), but much, much faster.

![Mailpit](https://raw.githubusercontent.com/axllent/mailpit/develop/docs/screenshot.png)


## Features

- Runs entirely from a single binary, no installation required
- SMTP server (default `0.0.0.0:1025`)
- Web UI to view emails (formatted HTML, highlighted HTML source, text, headers, raw source and MIME attachments including image thumbnails)
- Mobile and tablet HTML preview toggle in desktop mode
- Advanced mail search ([see wiki](https://github.com/axllent/mailpit/wiki/Mail-search))
- Message tagging ([see wiki](https://github.com/axllent/mailpit/wiki/Tagging))
- Real-time web UI updates using web sockets for new mail
- Optional browser notifications for new mail (HTTPS only)
- Configurable automatic email pruning (default keeps the most recent 500 emails)
- Email storage either in a temporary or persistent database ([see wiki](https://github.com/axllent/mailpit/wiki/Email-storage))
- Fast SMTP processing & storing - approximately 70-100 emails per second depending on CPU, network speed & email size, easily handling tens of thousands of emails
- SMTP relaying / message release - relay messages via a different SMTP server ([see wiki](https://github.com/axllent/mailpit/wiki/SMTP-relay))
- Optional SMTP with STARTTLS & SMTP authentication, including an "accept anything" mode ([see wiki](https://github.com/axllent/mailpit/wiki/SMTP-with-STARTTLS-and-authentication))
- Optional HTTPS for web UI ([see wiki](https://github.com/axllent/mailpit/wiki/HTTPS))
- Optional basic authentication for web UI ([see wiki](https://github.com/axllent/mailpit/wiki/Basic-authentication))
- A simple REST API ([see docs](docs/apiv1/README.md))
- Multi-architecture [Docker images](https://github.com/axllent/mailpit/wiki/Docker-images)


## Installation

The Mailpit web UI listens by default on `http://0.0.0.0:8025`, and the SMTP port on `0.0.0.0:1025`.

Mailpit runs as a single binary and can be installed in different ways:


### Install via Brew (Mac)

Add the repository to your taps with `brew tap axllent/apps`, and then install Mailpit with `brew install mailpit`.


### Install via bash script (Linux & Mac)

Linux & Mac users can install it directly to `/usr/local/bin/mailpit` with:

```bash
sudo bash < <(curl -sL https://raw.githubusercontent.com/axllent/mailpit/develop/install.sh)
```


### Download static binary (Windows, Linux and Mac)

Static binaries can always be found on the [releases](https://github.com/axllent/mailpit/releases/latest). The `mailpit` binary can extracted and copied to your `$PATH`, or simply run as `./mailpit`.


### Docker

See [Docker instructions](https://github.com/axllent/mailpit/wiki/Docker-images).


### Compile from source

To build Mailpit from source see [building from source](https://github.com/axllent/mailpit/wiki/Building-from-source).


### Testing Mailpit

Please refer to [the documentation](https://github.com/axllent/mailpit/wiki/Testing-Mailpit) of how to easily test email delivery to Mailpit.


### Configuring sendmail

Mailpit's SMTP server (by default on port 1025), so you will likely need to configure your sending application to deliver mail via that port. A common MTA (Mail Transfer Agent) that delivers system emails to a SMTP server is `sendmail`, used by many applications including PHP. Mailpit can also act as substitute for sendmail. For instructions of how to set this up, please refer to the [sendmail documentation](https://github.com/axllent/mailpit/wiki/Configuring-sendmail).


## Why rewrite MailHog?

I had been using MailHog for a few years to intercept and test emails generated from several projects. MailHog has a number of performance issues, many of the frontend and Go modules are horribly out of date, and it is not actively developed.

Initially I tried to upgrade a fork of MailHog (both the UI as well as the HTTP server & API), but soon discovered that it is (with all due respect to its authors) poorly designed. It is in my opinion over-engineered (split over 9 separate projects), and performs very poorly when dealing with large amounts of emails or processing emails with an attachments (a single email with a 3MB attachment can take over a minute to ingest). Finally, the API transmits a lot of duplicate and unnecessary data on every browser request, and there is no HTTP compression.

In order to improve it I felt it needed to be completely rewritten, and so Mailpit was born.
