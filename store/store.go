package store

import (
	"errors"
	"fmt"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"strings"
)

var p *pool.Pool

func df(network, addr string) (*redis.Client, error) {
	client, err := redis.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	if err = client.Cmd("AUTH", "$V|>3|--p4//0|>D").Err; err != nil {
		client.Close()
		return nil, err
	}
	return client, nil
}

func init() {
	var err error
	// Establish a pool of 10 connections to the Redis server listening on
	// port 6379 of the local machine.
	p, err = pool.NewCustom("tcp", "store:6379", 10, df)
	if err != nil {
		log.Panic(err)
	}
}

// Get represents the accessor of redis' pool.
func Get() *pool.Pool {
	return p
}

func execute(req string) (*redis.Resp, error) {
	conn, err := p.Get()
	if err != nil {
		return nil, err
	}
	defer p.Put(conn)
	f := strings.Fields(req)
	var fLen = len(f)
	if fLen-1 <= 0 {
		return nil, errors.New("Bad request")
	}
	s := make([]interface{}, fLen-1)
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
	for _, e := range tmp {
		s, err := e.Str()
		if err == nil {
			arr = append(arr, s)
		}
	}
	return arr
}

// GetAllKeys returns all keys from redis' pool.
func GetAllKeys() []string {
	reply, err := execute("KEYS *")
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return respToArray(reply)
}

// Exec send a req redis request to redis and may return an error from redis.
func Exec(req string) error {
	fmt.Print("exec : `", req, "`\n")
	_, err := execute(req)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
