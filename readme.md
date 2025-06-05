# Telegram Arithmetic Bot

This is a simple Telegram bot that can solve mathematical expressions.

## Features

*   Solves mathematical expressions with the following operations:
    *   Addition (+)
    *   Subtraction (-)
    *   Multiplication (*)
    *   Division (/)
*   Supports parentheses.
*   Error handling for invalid expressions.

## Usage

1.  Start the bot in Telegram.
2.  Send the bot a mathematical expression.
    *   Example: `(43 + 27) * (5 - 3)`
3.  The bot will respond with the solution.

## To Run Bot
```bash
make run
# OR
go run cmd/api/main.go
```

## Configuration

The bot can be configured using environment variables:

*   `TELEGRAM_BOT_TOKEN`: (Required) The Telegram Bot Token.
*   `LOG_LEVEL`: (Optional) The log level.  Possible values: `debug`, `info`, `warn`, `error`.  Defaults to `info`.