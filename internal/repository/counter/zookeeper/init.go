package zookeeper

import (
	"fmt"
	counterModel "github.com/aryahmph/url-shortener/internal/model/counter"
	counterRepo "github.com/aryahmph/url-shortener/internal/repository/counter"
	"github.com/go-zookeeper/zk"
	"strconv"
	"sync"
)

type zookeeperCounterRepository struct {
	conn        *zk.Conn
	counter     counterModel.Counter
	counterLock sync.RWMutex
}

func NewZookeeperCounterRepository(conn *zk.Conn) counterRepo.Repository {
	counter := counterModel.Counter{}

	exists, _, err := conn.Exists(counterModel.CounterPath)
	if err != nil {
		panic(err)
	}
	if !exists {
		// Counter range start from 1
		// then become 1 - 1_000_000
		counter.Start = 1
		counter.Current = 1
		counter.End = counterModel.CounterRange
		counter.Version = 0

		flags := int32(0) // No special flags
		_, err = conn.Create(counterModel.CounterPath, []byte(strconv.Itoa(int(counter.End))), flags, zk.WorldACL(zk.PermAll))
		if err != nil {
			fmt.Println("create", err)
			panic(err)
		}
	} else {
		b, stat, err := conn.Get(counterModel.CounterPath)
		if err != nil {
			fmt.Println("get", err)
			panic(err)
		}

		c, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}

		counter.End = uint64(c + counterModel.CounterRange)
		counter.Version = stat.Version + 1
		_, err = conn.Set(counterModel.CounterPath, []byte(strconv.Itoa(int(counter.End))), -1)
		if err != nil {
			fmt.Println("set", err)
			panic(err)
		}

		// Counter range start example x_000_000
		// then become x_000_001 - x_000_000
		counter.Start = uint64(c + 1)
		counter.Current = counter.Start
	}

	return &zookeeperCounterRepository{
		conn:    conn,
		counter: counter,
	}
}

func (r *zookeeperCounterRepository) GetCurrentCounter() (uint64, error) {
	r.counterLock.Lock()
	defer r.counterLock.Unlock()

	value := r.counter.Current
	if r.counter.Current > r.counter.End {
		b, stat, err := r.conn.Get(counterModel.CounterPath)
		if err != nil {
			return 0, nil
		}

		c, err := strconv.Atoi(string(b))
		if err != nil {
			return 0, nil
		}

		r.counter.End = uint64(c + counterModel.CounterRange)
		r.counter.Version = stat.Version + 1
		_, err = r.conn.Set(counterModel.CounterPath, []byte(strconv.Itoa(int(r.counter.End))), -1)
		if err != nil {
			return 0, nil
		}

		// Counter range start example x_000_000
		// then become x_000_001 - x_000_000
		r.counter.Start = uint64(c + 1)
		r.counter.Current = r.counter.Start
		return 0, nil
	} else {
		r.counter.Current += 1
	}
	return value, nil
}
