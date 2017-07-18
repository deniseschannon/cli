package handlers

import (
	"bytes"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/rancher-compose-executor/project"
	"github.com/rancher/rancher-compose-executor/project/events"
)

func NewListenLogger(logger *logrus.Entry, p *project.Project) chan<- events.Event {
	listenChan := make(chan events.Event)
	go func() {
		for event := range listenChan {
			buffer := bytes.NewBuffer(nil)
			if event.Data != nil {
				for k, v := range event.Data {
					if buffer.Len() > 0 {
						buffer.WriteString(", ")
					}
					buffer.WriteString(k)
					buffer.WriteString("=")
					buffer.WriteString(v)
				}
			}

			logger.Infof("[%s:%s]: %s %s", p.Name, event.ServiceName, event.EventType, buffer.Bytes())
		}
	}()
	return listenChan
}
