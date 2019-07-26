# _Amplitude_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)


An OMG service for amplitude, it is a product analytics service that makes it easier for companies to understand user behavior, ship the right features and improve business outcomes.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Check domain availability
```coffee
godaddy checkDomainAvailability domain:'domain.com'
{"available": true,"currency": "USD","definitive": true,"domain": "domain.com"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Check domain availability
```shell
$ omg run checkDomainAvailability -a domain=<DOMAIN_NAME> -e API_KEY=<API_KEY> -e API_SECRET=<API_SECRET>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/godaddy/blob/master/LICENSE).
