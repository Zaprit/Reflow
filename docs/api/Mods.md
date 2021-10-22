# Mod Api
This api "governs" access to mods.  
In reality, it's a fancy directory listing with some md5 hashes.

## Mod List

```http request
GET /mod
```

This returns a list of all mods in the system in a (terrible) use of json objects.

### Mod List JSON Example

```json
{"mods":{
  "mod1": "First Mod",
  "mod2": "Second Mod"
}}
```

## Mod

```http request
GET /mod/<slug>
```

This gets a mod's data and a list of builds

### Mod JSON Example

```json
{
  "name":"test-mod",
  "pretty_name":"Test Mod",
  "author":"Test Author",
  "description":"Test Description",
  "link":"https://example.org",
  "versions":["1.0"]
}
```