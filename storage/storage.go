package storage

import "github.com/boltdb/bolt"

type Storage interface {
	Get(key []byte) ([]byte, error)
	Store(key []byte, data []byte) error
	Close() error
}

type boltStorage struct {
	db *bolt.DB
}

func NewStorage(fname string) (Storage, error) {
	db, err := bolt.Open(fname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &boltStorage{db: db}, nil
}

func (s *boltStorage) Get(key []byte) ([]byte, error) {
	var res []byte
	if err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("contacts"))
		res = b.Get(key)
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *boltStorage) Store(key []byte, data []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("contacts"))
		if err != nil {
			return err
		}
		return b.Put(key, data)
	})
}

func (s *boltStorage) Close() error {
	return s.db.Close()
}
