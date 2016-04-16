package da

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/lcaballero/hitman"
)

type DataProvider func() interface{}
type Access func(interface{}) bool

// DataStore holds some data an periodically writes that to file.
type Locker struct {
	modified bool
	filename string
	lock     *sync.Mutex
	data     interface{}
}

// NewDataStore loads the db from the given file, or if the file doesn't
// exist then uses the forNew constructor, else it fills in the instance
// from forExisting instance.
func NewDataStore(
	dbname string,
	forNew DataProvider,
	forExisting DataProvider) (*Locker, error) {

	_, err := os.Stat(dbname)
	if os.IsNotExist(err) {
		return &Locker{
			filename: dbname,
			lock:     &sync.Mutex{},
			data:     forNew(),
		}, nil
	} else {
		return LoadFromFile(dbname, forExisting())
	}
}

// LoadFromFile loads the db from the given file and populates the data
// instance based on the provided interface.
func LoadFromFile(dbname string, data interface{}) (*Locker, error) {
	bits, err := ioutil.ReadFile(dbname)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bits, data)
	if err != nil {
		return nil, err
	}
	d := &Locker{
		filename: dbname,
		lock:     &sync.Mutex{},
		data:     data,
	}
	return d, nil
}

// WriteTo write the data to the given writer which can be loaded by
// LoadFromFile to recreate the DataStore from file.
func (d *Locker) WriteTo(w io.Writer) (int64, error) {
	bits, err := json.MarshalIndent(d.data, "", "  ")
	if err != nil {
		return 0, nil
	}
	n, err := w.Write(bits)
	return int64(n), err
}

func (d *Locker) Flush() (bool, error) {
	if !d.modified {
		return false, nil
	}
	log.Printf("Flushing data to file: %s\n", d.filename)
	file, err := os.Create(d.filename)
	if err != nil {
		log.Println("Unable to write file to disk", err)
		return false, err
	}
	defer file.Close()
	_, err = d.WriteTo(file)
	if err != nil {
		log.Println("Error occured while writing data", err)
		return false, err
	}

	return true, nil
}

// Implementation of the hitman.NamedTarget interface.
func (d *Locker) Name() string {
	return "DataStore"
}

// Implements the hitman.Target interface.
func (d *Locker) Start() hitman.KillChannel {
	log.Println("Starting DataStore")
	done := hitman.NewKillChannel()
	writeTic := time.NewTicker(5 * time.Second).C
	go func() {
		for {
			select {
			case cleaner := <-done:
				log.Println("Flushing to disk and shutting down data_store")
				d.Flush()
				cleaner.WaitGroup.Done()
				return

			case <-writeTic:
				log.Println("Flushing to disk")
				d.Flush()
			}
		}
	}()

	return done
}

// Provides locking around the function and access to the function
// to the underlying data.
func (d *Locker) DataStore(fn Access) {
	defer d.lock.Unlock()
	d.lock.Lock()
	d.modified = fn(d.data)
}
