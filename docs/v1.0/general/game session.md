---
title: Game Session
excerpt: How the Game Session works.
category: 65b8f16eb5f2930078659f93
order: 2
hidden: false
slug: general-game-session
---

When you request Public API to create a new game for you it actually creates a new Game Session for you.
Game Session is clearly identified by Game Session ID returned in the response.  

Every Game Session linked to a specific user, game and currency and has a TTL (time to live) which is 4 hours by default. 
Once time if over you need to create a new Game Session to continue playing.  
Client has ability to refresh this TTL time and theoretically extend it forever. But if you don't play game for 4 hours, this Game Session will be closed automatically.  

Game Session is a container for all user's game data. It contains balance, bets, wins, bonuses, etc. Also it uses in all client-to-backend calls from our games.

If user plays multiple games at the same time, he has multiple Game Sessions, also in case of different currencies you have different Game Sessions.  

Every Game Session should be linked to a single External Game Session ID (your game session ID) to be able to clearly identify user's game data (user, game, currency etc).  



