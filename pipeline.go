
package go4redis

import (
  "fmt"
  "errors"
)

func (c* Client) StartPipeline() {
  c.pipelineMode = true
  c.pipelineBuffer.Reset()
  c.pipelineChan = make(chan *Response)
}

func (c* Client) AddToPipeline(command string) error {
  if c.pipelineMode == true {
    c.pipelineBuffer.WriteString(command)
    return nil
  } else {
    return errors.New("Call StartPipeline() before calling AddToPipeline(string)")
  }
}

func (c *Client) ExecPipeline() <-chan *Response {
  fmt.Fprintf(c.conn, c.pipelineBuffer.String())
  return c.pipelineChan
}

func (c* Client) EndPipeline() {
  c.pipelineMode = false
  c.pipelineBuffer.Reset()
  close(c.pipelineChan)
}
