# TicTacToe CLI client

![License](https://img.shields.io/github/license/Bananenpro/tictactoe-cli)
![Go version](https://img.shields.io/github/go-mod/go-version/Bananenpro/tictactoe-cli)

A simple cli frontend for my [TicTacToe multiplayer server](https://github.com/Bananenpro/tictactoe-backend).

## Features

- Implements all features of the server
- Colorful winning row
- Optional AI using the [minimax search](https://en.wikipedia.org/wiki/Minimax) algorithm

## Setup

### Prerequisites

- [Go](https://go.dev/) 1.17+
- A terminal emulator that supports ANSI escape sequences

### Cloning the repo

```sh
git clone https://github.com/Bananenpro/tictactoe-cli.git
cd tictactoe-cli
```

### Building

```sh
go build -o tictactoe-cli ./cmd/main.go
```

### Running

To play:

```sh
./tictactoe-cli
```

To let the computer play for you:

```sh
./tictactoe-cli --ai
```

## License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Copyright

Copyright © 2022 Julian Hofmann
