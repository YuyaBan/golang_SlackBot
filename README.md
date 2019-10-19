# golang_SlackBot
## Description
* A Simple SlackBot made in Go (Golang).
* this Bot imitates a respected senior!

## How to build in localhost
1. Please prepare config.toml. config.toml contain Token and BotId.
* You can deploy easily via Dockerfile.
1. docker build -t <SlackBot_Name> . 
2. docker run -e "PORT=<port_number>" -p <port_number>:<port_number> -t <SlackBot_Name>

## License
[MIT](https://github.com/YuyaBan/golang_SlackBot/blob/master/LICENSE)

## Author
[Yuya Ban](https://github.com/YuyaBan)