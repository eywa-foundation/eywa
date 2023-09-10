# eywa

# API DOCS

## User

### Register User

```bash
  eywad tx eywa register-user [PubKey] --from alice
```

- PubKey : User 공개키

### Get User By Address

```bash
  eywad q eywa get-user [Address]
```

- Address : USer 지갑주소

## Handshake

### Create Handshake

```bash
  eywad tx eywa create-handshake [Receiver] [RoomId] [ServerAddress] --from alice
```

- Receiver : Receiver 지갑주소
- RoomId : Chat Room Id
- ServerAddress : Server 주소 (혹은 도메인 혹은 이름 뭐가 되도 liter sting이라 그냥 적어주면 됨)

### Get Handshake By Receiver

```bash
  eywad q eywa get-handshake [Receiver]
```

- Receiver : Receiver 지갑주소

## Relay server

### Create Relay Server

```bash
  eywad tx eywa create-relay-server [Nickname] [Location] --from alice
```

- Nickname : Relay Server 닉네임
- Location : Relay Server 주소 (ip , domain)

### Get Relay Server All

```bash
  eywad q eywa get-relay-server-all
```

## Chat Room

### Create Chat Room

```bash
  eywad tx eywa create-chat-room [RoomId] [Receiver] [Message] [Time] --from alice
```

- RoomId : Chat Room Id
- Receiver : Receiver 지갑주소
- Message : 메세지
- Time : 메세지 보낸 시간

### Get Chat Room By RoomId

```bash
  eywad q eywa get-chat-room [RoomId]
```

- RoomId : Chat Room Id
