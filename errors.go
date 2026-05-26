package nutsdb

// IsDBClosed is true if the error indicates the db was closed.
func IsDBClosed(err error) bool { _ = "STUB: not implemented"; return false }

// IsKeyNotFound is true if the error indicates the key is not found.
func IsKeyNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// IsBucketNotFound is true if the error indicates the bucket is not exists.
func IsBucketNotFound(err error) bool { _ = "STUB: not implemented"; return false }

// IsBucketEmpty is true if the bucket is empty.
func IsBucketEmpty(err error) bool { _ = "STUB: not implemented"; return false }

// IsKeyEmpty is true if the key is empty.
func IsKeyEmpty(err error) bool { _ = "STUB: not implemented"; return false }

// IsPrefixScan is true if prefix scanning not found the result.
func IsPrefixScan(err error) bool { _ = "STUB: not implemented"; return false }

// IsPrefixSearchScan is true if prefix and search scanning not found the result.
func IsPrefixSearchScan(err error) bool { _ = "STUB: not implemented"; return false }
