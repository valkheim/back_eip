package main

import (
  "time"
  "github.com/BurntSushi/toml"
)

type duration struct {
  time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
  var err error
  d.Duration, err = time.ParseDuration(string(text))
  return err
}

type General struct {
  Version string
  Address string
  Port    int
  Timeout duration
}

type Config struct {
  Title   string
  General General
}

func (c *Config) Get(filename string) error {
  if _, err := toml.DecodeFile(filename, &c); err != nil {
    return err
  }
  return nil
}
