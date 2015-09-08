package go4redis

// SADD key member [member ...]
// Add one or more members to a set
// SCARD key
// Get the number of members in a set
// SDIFF key [key ...]
// Subtract multiple sets
// SDIFFSTORE destination key [key ...]
// Subtract multiple sets and store the resulting set in a key
// SINTER key [key ...]
// Intersect multiple sets
// SINTERSTORE destination key [key ...]
// Intersect multiple sets and store the resulting set in a key
// SISMEMBER key member
// Determine if a given value is a member of a set
// SMEMBERS key
// Get all the members in a set
// SMOVE source destination member
// Move a member from one set to another
// SPOP key [count]
// Remove and return one or multiple random members from a set
// SRANDMEMBER key [count]
// Get one or multiple random members from a set
// SREM key member [member ...]
// Remove one or more members from a set
// SUNION key [key ...]
// Add multiple sets
// SUNIONSTORE destination key [key ...]
// Add multiple sets and store the resulting set in a key
// SSCAN key cursor [MATCH pattern] [COUNT count]
// Incrementally iterate Set elements
