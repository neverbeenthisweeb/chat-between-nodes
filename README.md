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
```
