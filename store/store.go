package store

import "context"

type Store interface {
	// Put writes to a transaction, which might be flushed from time to time. Call FlushPuts() to ensure all Put entries are properly written to the database.
	Put(ctx context.Context, key, value []byte) (err error)
	// FlushPuts takes any pending writes (calls to Put()), and flushes them.
	FlushPuts(ctx context.Context) (err error)

	// Get a given key.  Returns `kdb.ErrNotFound` if not found.
	Get(ctx context.Context, key []byte) (value []byte, err error)
	// BatchGet get a batch of keys.  Returns `kdb.ErrNotFound` the first time a key is not found: not finding a key is fatal and interrupts the result set from being fetched completely.  BatchGet guarantees that Iterator return results in the exact same order as keys
	BatchGet(ctx context.Context, keys [][]byte) *Iterator

	Prefix(ctx context.Context, prefix []byte, limit int, options ...ReadOption) *Iterator

	// Delete a given key.  Returns `kdb.ErrNotFound` if not found.
	Delete(ctx context.Context, key []byte) (err error)

	BatchDelete(ctx context.Context, keys [][]byte) (err error)

	// Close the underlying store engine and clear up any resources currently hold
	// by this instance.
	//
	// Once this instance's `Close` method has been called, it's assumed to be terminated
	// and cannot be reliably used to perform read/write operation on the backing engine.
	Close() error
}
