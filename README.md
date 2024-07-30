# Simple TOTP generator

Reads a secret key, base32-encoded, and outputs TOTP (Time-based One Time Password) as defined in https://datatracker.ietf.org/doc/html/rfc6238

Typical usage:
```
% cat /path/to/secret/key.txt | totp
```

For example:
```
% cat example_key.txt | totp
012359%
```

Supported flags are (all optional):
```
% totp
Usage of totp:
  -digits int
    	number of digits (default 6)
  -t string
    	timestamp (T) as Unix epoch (default current-timestmp)
  -t0 int
    	start timestamp (T0) as Unix epoch (default 0)
  -x int
    	time step in seconds (X) (default 30)
```


## SwiftBar integration

The `totp-swiftbar.sh` is an example script of how to use this tool with the awesome SwiftBar


## License

Public domain
