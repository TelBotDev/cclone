package main

import (
	_ "github.com/TelBotDev/cclone/backend/all" // import all backends
	"github.com/rclone/rclone/cmd"
	_ "github.com/TelBotDev/cclone/cmd/copy"
	_ "github.com/rclone/rclone/cmd/all"    // import all commands
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

func main() {
	cmd.Main()
}
