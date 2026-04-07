package utils

import (
	"time"
)

// NOTE: stop a RepeatingTimer obj can using
// Stop() or set t.Enable = false

// example
// gTick = utils.NewRepeatingTimer(1, call)
// gTick.Start()
// gTick.Stop()  // clean

type CallbackFunc func() error

type RepeatingTimer struct {
  Ticker *time.Ticker
  Runner CallbackFunc
  Enable bool
}

func NewRepeatingTimer(interval int, f CallbackFunc) *RepeatingTimer {
  return &RepeatingTimer{
    Ticker: time.NewTicker(time.Duration(interval) * time.Second),
    Runner: f,
    Enable: true,
  }
}

func (t *RepeatingTimer) Start()  {
  t.Enable = true

  go func() {
    for {
      if ! t.Enable {
        t.Ticker.Stop()
        break
      }
      select {
      case <-t.Ticker.C:
        t.Runner()
      }
    }
  }()
}

func (t *RepeatingTimer) Stop()  {
  t.Ticker.Stop()
  t.Enable = false
}
