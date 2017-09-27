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

	app.Name = "Flower"
	app.HelpName = "flower"

	app.Commands = []cli.Command{
		{
			Name: "run",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "code",
					Usage: "give a code name for store state",
				},
				cli.StringFlag{
					Name:  "cwd",
					Usage: "current work dir",
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

func run(ctx *cli.Context) (err error) {
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

	deployConf := config.NewConfig(
		config.ConfigFile(scriptConfig),
	)

	flowConf := config.NewConfig(
		config.ConfigFile(flowConfig),
	)

	redisAddr := deployConf.GetString("deployer.redis.address", "127.0.0.1:6379")
	redisDb := deployConf.GetInt32("deployer.redis.db", 0)
	redisPassword := deployConf.GetString("deployer.redis.password", "")

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       int(redisDb),
		})

	redisClient.Ping()

	f, err := flow.NewFlow(
		flowConf.GetString("name"),
		flow.ConfigFile(flowConfig),
	)

	if err != nil {
		return
	}

	task := f.NewTask()

	task.Context().
		Set("config", deployConf).
		Set("log", logrus.StandardLogger()).
		Set("redis", redisClient).
		Set("code", code)

	err = task.Run()

	return
}
