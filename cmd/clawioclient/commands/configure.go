package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/clawio/system/cmd/clawioclient/config"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

var ConfigureCommand = cli.Command{
	Name:  "configure",
	Usage: "Configure ClawIO CLI options",
	Description: `
 Configure ClawIO  CLI options. If  this command is  run with no  arguments,
 you will be  prompted for configuration values such as  your ClawIO credentials
 (username  and password).  If  your config  file does  not  exist (the  default
 location is ~/.clawio/config),  the ClawIO CLI will create it  for you. To keep
 an existing value, hit enter when prompted for the value. When you are propmted
 for information,  the current  value will  be displayed  in [brackets].  If the
 config  item has  no value,  it  will be  displayed  as [None].  Note that  the
 configure command only work  with values from the config file.  It does not use
 any values from environment variables.

 Note: the ClawIO Access Token  obtained after validating the ClawIO Credentials
 will be written to the shared credentials file (~/.clawio/credentials).

 Note: ClawIO will log additional information to a log file (default location is
 ~/.clawio/log).
`,
	ArgsUsage: "",
	Action:    configure,
}

func configure(c *cli.Context) {
	cfg := ask()
	config.Set(cfg)
	fmt.Printf("Configuration saved to %q\n", config.ConfigFile)
}

func ask() *config.Config {
	cfg := config.Get()
	if cfg.AuthSVC == "" {
		cfg.AuthSVC = config.DefaultAuthSVC
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("username: ")
	username, _ := reader.ReadString('\n')

	fmt.Printf("password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println()

	fmt.Printf("authsvc: ")
	authenticationServiceBaseURL, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	authenticationServiceBaseURL = strings.TrimSpace(authenticationServiceBaseURL)

	if username != "" {
		cfg.Username = username
	}
	if password != "" {
		cfg.Password = password
	}
	if authenticationServiceBaseURL != "" {
		cfg.AuthSVC = authenticationServiceBaseURL
	}
	return cfg
}
