# Simple TOTP generator

Reads a secret key, encoded with base32, and outputs

```
% cat /path/to/secret/key.txt | totp
```

For example:

```
echo | ./totp
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
