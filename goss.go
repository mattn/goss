package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"launchpad.net/goamz/aws"
	"os"
	"regexp"
)

var pat = regexp.MustCompile("^sakura://([^/]+)/(.*)")

func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("SAKURA_STORAGE_USER_ID"),
		SecretKey: os.Getenv("SAKURA_STORAGE_ACCESS_TOKEN"),
	}
	region := aws.Region{
		"sakura-storage",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"",
		true,
		true,
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
		"https://b.storage.sakura.ad.jp",
	}

	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "sakura storage utility",
	}
	command.Subcommands = []*commander.Command{
		make_cmd_ls(auth, region),
		make_cmd_put(auth, region),
		make_cmd_get(auth, region),
		make_cmd_cat(auth, region),
		make_cmd_del(auth, region),
	}

	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
