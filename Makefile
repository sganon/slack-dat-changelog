# Intended for local use only
NAME = slack-dat-changelog
OS = $(shell go env GOOS)

all: $(OS)

PLATFORMS := linux darwin
TARGET_OS = $(word 1, $@)
$(PLATFORMS):
	@printf "\e[36mBuilding $(NAME) for target os: $(TARGET_OS)\e[0m\n"
	@GOOS=$(TARGET_OS) GOARCH=amd64 go build -o $(NAME)

