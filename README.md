# netsumm
readsa json file and runs a series of network tests. save the output during normal operations as a baseline
```bash
echo '{
"targets": [
{
"type": "DNS Lookup",
"destination": "8.8.8.8",
"data": "www.google.com"
},
{
"type": "TCP Connection",
"destination": "www.google.com",
"data": "443"
}
],
"source": "",
"iterations": 20
}' > netsumm.json

netsumm
```

```bash


```