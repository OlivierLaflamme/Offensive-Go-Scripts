# findIDHash
Go script that identifies a hash type.


findIDHash is written in Go. 
It takes a hash, calculates the length, and compares the length against the [hashcat example hashes](https://hashcat.net/wiki/doku.php?id=example_hashes)

*   Hash is provided in a file to clean up whitespace and prevent interpolation.
*   Can only evaluate one hash at a time. So there should be one hash per file. 
*   Remove any usernames, domain information, etc. from the raw hash provided. This tool compares hashes by length so usernames with varying length sizes will not match the examples in hashcathash.csv
*   The next upgrade will address some of these limitations.

Usage
------------
go run findIDHash.go -file <<Filename>> [default: hash.txt] 

Example: go run findIDHash.go -file hash.txt

or build

*   env GOOS=windows GOARCH=386 go build -v findIDHashgo
*   env GOOS=linux GOARCH=amd64 go build -v findIDHash.go

hashcathash.csv file provided in this repo is required


|Parameter     |Description  |
|-----------|-------------------------------------------------------|
|-file      |inpute file with hash <<Filename>> [default: hash.txt] |
|-h, --help |show help message and exit                             |



Example
----------

    $ go run findIDHash.go -file hash.txt
	findIDHash v1          Olivier Laflamme
	=========================================

	NetNTLMv1 / NetNTLMv1+ESS [-m 5500]


    $ cat hash.txt
    338d08f8e26de93300000000000000000000000000000000:9526fb8c23a90751cdd619b6cea564742e1e4bf33006ba41:cb8086049ec4736c
    

Resources
---------

-  https://hashcat.net/wiki/doku.php?id=example_hashes
