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
{}
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
      "h": 38,
      "w": 38,
      "value": "yazirusi"
    },
    {
      "h": 50,
      "w": 49,
      "value": "otakara"
    },
    ...
  ]
}
```

### POST `/fill`

fill a cell in the game field with value (wanawan | otakara | yazirusi)

```json
{
    "h": 120,
    "w": 49,
    "value": "wanawana"
}
```

#### `response body`

```json
{}
```