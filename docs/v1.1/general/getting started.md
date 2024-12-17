---
title: Getting started
excerpt: First steps in integration process.
category: 671b459dfe48cf0030150ea1
order: 0
hidden: false
slug: general-getting-started
---

# Terms Used

_Playhub/Playhub Game Provider/Game Provider or Provider_ - Our platform, which provides games and a Public API for integration. We also provide the Playhub Integration Adapter API for you to integrate your system with our platform.  
_Aggregator/Operator or Client_ - Your system, which integrates with our platform.

# Provided APIs, Tools, and Services

Playhub provides two APIs for integration:

- `Playhub Core API` should be integrated with your backend system and used to launch games and set up free rounds campaigns. You can also use this API to communicate with our platform to get game history, free rounds information, user bonus balance, and other information.
- `Playhub Integration Adapter API` should be integrated with your backend system and used to receive game events from our platform.

Upon request, we can provide you with OpenAPI specifications for our Public APIs.

# Introduction

The integration process consists of three phases:  

1. Registering and providing us with your integration details.
2. Integrating your backend with our Public APIs.
3. Setting up your frontend/client to launch our games.

## Registration Process

To start using our Public API, you need to register an account on our portal (or request via any communication channel) and obtain the following details:

- Client ID: This is your unique identifier in our system, which identifies your system as our client.
- API Token Key: This token key is used for all requests to our Public API. See the [Authentication](general-auth) section for more details. Please handle this token key securely and do not expose it to the public; otherwise, the whole integration could be compromised and blocked.
- Callback Token Key: This token key is used in all callback requests from our platform to your system.
- Legal bets table (if free rounds is requested). See the [Legal Bets](core-free-rounds-legal-bet).  

From your side, we require the following details:

- URL for `Get Balance` endpoint: This endpoint should return the current user's balance in your system. We will use this endpoint to get the user's balance before launching a game.
- URL for `Bet` endpoint: This endpoint should be called when a user places a bet in our game. We will call this endpoint when a user places a bet in our game.
- URL for `Win` endpoint: This endpoint should be called when a user wins in our game. We will call this endpoint when a user wins in our game.
- URL for `Refund` endpoint: This endpoint should be called when a user is refunded in our game. We will call this endpoint when a user is refunded in our game.  



All URLs should support HTTPS and be accessible from our servers. We will use the `POST` method to call your endpoints.  
For more details on callbacks, see [Playhub Integration Adapter API](playhub-integration-adapter-intro). Please note that Open API specification provides different endpoints for 
different requests, but depends on your needs we could use a single endpoint for all requests. Just make sure that in this case you need to skip Open API specification endpoints list and use your own you provided with us.

Additionally, if you require, we can set up whitelisting for your platform. Please share a list of IP addresses that should be whitelisted via the integration Skype chat or via email: integration@playhub.studio.

More about whitelisting can be found in the [IP Addresses](playhub-integration-adapter-intro) section.

## Development Process

You are free to use our staging environment for any kind of tests and development without any rate limiting.
