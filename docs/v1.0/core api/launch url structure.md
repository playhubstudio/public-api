---
title: Launch URL
excerpt: Launch URL structure and options.
category: 65bc99b3a8f20d00395a35e7
order: -1
hidden: false
slug: core-launch-url-structure
---

# Game Launcher URL Logic and Parameters

## URL Template

The URL template for launching a game should follow this format:

```
https://<game_provider_base_url>?<query_parameters>
```

- **\<game_provider_base_url>**: This is the fundamental URL address required for launching the game.
- **\<query_parameters>**: These are the request parameters formatted as a query string, essential for initiating the specific product.

## Base URLs

- **Staging Environment**: `https://staging.fastplaynetwork.com/games/<gameid>`
- **Production Environment**: `https://prod.fastplaynetwork.com/games/<gameid>`

Example: https://staging.fastplaynetwork.com/games/ph_limbo_97/index.html?

## Requirements

The game launcher requires certain parameters to be passed as GET parameters. These parameters must be URL encoded. Below are the required parameters along with examples.

### Required Parameters

These parameters are essential for the game launcher to function correctly:

- **game**: The identifier of the game you want to launch.
- **locale**: The language and region code based on ISO 639-1 standard.
- **currency**: The currency code based on ISO 4217 standard.

#### Example

```
https://staging.fastplaynetwork.com/games/ph_limbo_97/index.html?
&locale=en-gb
&currency=EUR
```

### Notes

- All parameters should be URL encoded to ensure they are correctly interpreted by the game launcher.
- Replace **\<game_provider_base_url>** with either the staging or production URL based on your environment.
- Construct the **query_parameters** as a valid query string, with each parameter separated by an ampersand (&).


