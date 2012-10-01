package kdfutil

import (
	"hash"
	"time"
	"code.google.com/p/go.crypto/pbkdf2"
)

const maxInt = int(^uint(0) >> 1) 
 
// BUG: clocks are not monotonic.

// CalibratePBKDF2 returns the estimated number of iterations required to 
// derive key of length keyLen with hash function h in the given time. 
func CalibratePBKDF2(dur time.Duration, h func() hash.Hash, keyLen int) int { 
	// XXX uses block size as input, maybe provide password and salt len?
        block := make([]byte, h().BlockSize()) 
        start := time.Now() 
        var ns int64 
        iter := 0 
        for { 
                pbkdf2.Key(block, block, 1000, keyLen, h) 
                iter += 1000 
                ns = time.Since(start).Nanoseconds() 
                if ns > 0 { 
                        break 
                } 
        } 
        if dur.Nanoseconds()/ns > int64(maxInt/iter) { 
                return maxInt 
        } 
        return iter * int(dur.Nanoseconds()/ns) 
} 
