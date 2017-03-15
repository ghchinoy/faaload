# Airport sequential GET

Simplistic sequential GET calls against a randomized list of airports.


## Usage

Options

* `FAA_API` - endpoint; defaults to http://services.faa.gov/airport/status
* `FAA_MAX` - max calls to make; defaults to 50
* `FAA_IATA` - comma-separated list of airport codes to use, defaults to a set list, see below

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

## FAA_IATA

Choose your own FAA airport codes to call.

The default list is: JFK,ORD,MDW,LGA,LAX,LGB,SEA,ROC,DEN,COS,DAL,HOU

Example:

```
$ FAA_IATA=DEN,ORD faaload
3
200 http://services.faa.gov/airport/status/DEN
200 http://services.faa.gov/airport/status/ORD
200 http://services.faa.gov/airport/status/ORD
```