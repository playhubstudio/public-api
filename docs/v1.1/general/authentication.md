---
title: Authentication
excerpt: Authentication for our Public APIs requests.
category: 671b459dfe48cf0030150ea1
order: 1
hidden: false
slug: general-auth
---

When you register an account on our portal, you will be given an API token key. 
This key is used to authenticate your requests to our APIs.

## Signing your request  

We assume your request is authenticated if you provide valid signature for the whole request in the `X-REQUEST-SIGN` header.

**Request example with curl:**
```curl
curl -X GET '<server url>' \    
    -H "X-REQUEST-SIGN: <singature>"
```

where `<server url>` is a URL to our public API (staging or production environment),   `<singature>` is a signature for the whole request.

The whole request payload should be signed, in case of `GET` requests, the payload is virtual - we use query parameters to construct the JSON  
and use this JSON for signing. Please take a look at the following examples.

Before signing your JSON request should be in a compact form and sorted by keys. It's called canonical JSON.

**Sign example in Go:**
```go
func Sign(payload, token string) string {
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(payload))
return hex.EncodeToString(hash.Sum(nil))
}
```

**Sign example for GET request in Go:**
```go
func Sign(queryParams url.Values, token string) string {
  reqJSON, err := URLValuesFlatteredToJSON(params)
  _ = err
	hash := hmac.New(sha256.New, []byte(token))
	hash.Write([]byte(reqJSON))
return hex.EncodeToString(hash.Sum(nil))
}

func URLValuesFlatteredToJSON(values url.Values) ([]byte, error) {
	flattened := make(map[string]string)
	for key, vals := range values {
		if len(vals) > 0 {
			flattened[key] = vals[0]
		}
	}

	jsonData, err := json.Marshal(flattened)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
```

**Sign example with JavaScript:**
```js
const crypto = require('crypto');

function sign(payload, token) {
  const hmac = crypto.createHmac('sha256', token);
  hmac.update(payload);
  const hash = hmac.digest('hex');
  return hash;
}

const payload = 'example payload';
const token = 'example token';
const signature = sign(payload, token);
console.log(signature);