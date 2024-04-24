STEAMPIPE_INSTALL_DIR?=~/.steampipe

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/local/teamwork@latest/steampipe-plugin-teamwork.plugin *.go
