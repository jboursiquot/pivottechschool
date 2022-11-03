# Challenge Summary
We have a monitoring service which regularly pings an inventory management service and records metrics that indicate the health of that service. The inventory management service responds with the following fields:

- `timestamp`: the unix time when pinging the service, in seconds
- `version`: the git SHA of the code run by the micro service
- `query_time`: how long it took the micro service to generate its response, in nanoseconds

## Assumptions

- The monitoring service pings the inventory service every hour on the hour
- All event responses are stored in a file called `events.json`
- Average query time should be formatted to two decimal place or else the tests won't pass

## Sample Data

```json
[
    {
        "timestamp": 1536051600,
        "version": "356a192b7913b04c54574d18c28d46e6395428ab",
        "query_time": 189
    },
    {
        "timestamp": 1536832800,
        "version": "77de68daecd823babbb58edb1c8e14d7106e83bb",
        "query_time": 124
    }
]
```

## Tasks

1. Find the minimum, average and maximum query times by version.
2. Find the best and worst performing releases.
3. Using the health data, reconstruct the release history of the service.

## Instructor's Recommendations

The following are your instructor's recommendations on what you need to know and be comfortable with as concepts in order to successfully complete this challenge:

- Know how to work with the `os` package to open files
- Know how to use with `json` package to unmarshal your file contents into custom types
- Know how to work with slices
- You may need/want to work with the `time` package