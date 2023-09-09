# eywa

**eywa** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/eywa@latest! | sudo bash
```

`username/eywa` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

# Example

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
