# talpidae-backend 


## block-type enums

```
SakuSaku       = 0
KachiKachi     = 1
GochiGochi     = 2
Treasure       = 3
ArrowUp        = 4
ArrowDown      = 5
ArrowLeft      = 6
ArrowRight     = 7
WanaArrowUp    = 8
WanaArrowDown  = 9
WanaArrowLeft  = 10
WanaArrowRight = 11
```

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
  "fields": [
    [0,0,0,0],
    [1,1,1,0],
    [1,4,1,0],
    [1,1,1,0],
    [0,0,0,0]
  ]
}
```

### GET `/field`

get current field of the game

#### request body

```json
{}
```

#### response body

```json
{
  "fields": [
    [0,0,0,0],
    [1,1,1,0],
    [1,4,1,0],
    [1,1,1,0],
    [0,0,0,0]
  ]
}
```

### POST `/fill`

fill a cell in the game field with value (arrow-left | arrow-right | arrow-up | arrow-down | treasure)

#### request body

```json
{
  "user_id": "76d96dd2-e0a5-3e7f-a947-2469b8804d06",
  "value": 8,
  "h": 120,
  "w": 49
}
```

#### response body

```json
{}
```

### GET `/logs`

get change logs of game field by a fill endpoint

#### request body

```json
```

#### response body

```json
{
  "logs": [
    {"user_id": "2a3df377-4ce5-3b8f-bb7e-9e0a0f66176f" , "value": 9, "w": 12, "h": 21},
    {"user_id": "bb673d5e-d206-3bfb-a0e2-0938fc1fe4cf" , "value": 11, "w": 2, "h": 18}
  ]
}
```
