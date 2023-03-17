# Chat Between Nodes

A service endpoint for a chat client that shows chats between different user nodes.

Dependencies

- Go (`1.18.7`)

How to run

- Run `go mod tidy`
- Run `go run .`
- Hit the `GET /node/:id/chats` HTTP endpoint. Where `:id` is your node ID

```
// Example

curl -X GET \
  'localhost:3000/node/1002/chats' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'

[
  {
    "id": "3",
    "node_id": "1002",
    "from": "Sandy",
    "message": "Hi, anybody here?",
    "timestamp": "2023-03-17T21:02:00+07:00"
  },
  {
    "id": "4",
    "node_id": "1002",
    "from": "Larry",
    "message": "Hi",
    "timestamp": "2023-03-17T21:03:00+07:00"
  },
  {
    "id": "5",
    "node_id": "1002",
    "from": "Larry",
    "message": "HII!!! I AM HERE!",
    "timestamp": "2023-03-17T21:04:00+07:00"
  }
]
```
