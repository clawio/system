package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/clawio/system/cmd/clawioclient/config"
	"github.com/clawio/system/pb"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var AuthenticateCommand = cli.Command{
	Name:  "authenticate",
	Usage: "Authenticate user with username and password",
	ArgsUsage: "Usage: authenticate USERNAME PASSWORD",
	Description: `
`,
	Action:    authenticate,
}

func authenticate(c *cli.Context) {
	if c.NArg() != 2 {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	con, err := grpc.Dial("localhost:1502", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	username := c.Args().First()
	password := c.Args().Get(1)
	creds := &pb.Credentials{}
	creds.Value = fmt.Sprintf("%s:%s", username, password)
	client := pb.NewAuthClient(con)
	res, err := client.Authenticate(context.Background(), creds)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	config.SetToken(res.Value)
	fmt.Printf("Authentication ok. Credentials stored in %q\n", config.CredentialsFile)
}
