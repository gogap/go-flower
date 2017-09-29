package main

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/gogap/flow"
	"github.com/urfave/cli"

	"github.com/Sirupsen/logrus"
	"github.com/go-redis/redis"
	"github.com/gogap/config"
)

type NullWriter struct {
}

func (p *NullWriter) Write(v []byte) (int, error) {
	return len(v), nil
}

func init() {
	log.SetOutput(&NullWriter{})
}

func main() {

	app := cli.NewApp()

	app.Name = "Go Flower"
	app.HelpName = "go-flower"
	app.Usage = "A tiny script runner"

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run flow",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config",
					Usage: "the flow configuration",
					Value: "flow.conf",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "give a code name for store state",
				},
				cli.StringFlag{
					Name:  "cwd",
					Usage: "change current work dir",
				},
				cli.StringSliceFlag{
					Name:  "args",
					Usage: "passthrough argement to context as key:value format e.g.: --args group:gogap --args debug --args user:zeal",
				},
				cli.StringFlag{
					Name:  "loglevel",
					Usage: "change log level: debug, info, warn, error, fatal, panic",
					Value: "info",
				},
			},
			Action: run,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Errorln(err)
	}

	return
}

func fileExist(fileName string) bool {
	fi, err := os.Stat(fileName)
	if err != nil {
		return false
	}

	if fi.IsDir() {
		return false
	}

	return true
}

func changeLogLevel(ctx *cli.Context) (err error) {

	strLvl := ctx.String("loglevel")

	lvl, err := logrus.ParseLevel(strLvl)
	if err != nil {
		return
	}

	logrus.SetLevel(lvl)
	return
}

func run(ctx *cli.Context) (err error) {

	err = changeLogLevel(ctx)
	if err != nil {
		return
	}

	code := ctx.String("code")

	if len(code) == 0 {
		cli.ShowSubcommandHelp(ctx)
		err = errors.New("code name is empty")
		return
	}

	cwd := ctx.String("cwd")

	if len(cwd) == 0 {
		cwd, _ = os.Getwd()
	} else {
		err = os.Chdir(cwd)
	}

	if err != nil {
		return
	}

	configName := path.Base(cwd) + ".conf"

	flowConfig := "flow.conf"
	scriptConfig := configName

	if !fileExist(flowConfig) {
		cli.ShowSubcommandHelp(ctx)
		err = errors.New("flow.conf not exist")
		return
	}

	if !fileExist(scriptConfig) {
		cli.ShowSubcommandHelp(ctx)
		err = errors.New(scriptConfig + " not exist")
		return
	}

	flowConf := config.NewConfig(
		config.ConfigFile(flowConfig),
	)

	scriptConf := config.NewConfig(
		config.ConfigFile(scriptConfig),
	)

	redisAddr := flowConf.GetString("redis.address", "127.0.0.1:6379")
	redisDb := flowConf.GetInt32("redis.db", 0)
	redisPassword := flowConf.GetString("redis.password", "")

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       int(redisDb),
		})

	if redisClient.Ping().Val() != "PONG" {
		logrus.WithField("ADDRESS", redisAddr).WithField("DB", redisDb).Warnln("Redis are not available.")
	}

	f, err := flow.NewFlow(
		flowConf.GetString("name"),
		flow.Config(
			flowConf.GetConfig("flow"),
		),
	)

	if err != nil {
		return
	}

	task := f.NewTask()

	strArgs := ctx.StringSlice("args")

	args := NewArgs(strArgs...)

	task.Context().
		Set("CONFIG", scriptConf).
		Set("LOG", logrus.StandardLogger()).
		Set("REDIS", redisClient).
		Set("CODE", code).
		Set("ARGS", args)

	err = task.Run()

	return
}
