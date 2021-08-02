<p align="center">
    <img src="https://user-images.githubusercontent.com/56378698/127357452-1b57af9c-be5a-42ff-aecb-bd2e2c006716.png" alt="Pioneer logo" width="200" height="200">
</p>

# Go SDK for Pioneer

[![Test](https://github.com/pioneer-io/go_sdk/actions/workflows/verify.yml/badge.svg)](https://github.com/pioneer-io/go_sdk/actions/workflows/verify.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/pioneer-io/go_sdk.svg)](https://pkg.go.dev/github.com/pioneer-io/go_sdk)
[![Go](https://img.shields.io/github/go-mod/go-version/pioneer-io/go_sdk)](https://github.com/gomods/athens)
[![Maintainability](https://api.codeclimate.com/v1/badges/1d04449ece98968a14c1/maintainability)](https://codeclimate.com/github/pioneer-io/go_sdk/maintainability)


This module is a server-side SDK for applications written in Golang using Pioneer's feature management service.

Visit the [pioneer-io/pioneer](https://github.com/pioneer-io/pioneer) repo for more or check out Pioneer's [case study](https://pioneer-io.github.io/).

## Getting started

From your Go module:

```
go get github.com/pioneer-io/go_sdk
```

To initialize a new SDK client, you'll need to provide the location of Pioneer's Scout daemon. The default port for Scout is `:3030`. The `/features` endpoint is the correct endpoint to receive feature flag updates. Once connected, communication with Scout is uni-directional. The SDK client will receive ruleset updates in real-time via SSE any time a feature flag is created/updated/deleted via the Compass dashboard.

You'll also need to provide a valid SDK key. The SDK key can be found on the Compass application's user interface via the 'Account' tab.

```Go
// Initialize an SDK client
scoutDaemon := "http://localhost:3030"
sdkKey := "a-dummy-key"
client := sdk.InitMember(scoutDaemon, sdkKey)

// connect SDK client to Scout to listen for SSE updates
client.Connect()
client.Listen()
```

### Failed SSE Connections

If the SDK fails to connect to the Scout daemon as an eventsource client, the connection attempt will be retried up to 10 times. The SDK will 'jitter' these connection attempts-- pausing for a random length of time between 1 and 10 seconds (inclusive) in between each attempt.

If the connection fails 10 times, an error will be logged to the user (example below) and the SDK will stop trying to connect.

```
2021/07/17 12:22:54 ERROR: Get "http://localhost:3030/features": dial tcp [::1]:3030: connect: connection refused
exit status 1
```

## Usage

### `client.Get(flagKey string) -> bool`

```Go
flagKey := "login_service"
client.Get(flagKey) // returns a boolean

// potential example usage
if client.Get(flagKey) {
  executeMicroservice()
} else {
  doSomethingElse()
}

```

The `clientGet()` function returns a boolean indicating whether the feature flag is currently toggled on or off.

**NOTE**: Passing in an invalid flag name results in an error being logged, as below:

```
2021/07/17 10:34:05 The flag 'non existent flag' is not in the ruleset
exit status 1
```

### `client.GetWithDefault(flagKey string, default bool) -> bool`

There may be situations in which you do not want your code to log an error, even if the flag does not exist in the ruleset. For this case, we provide the `GetWithDefault()` function. If the flag exists in the ruleset, its toggle status (`true` or `false` will be returned). If it does not exist in the ruleset, a message will be logged to let you know and the default value provided will be returned.

```Go
// default value of true is passed in

// example: the flag titled 'a_flag' does exist, and is toggled off
client.GetWithDefault("a_flag", true)
// returns false -> the flag exists and is toggled off

// example: the flag 'a_flag' does not exist in the ruleset
client.GetWithDefault("a_flag", true)
// returns true -> the flag doesn't exist in the ruleset; provided default value is returned.

// because the flag doesn't exist, the following message
// will be logged:
// The flag 'a_flag' is not in the ruleset. Returning the default value you provided,  true
```

### `client.GetWithContext(flagKey, context string) -> bool`

`GetWithContext()` allows you to take advantage of the rollout percentage set on a particular flag. The `GetWithContext()` function accepts the flag name, as well as a string context for the user. We recommend passing in something unique to the user such as a unique user identifier used within your application, or even an IP address.

The SDK will determine whether the user's context falls within the rollout percentage. This is done by summing the values of the code points within the string argument, and modding by 100 (the maximum rollout percentage possible). If the resulting value is less than or equal to the flag's rollout percentage, `GetWithContext()` will return `true`, and the user will funneled to the feature. If the user's context falls above the set rollout percentage, `GetWithContext()` returns false.

**NOTE:** If a flag is toggled off, `GetWithContext()` will return `false` no matter what the context argument is. This function is only relevant for flags that are toggled on.

Example usage below. In the below example, the flag is toggled on, and the flag's rollout percentage is set to 54%:

```Go
  dummy_uuid := "it-is-a-dummy-uuid"
  client.GetWithContext("test_flag", dummy_uuid) // true
  // the dummy_uuid sums to 13; the return value is true

  dummy_uuid2 := "ITSZ A DUMMY"
  client.GetWithContext("test_flag", dummy_uuid2) // false
  // the dummy_uuid2 sums to 55; the return value is false
```

Note that just like the `Get()` function, `GetWithContext()` will log an error if the flag does not exist within the ruleset. If you need flexibility there, use `GetWithContextWithDefault()`.

### `client.GetWithContextWithDefault(flagKey, context string, defaultVal bool) -> bool`

This function works just like `getWithContext()`, however it will not raise an error if the flag does not exist in the ruleset.

If the flag does not exist in the ruleset, the function will log a message letting the user know, and return the default value provided.

Example usage:

```Go
dummy_uuid := "it-is-a-dummy-uuid"
client.GetWithContextWithDefault("non-exist", dummy_uuid, false))
// returns false
// logs The flag ' non-exist ' is not in the ruleset.
// Returning the default value you provided,  false

client.GetWithContextWithDefault("test_flags", dummy_uuid, false))
// returns true
// the flag exists in the ruleset; the default value is ignored.
```

### Testing

To run unit tests, `go test` from root directory.

### Integrating with Google Analytics

You can set up an integration with Google Analytics by providing your GA tracking ID to `InitAnalytics()`. This tracking ID should begin with `UA-`.

You can then log an analytics event any time you wish using `LogAnalyticsEvent()`.

Example:

```Go
analytics := sdk.InitAnalytics("UA-XXXXXX-X")
// some application code ...
// when an event happens that you want to log to GA:
analytics.LogAnalyticsEvent("pioneer", "log", "1")
```
