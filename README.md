# Airport sequential GET

Simplistic sequential get against a random list of airports.


## Usage

Options

* `FAA_API` - endpoint
* `FAA_MAX` - max calls to make

These can be combined, example:

```
$ FAA_MAX=10 FAA_API=https://nd.akana.dev:9982/api10973live faaload
4
200 https://nd.akana.dev:9982/api10973live/LGB
200 https://nd.akana.dev:9982/api10973live/LGB
200 https://nd.akana.dev:9982/api10973live/SEA
200 https://nd.akana.dev:9982/api10973live/DAL
```

## FAA_API
Set an environment variable for the base API, the airport code will be appended to the base API, otherwise http://services.faa.gov/airport/status will be used.

example:

```
$ FAA_API=https://nd.akana.dev:9982/api10973live faaload
```

## FAA_MAX

One can also set a maximum amount of calls to make, using the env var `FAA_MAX`.

example:

```
$ FAA_MAX=20 faaload
```