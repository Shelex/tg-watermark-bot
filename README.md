# tg-watermark-bot

Simple-stupid telegram bot that can use watermark image overlay over video input with ffmpeg.  
Uses long-polling, which is not ideal (webhook would be better), but just works.  
Supports localization: Ukrainian (uk) and English (en), could be configured via `config.env` file or env variables (with `BOT_` prefix).  
~~Working example: [@odarka_watermarka_bot](https://t.me/odarka_watermarka_bot)~~

# Requirements

- [Go](https://go.dev/)
- Register bot with [@BotFather](https://t.me/Botfather)
- ffmpeg should be installed
    - macos (with brew, fyi: may require `xcode-select --install`)
        ```bash
            brew install ffmpeg
        ```
    - ubuntu
        ```bash
            sudo apt install ffmpeg
        ```
- Configure `config.env` (you can check `sample.env` as an example):
    - set `TELEGRAM_TOKEN` (could be obtained from [@BotFather](https://t.me/Botfather))
    - set `ENVIRONMENT` - if it is `dev` - debug logs will be displayed
    - set `LOCALE` - supports `en` for english and `uk` for ukrainian, `uk` by default.

# Run

```bash
    go run .
```

# Build

binary will be available in `bin` folder.  

- for ubuntu
    ```bash
        GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/tgbot-watermark
    ```
- for macos
    ```bash
        GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/tgbot-watermark
     ```
