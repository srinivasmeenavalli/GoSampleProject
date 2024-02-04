package main

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type InputParams struct {
	EventSourceName  string
	EmittedEventType string
}

type Function struct {
	params *InputParams
}

func NewFunction(params *InputParams) *Function {
	return &Function{
		params: params,
	}
}

func (fn *Function) Run() error {
	// log.Info("starting function store")

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		return errors.Wrap(err,
			"unable to create a cloud events client")
	}

	// blocking operation
	return c.StartReceiver(context.Background(), fn.Receive)
}

func (fn *Function) Receive(ctx context.Context,
	event cloudevents.Event) (*event.Event, protocol.Result) {

	log.Debugf("event received: %+v", event.Type())
	log.Debugf("with content: %+v", string(event.Data()))

	newEvent := cloudevents.NewEvent()
	newEvent.SetID(uuid.New().String())
	newEvent.SetSource(fn.params.EventSourceName)
	newEvent.SetType(fn.params.EmittedEventType)
	if err := newEvent.SetData(cloudevents.ApplicationJSON,
		event.Data()); err != nil {
		log.Errorf("unable to set data for the event: %s",
			err.Error())
		return nil, protocol.ResultNACK
	}

	return &newEvent, protocol.ResultACK
}
