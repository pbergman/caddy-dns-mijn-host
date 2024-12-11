Mijn Host module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with [mijn host](https://mijn.host/).
It makes use of [pbergman/libdns-mijn-host](https://github.com/pbergman/libdns-mijn-host)

## Install caddy module

use xcaddy to build a version with this module

```
xcaddy build --with github.com/pbergman/caddy-dns-mijn-host
```


## Caddy module name

```
dns.providers.mijn-host
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "mijn-host",
				"api_key": "YOUR_API_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns mijn-host <api_key>
}
```

```
# one site
tls {
	dns mijn-host <api_key>
}
```

or alternatively:


```
tls {
	dns transip {
		account_name <accountName> 
		private_key_path <privateKeyPath>
	}
}
```