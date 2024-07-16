---
title: Getting started
excerpt: First steps in integration process.
category: 65b8f16eb5f2930078659f93
order: 0
hidden: false
slug: general-getting-started
---

# Terms Used

_Playhub/Playhub **_Playhub/Playhub Game Provider/Game Provider or Provider_ - Provider_**: Our platform, which platform that provides games and a Public API for integration. We also provide the Playhub Integration Adapter API for you to integrate your system with our platform.  platform.

_Aggregator/Operator **_Aggregator/Operator or Client_ - Client_**: Your system, which integrates with our platform.

# Provided APIs, Tools, and Services

Playhub provides two APIs **two APIs** for integration:

- `Playhub **Playhub Core API` should API**: 
  - **Integration**: Should be integrated with your backend system and used system.
  - **Usage**: Used to launch games and set up free rounds campaigns. You can also use this API to communicate campaigns.
  - **Additional Functions**: Communicate with our platform to get game history, free rounds information, user bonus balance, and other information.

- `Playhub **Playhub Integration Adapter API` should API**: 
  - **Integration**: Should be integrated with your backend system and used system.
  - **Usage**: Used to receive game events from our platform.

**Note**: Upon request, we can provide you with OpenAPI specifications for our Public APIs.

# Introduction

The integration process consists of three phases:  
**three phases**:

1. Registering **Registering** and providing us with your integration details.
2. Integrating **Integrating** your backend with our Public APIs.
3. Setting up **Setting up** your frontend/client to launch our games.

## Registration Process

To start using our Public API, you need to register follow these steps:

1. **Register an account account** on our portal (or request via any communication channel) and obtain channel).

2. **Obtain the following details:

details**:
   - Client ID: This is your **Client ID**: Your unique identifier in our system, which identifies your system as our client.
system.
   - API **API Token Key: This token key is used Key**: Used for all requests to our Public API. See the [Authentication](general-auth) section for more details. Please handle **Important**: Handle this token key securely and do not expose it publicly to avoid compromising the public; otherwise, the whole integration could be compromised and blocked.
integration.
   - Callback **Callback Token Key: This token key is used Key**: Used in all callback requests from our platform to your system.

From your side, we require 3. **Provide us with the following details:

details**:
   - URL **URL for `Get Balance` endpoint: This endpoint should return endpoint**: Returns the current user's balance in your system. We will use this endpoint to get the user's balance before launching a game.
system.
   - URL **URL for `Bet` endpoint: This endpoint should be called when a user places a bet in our game. We will call this endpoint endpoint**: Called when a user places a bet in our game.
   - URL **URL for `Win` endpoint: This endpoint should be called when a user wins in our game. We will call this endpoint endpoint**: Called when a user wins in our game.
   - URL **URL for `Refund` endpoint: This endpoint should be called when a user is refunded in our game. We will call this endpoint endpoint**: Called when a user is refunded in our game.

**Notes**:

- All URLs should support HTTPS and be accessible from our servers. servers.
- We will use the `POST` method to call your endpoints.  
endpoints.
- For more details on callbacks, see [Playhub Integration Adapter API](playhub-integration-adapter-intro).

Additionally, if you require, - If required, we can set up whitelisting for your platform. Please attach provide a list of IP addresses that should to be whitelisted.

From our side, we use the following ### IPs We Use to call your endpoints:  
Staging environment: 34.76.77.205  Call Your Endpoints:

Production environment: - **Staging environment**: 34.76.77.205
- **Production environment**: 34.140.75.241

## Development Process

- **Testing and Development**: You are free to use our staging environment for any kind of tests and development without any rate limiting.

## Setting Up Your Frontend/Client

**Ensure** your frontend/client is set up Your Frontend/Clientto launch our games.

***

**Important Points to Remember**:

- Secure your **API Token Key**.
- Ensure your endpoints are accessible via **HTTPS**.
- Use the **staging environment** for testing and development.

For any additional clarifications or assistance, please contact our support team.# Terms Used

**_Playhub/Playhub Game Provider/Game Provider or Provider_**: Our platform that provides games and a Public API for integration. We also provide the Playhub Integration Adapter API for you to integrate your system with our platform.

**_Aggregator/Operator or Client_**: Your system, which integrates with our platform.

# Provided APIs, Tools, and Services

Playhub provides **two APIs** for integration:

- **Playhub Core API**: 
  - **Integration**: Should be integrated with your backend system.
  - **Usage**: Used to launch games and set up free rounds campaigns.
  - **Additional Functions**: Communicate with our platform to get game history, free rounds information, user bonus balance, and other information.

- **Playhub Integration Adapter API**: 
  - **Integration**: Should be integrated with your backend system.
  - **Usage**: Used to receive game events from our platform.

**Note**: Upon request, we can provide you with OpenAPI specifications for our Public APIs.

# Introduction

The integration process consists of **three phases**:

1. **Registering** and providing us with your integration details.
2. **Integrating** your backend with our Public APIs.
3. **Setting up** your frontend/client to launch our games.

## Registration Process

To start using our Public API, follow these steps:

1. **Register an account** on our portal (or request via any communication channel).

2. **Obtain the following details**:
   - **Client ID**: Your unique identifier in our system.
   - **API Token Key**: Used for all requests to our Public API. See the [Authentication](general-auth) section for more details. **Important**: Handle this token key securely and do not expose it publicly to avoid compromising the integration.
   - **Callback Token Key**: Used in all callback requests from our platform to your system.

3. **Provide us with the following details**:
   - **URL for `Get Balance` endpoint**: Returns the current user's balance in your system.
   - **URL for `Bet` endpoint**: Called when a user places a bet in our game.
   - **URL for `Win` endpoint**: Called when a user wins in our game.
   - **URL for `Refund` endpoint**: Called when a user is refunded in our game.

**Notes**:

- All URLs should support HTTPS and be accessible from our servers.
- We will use the `POST` method to call your endpoints.
- For more details on callbacks, see [Playhub Integration Adapter API](playhub-integration-adapter-intro).
- If required, we can set up whitelisting for your platform. Please provide a list of IP addresses to be whitelisted.

### IPs We Use to Call Your Endpoints:

- **Staging environment**: 34.76.77.205
- **Production environment**: 34.140.75.241

## Development Process

- **Testing and Development**: You are free to use our staging environment for any kind of tests and development without any rate limiting.

## Setting Up Your Frontend/Client

**Ensure** your frontend/client is set up to launch our games.

***

**Important Points to Remember**:

- Secure your **API Token Key**.
- Ensure your endpoints are accessible via **HTTPS**.
- Use the **staging environment** for testing and development.

For any additional clarifications or assistance, please contact our support team.