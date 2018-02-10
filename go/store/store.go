package store

import (
  "errors"
  "fmt"
  "log"
  "strings"
  "github.com/mediocregopher/radix.v2/redis"
  "github.com/mediocregopher/radix.v2/pool"
)

var p *pool.Pool

func init() {
    var err error
    // Establish a pool of 10 connections to the Redis server listening on
    // port 6379 of the local machine.
    //p, err = pool.New("tcp", "localhost:6379", 10)
    p, err = pool.New("tcp", "store:6379", 10)
    if err != nil {
        log.Panic(err)
    }
}

func Get() *pool.Pool {
  return p;
}

func execute(req string) (*redis.Resp, error) {
  conn, err := p.Get()
  if err != nil {
    return nil, err
  }
  defer p.Put(conn)
  f := strings.Fields(req)
  var fLen = len(f)
  if fLen - 1 <= 0 {
    return nil, errors.New("Bad request")
  }
  s := make([]interface{}, fLen - 1)
  for i, v := range f[1:] {
    s[i] = v
  }
  res := conn.Cmd(f[0], s)
  if err != nil {
    return nil, err
  }
  return res, nil
}

func respToArray(r *redis.Resp) []string {
  var arr []string

  tmp, err := r.Array()
  if err != nil {
    return arr
  }
  for _, e:= range tmp {
    s, err := e.Str()
    if err == nil {
      arr = append(arr, s)
    }
  }
  return arr
}

func GetAllKeys() []string {
  reply, err := execute("KEYS *")
  if err != nil {
    fmt.Print(err)
    return nil
  } else {
    return respToArray(reply)
  }
}

func Exec(req string) error {
  fmt.Print("exec : `", req, "`\n")
  _, err := execute(req)
  if err != nil {
    fmt.Print(err)
    return err
  }
  return nil
}
