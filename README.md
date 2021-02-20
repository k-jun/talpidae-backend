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
FakeArrowUp    = 8
FakeArrowDown  = 9
FakeArrowLeft  = 10
FakeArrowRight = 11
TrapArrow      = 12
TrapTreasure   = 13
```

## apis

### GET `/game/start?game_id=3f984f86-49b0-371e-9449-6047e9241b68`

start/reset specified game meaning make the game very first state

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

### GET `/game/field?game_id=54d44bb1-4777-3aec-af20-62f94ca9ee04`

get current field of specified game

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
  ],
  "current_number_of_users": 2
}
```

### POST `/game/fill?game_id=b9f73c4d-7bda-32b8-9877-a2e80bf021ba`

fill a cell in specified game field with value (WanaArrowUp | WanaArrowDown | WanaArrowLeft | WanaArrowRight)

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

### GET `/game/logs?game_id=abfde535-6e86-3259-8875-2b19be76664f`

get change logs of specified game field by a fill endpoint

#### request body

```json
{}
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

### POST `/match/join`

join to a random match, after joining you need to keep calling field endpoint until the number of users reaching to max.

#### request body

```json
{
  "id": "72f10c8b-b881-30ee-8aa3-97b9b132a017",
  "name": "Reba Walker"
}
```

#### response body

```json
{
  "game_id": "720042dc-f72a-3dbd-a265-2e367654dbf2",
  "user_id": "72f10c8b-b881-30ee-8aa3-97b9b132a017"
}
```

### POST `/match/join`

leave from a match

#### request body

```json
{
  "game_id": "c5f3a644-3cac-3841-af03-bb2409f55674",
  "user_id": "6105f442-fcf2-3362-9ba9-1d9f72385a76"
}
```

#### response body

```json
{}
```
