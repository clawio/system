package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/clawio/system/cmd/clawioclient/config"
	"github.com/clawio/system/pb"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var WhoamiCommand = cli.Command{
	Name:  "whoami",
	Usage: "Whoami shows the current logged in user",
	ArgsUsage: "Usage: whoami",
	Description: `
`,
	Action: whoami,
}

func whoami(c *cli.Context) {
	ctx := config.GetTokenContext()
	con, err := grpc.Dial(config.Get().AuthSVC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	client := pb.NewAuthClient(con)
	empty := &pb.Empty{}
	res, err := client.Whoami(ctx, empty)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("AccountId: %s\n", res.AccountId)
	fmt.Printf("DisplayName: %s\n", res.DisplayName)
	fmt.Printf("Email: %s\n", res.Email)
	fmt.Printf("Opaque: %s\n", res.Opaque)
}
