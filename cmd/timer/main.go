package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"timer/internal/timer"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "timer",
		Action: func(cCtx *cli.Context) error {
			minutes, err := parseMinutes(cCtx.Args())
			if err != nil {
				return err
			}

			if err := timer.New(minutes); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseMinutes(args cli.Args) (int, error) {
	if args.Len() != 1 {
		return 0, fmt.Errorf("expected 1 argument, got %v", args.Len())
	}

	minutes, err := strconv.Atoi(args.First())
	if err != nil {
		return 0, fmt.Errorf("invalid argument: %v", args.First())
	}

	return minutes, nil
}
