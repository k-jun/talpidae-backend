# talpidae-backend 


## apis

### GET `/start`

start/reset the game meaning make the game very first state

#### request body

```json
{}
```

#### response body

```json
{
  "positions": [
    {
      "h": 8,
      "w": 64,
      "value": "arrow-left"
    },
    {
      "h": 33,
      "w": 6,
      "value": "treasure"
    },
    ...
}
```

### GET `/status`

get current state of the game

#### request body

```json
{}
```

#### response body

```json
{
  "positions": [
    {
      "h": 8,
      "w": 64,
      "value": "arrow-left"
    },
    {
      "h": 33,
      "w": 6,
      "value": "treasure"
    },
    ...
}
```

### POST `/fill`

fill a cell in the game field with value (arrow-left | arrow-right | arrow-up | arrow-down | treasure)

#### request body

```json
{
    "h": 120,
    "w": 49,
    "value": "arrow-right"
}
```

#### response body

```json
{}
```
