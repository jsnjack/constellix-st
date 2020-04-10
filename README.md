constellix-st
========

Generates constellix security token

# How to use?
```
$ constellix-st -h
Usage of constellix-st:
  -c	Print as a curl header
  -d	Output debug information
```
Example with curl:
```bash
export CONSTELLIX_API_KEY=your_key
export CONSTELLIX_SECRET_KEY=your_secret
curl -H "Content-Type: application/json" -H "`constellix-st -c`" "https://api.dns.constellix.com/v1/domains" | jq .
```

# How to install?
 Download a binary file from Releases page
