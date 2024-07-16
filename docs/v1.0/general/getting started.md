---
title: Getting started
excerpt: First steps in integration process.
category: 65b8f16eb5f2930078659f93
order: 0
hidden: false
slug: general-getting-started
---

# Terms used

*Playhub/Playhub Game Provider/Game Provider or Provider* - our platform, which provides games and Public API for integration. Also we provide Playhub Integration Adapter API for you to integrate your system with our platform.  
*Aggregator/Operator or Client* - your system, which integrates with our platform.

# Provided APIs, tools and services

Playhub provides 2 APIs for integration:
- `Playhub Core API` which should be integrated with your back end system and used to launch games and setup free rounds campaigns. 
Also you can use this API to communicate with our platform to get game history, free rounds information, user bonus balance and other information.
- `Playhub Integration Adapter API` which should be integrated with your back end system and used to receive game events from our platform.

On a request we can provide you OpenAPI specifications for our Public APIs (including Playhub Integration Adapter API).

# Introduction

Integration process consists of 3 phases:  
1. Registering and providing us with your integration details.
2. Integrate your back end with our Public APIs.
3. Setup your front end/client to launch our games.


## Registration Process

To start using our Public API you need to register an account on our portal (or request via any communication channel) and get the following details:
- Client ID. This is your unique identifier in our system which identifies your system as our client.
- API token key. This token key is used to all requests to our Public API. See [Authentication](general-auth) section for more details.  Please handle this token key securely and do not expose it to public otherwise the whole integration could be compromised and blocked.
- Callback token key. This token key is used in all callback requests from our platform to your system.

From your side we require the following details:
- URL for `Get Balance` endpoint. This endpoint should return current user's balance in your system. We will use this endpoint to get user's balance before launching a game.
- URL for `Bet` endpoint. This endpoint should be called when user places a bet in our game. We will call this endpoint when user places a bet in our game.
- URL for `Win` endpoint. This endpoint should be called when user wins in our game. We will call this endpoint when user wins in our game.
- URL for `Refund` endpoint. This endpoint should be called when user is refunded in our game. We will call this endpoint when user is refunded in our game.

All URLs should support HTTPS and be accessible from our servers. We will use `POST` method to call your endpoints. 
More details on callbacks see at [Playhub Integration Adapter API](playhub-integration-adapter-intro).

In additional if you require we could setup whitelisting for your platform - please attach list of IP addresses which should be whitelisted.

From our side we use the following IPs to call your endpoints:
staging environment: 34.76.77.205
production environment: 34.140.75.241

## Development Process

You are free to use our staging environment to do any kind of tests and development without any rate limiting. 

## Setting up your frontent/client