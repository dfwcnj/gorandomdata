package randomdata

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var alphanum = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSeq(n int)
// generate a random string length n with lower, upper case alphanum and digits
func randSeq(slen int, rlen bool) string {
	b := make([]rune, slen)
	ll := len(alphanum)
	for i := range b {
		b[i] = alphanum[rand.Intn(ll)]
	}
	if rlen == true {
		rl := rand.Intn(slen)
		if rl != 0 {
			b = b[:rl]
		}
	}
	return string(b)
}

// randomstrings(n int64, slen int)
// generate n random strings with length slen
// return a slice containing the strings
func Randomstrings(n int64, slen int, rlen bool) []string {
	//log.Print("rlen ", rlen)
	ssl := make([]string, 0)
	for _ = range n {
		ssl = append(ssl, randSeq(slen, rlen))
	}
	return ssl
}

// randomints(ņ int)
// generate n random int64 values
// return a slice containing the int64 values
func Randomuints(n int64) []uint64 {
	usl := make([]uint64, 0)
	for _ = range n {
		usl = append(usl, rand.Uint64())
	}
	return usl
}

// randomdates(n int64, format string)
// generate n random dates with format
// return a slice containing the random date strings
func Randomdates(n int64, format string, emit bool) []string {
	now := time.Now().Unix()
	var mod = int64(now)
	var s string
	dsl := make([]string, 0)
	if emit == true {
		fp := os.Stdout
		wp := bufio.NewWriterSize(fp, 1<<22)
		var s string
		for _ = range n {
			ri := rand.Int63() % mod
			tm := time.Unix(int64(ri), int64(0))

			switch format {
			case "DateTime":
				s = fmt.Sprintln(tm.Format(time.DateTime))
			case "Layout":
				s = fmt.Sprintln(tm.Format(time.Layout))
			case "RubyDate":
				s = fmt.Sprintln(tm.Format(time.RubyDate))
			case "UnixDate":
				s = fmt.Sprintln(tm.Format(time.UnixDate))
			case "RFC3339":
				s = fmt.Sprintln(tm.Format(time.RFC3339))
			default:
				s = fmt.Sprintln(tm)
			}
			wp.WriteString(s)
		}
	} else {
		for _ = range n {
			ri := rand.Int63() % mod
			tm := time.Unix(int64(ri), int64(0))

			switch format {
			case "DateTime":
				s = fmt.Sprint(tm.Format(time.DateTime))
			case "Layout":
				s = fmt.Sprint(tm.Format(time.Layout))
			case "RubyDate":
				s = fmt.Sprint(tm.Format(time.RubyDate))
			case "UnixDate":
				s = fmt.Sprint(tm.Format(time.UnixDate))
			case "RFC3339":
				s = fmt.Sprint(tm.Format(time.RFC3339))
			default:
				s = fmt.Sprint(tm)
			}
			dsl = append(dsl, s)
		}
	}
	return dsl
}
