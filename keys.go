package go4redis

// DEL key [key ...]
// Delete a key
// DUMP key
// Return a serialized version of the value stored at the specified key.
// EXISTS key [key ...]
// Determine if a key exists
// EXPIRE key seconds
// Set a key's time to live in seconds
// EXPIREAT key timestamp
// Set the expiration for a key as a UNIX timestamp
// KEYS pattern
// Find all keys matching the given pattern
// MIGRATE host port key destination-db timeout [COPY] [REPLACE]
// Atomically transfer a key from a Redis instance to another one.
// MOVE key db
// Move a key to another database
// OBJECT subcommand [arguments [arguments ...]]
// Inspect the internals of Redis objects
// PERSIST key
// Remove the expiration from a key
// PEXPIRE key milliseconds
// Set a key's time to live in milliseconds
// PEXPIREAT key milliseconds-timestamp
// Set the expiration for a key as a UNIX timestamp specified in milliseconds
// PTTL key
// Get the time to live for a key in milliseconds
// RANDOMKEY
// Return a random key from the keyspace
// RENAME key newkey
// Rename a key
// RENAMENX key newkey
// Rename a key, only if the new key does not exist
// RESTORE key ttl serialized-value [REPLACE]
// Create a key using the provided serialized value, previously obtained using DUMP.
// SORT key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]
// Sort the elements in a list, set or sorted set
// TTL key
// Get the time to live for a key
// TYPE key
// Determine the type stored at key
// WAIT numslaves timeout
// Wait for the synchronous replication of all the write commands sent in the context of the current connection
// SCAN cursor [MATCH pattern] [COUNT count]
// Incrementally iterate the keys space
