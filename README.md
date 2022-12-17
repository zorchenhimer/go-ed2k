# ED2K

This package implements the ed2k hashing algorithm.  There are two different
implementations that are mostly the same but differ on an edge case.  Both of
these are provided by this package.

## The Hash

A hash is calculated by first reading the input in 9728000 byte chunks.  Each
chunk is then hashed with the MD4 algorithm.  Once all chunks have been hashed,
the list of hashes for each chunk is hashed one more time, giving the final
hash.

If there is less than 9728000 bytes of data to hash, the MD4 of the data is
returned without further modification.

The two implementations differ when there are is a multiple of 9728000 bytes in
every chunk of the data.  The Red version will add a zero-length chunk to the
end and add the hash of this empty chunk to the list of hashes before computing
the final hash.  The Blue version does not do this.

The Sum() function uses the Blue method.
