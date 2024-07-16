---
title: Game Session
excerpt: How the Game Session works.
category: 65b8f16eb5f2930078659f93
order: 2
hidden: false
slug: general-game-session
---

# Game Session Management

When you request the Public API to create a new game, it actually creates a new Game Session for you. Each Game Session is identified by a Game Session ID, which is returned in the response.

## Key Details of Game Sessions

- **Identification**: Every Game Session is linked to a specific user, game, and currency.
- **TTL (Time to Live)**: The default TTL for a Game Session is 4 hours. Once the time is over, you need to create a new Game Session to continue playing. Clients can refresh this TTL to extend it indefinitely. However, if you do not play the game for 4 hours, the Game Session will be closed automatically.
- **Data Container**: A Game Session contains all user game data, including balance, bets, wins, bonuses, and more. It is used in all client-to-backend calls from our games.

## Multiple Game Sessions

- **Concurrent Games**: If a user plays multiple games simultaneously, they will have multiple Game Sessions.
- **Different Currencies**: Different currencies will result in separate Game Sessions.

## Linking Game Sessions

Every Game Session should be linked to a single External Game Session ID (your game session ID) to clearly identify user game data, including user, game, and currency.

***

**Important Points to Remember**:

- **Game Session ID**: Identifies each Game Session.
- **Default TTL**: 4 hours, refreshable by the client.
- **Data Containment**: Includes all relevant user game data.
- **Multiple Sessions**: Managed for concurrent games and different currencies.
- **Linking**: Use External Game Session ID for clear identification of user game data.When you request the Public API to create a new game, it actually creates a new Game Session for you. Each Game Session is identified by a Game Session ID, which is returned in the response.

## Key Details of Game Sessions

- **Identification**: Every Game Session is linked to a specific user, game, and currency.
- **TTL (Time to Live)**: The default TTL for a Game Session is 4 hours. Once the time is over, you need to create a new Game Session to continue playing. Clients can refresh this TTL to extend it indefinitely. However, if you do not play the game for 4 hours, the Game Session will be closed automatically.
- **Data Container**: A Game Session contains all user game data, including balance, bets, wins, bonuses, and more. It is used in all client-to-backend calls from our games.

## Multiple Game Sessions

- **Concurrent Games**: If a user plays multiple games simultaneously, they will have multiple Game Sessions.
- **Different Currencies**: Different currencies will result in separate Game Sessions.

## Linking Game Sessions

Every Game Session should be linked to a single External Game Session ID (your game session ID) to clearly identify user game data, including user, game, and currency.

***

**Important Points to Remember**:

- **Game Session ID**: Identifies each Game Session.
- **Default TTL**: 4 hours, refreshable by the client.
- **Data Containment**: Includes all relevant user game data.
- **Multiple Sessions**: Managed for concurrent games and different currencies.
- **Linking**: Use External Game Session ID for clear identification of user game data.