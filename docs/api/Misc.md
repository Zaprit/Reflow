# Misc API functions

These are the miscilaneous api endpoints that don't really fit anywhere else

## Root

```http request
GET /
```

This is the api endpoint that returns the server name, branch and version.  
The response format is JSON and looks like this

```json
{"api":"Reflow","version":"v0.1","stream":"master"}
```
Here's the api string from an official solder instance
```json
{"api":"TechnicSolder","version":"v0.7.7","stream":"DEV"}
```

As far as I can tell (and as of 2021), the client never checks the contents of this string just that it is valid json that fits the structure