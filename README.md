constellix-st
========

Generates constellix security token

# How to use it with curl
```bash
curl -H "Content-Type: application/json" -H "`constellix-st -c`" "https://api.dns.constellix.com/v1/domains" | jq .
```
