package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func createApp() *cli.App {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "event-source-name",
			Usage:   "event source of newly emitted event",
			Value:   "alcl-go-template-function",
			EnvVars: []string{"EVENT_SOURCE_NAME"},
		},
		&cli.StringFlag{
			Name:    "emitted-event-type",
			Usage:   "type of emitted event",
			Value:   "sample.functiontemplate.templated",
			EnvVars: []string{"EMITTED_EVENT_TYPE"},
		},
		&cli.IntFlag{
			Name:    "log-level",
			Usage:   "log-level",
			Value:   5,
			EnvVars: []string{"LOG_LEVEL"},
		},
	}

	return &cli.App{
		Name:  "alcl-go-template-function",
		Usage: "alcl-go-template-function is a knative template function that is subscribed on the TBD",
		Flags: flags,
		Action: func(c *cli.Context) error {

			log.SetFormatter(&log.JSONFormatter{})
			log.SetLevel(log.Level(c.Int("log-level")))

			eventSourceName := c.String("event-source-name")
			if eventSourceName == "" {
				log.Error("empty event-source-name parameter, you can use EVENT_SOURCE_NAME env var")
				cli.ShowSubcommandHelpAndExit(c, 6)
			}

			emittedEventType := c.String("emitted-event-type")
			if emittedEventType == "" {
				log.Error("empty emitted-event-type parameter, you can use EMITTED_EVENT_TYPE env var")
				cli.ShowSubcommandHelpAndExit(c, 7)
			}

			return runCMD(&InputParams{
				EventSourceName:  eventSourceName,
				EmittedEventType: emittedEventType,
			})
		},
	}
}

func runCMD(params *InputParams) error {
	return NewFunction(params).Run()
}
