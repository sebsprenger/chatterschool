# chatterschool
Some code to facilitate a [Hacker School workshop ](https://www.hacker-school.de/kurse/S09). This repository contains a chat server as well as a chat client both written in go. They use a websocket connection to talk to each other.

The server offers an interface to provide an opportunity to eavesdrop messages and even modify them to censor information or spread misinformation. We intend to play and discuss in this context.

The client provides two interfaces to modify messages 1) before they are being sent and 2) after they have been received, but before they are being displayed. With that we intend to implement commands (e.g. expansion of abbrevations) and emoji like replacements.

## Installation

  go get github.com/sebsprenger/chatterschool/cmd/server
  go get github.com/sebsprenger/chatterschool/cmd/client

## Start

  server [-port 9001]
  client [-port 9001] [-name Kitten2000]
