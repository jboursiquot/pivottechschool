# Capstone Project

## Summary

The capstone project outlined here will be to create a CLI tool (named `ziplook`) that, given a postal code, opens a browser that takes the user to Google Maps and drops a pin at the latitude and longitude of said postal code. The project will involve creating a API server for the CLI to interact with. The API will be hosted online using a third party Platform as a Service (PaaS) provider (Zeet) and be reacheable by issuing HTTP GET requests (no authorization required).

I will demonstrate the following:

1. Building a REST API in Go
2. Integrating with a remote API (https://api.zippopotam.us)
3. Using database behind the scenes to cache information that has already been fetched from a remote API (place information does not change)
4. Deploying my REST API on a third-party PaaS, Zeet (https://zeet.co).
5. Building a client that can interact with my remote REST API.
6. Building a CLI app that exposes a user-friendly way of leveraging my REST API server and client.
7. Good coding practices (e.g. handling errors properly, returning the proper error message to the user, etc)
8. Testing of HTTP handlers
9. Testing of HTTP client
10. Proper usage of Go interfaces, especially during testing

## User Stories

### As a user, I would like to invoke the CLI app with a postal code and have a browser window open at a Google Maps page showing me where in the world the zip code is located.

**Acceptance Criteria**

Given a postal code, when it is valid, I expect my default browser to open with a Google Map URL that drops a pin at the lat and long of the postal code.

Example:
```
$ ziplook show 21042
```

Default browser window opens at the following URL where `<lat>` is `39.2726` and `<long>` is `-76.8614` where these values are obtained by interacting with my API hosted on Zeet.

```
https://www.google.com/maps/search/?api=1&query=<lat>,<lng>
```

> https://www.google.com/maps/search/?api=1&query=39.2726,-76.8614
 
Given a postal code, when it is invalid, I expect the CLI app to indicate that the postal code was not found so that I may try again.

### As a user, I would like to invoke the CLI app with the option of getting the raw place data so that I can see what information is returned from the API

**Acceptance Criteria**

Given that I specify a flag when invoking the CLI app, when the postal code is valid, I will receive the raw JSON place data printed in my console/shell.

Example:
```
$ ziplook -raw 21042

{"post code": "21042", "country": "United States", "country abbreviation": "US", "places": [{"place name": "Ellicott City", "longitude": "-76.8614", "state": "Maryland", "state abbreviation": "MD", "latitude": "39.2726"}]}
```

Given that I specify a flag when invoking the CLI app, when the postal code is invalid, I will receive an error message indicating as such so that I may try again.

Example:

```
$ ziplook -show abc123
Error: "abc123" is not a valid postal code.
```

### As a user, I would like the CLI app to provide usage help when I invoke it without a command so that I may learn how to use it properly.

**Acceptance Criteria**

Given the CLI app is invoked, when no command is specified, usage help will be provided.

Example:

```
$ ziplook
Usage of ziplook:
  -show
      If valid, opens a Google Map URL in your default browser at the postal code's latitude and longitude.
  -raw
      If valid, shows the raw JSON data for the place found for that postal code.
  -help
      Shows this usage information.
```