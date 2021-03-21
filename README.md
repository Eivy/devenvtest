# DevEnvTest

開発環境のテスト用

## 開発

`query.sql`を編集して`go run github.com/kyleconroy/sqlc/cmd/sqlc generate`を実行してDBとのインターフェースを作成.

`graph/schema.graphqls`を編集して`go run github.com/99designs/gqlgen generate`を実行してGraphQL用のインターフェースを作成.
実際にデータの取得・登録等は`graph/schema.resolver.go`に実装する.

### テスト実行

`docker-compose up -d`でDBとAPIサーバーが起動する.
http://localhost:8080 にアクセスしてGraphQLのクエリの実行が行える.

http://localhost:8080/migrate?step=2 にアクセスすることでDBのマイグレーションを実行する. `step=2`の数字部分はDBのバージョンによる. 負数を指定するとデグレードできる.

## DBスキーマ
[users](./migrations/000001_create_table_user.up.sql)と[items](./migrations/000002_create_table_items.up.sql)があります.

## GraphQLクエリ

### ユーザー一覧
```graphql
query {
  users {
    id
    name
  }
}
```

### ユーザー取得
```graphql
query {
  user(id: 1) {
    id
    name
  }
}
```

### ユーザー追加
```graphql
mutation {
  createUser(input: {name: "ユーザー名"}){
      name
  }
}
```

### ユーザー削除
```graphql
mutation {
  deleteUser(id: 1) {
    id
    name
  }
}
```

### アイテム一覧
```graphql
query {
  items {
    id
    name
    location
    manager {
      id
      name
    }
  }
}
```

### アイテム取得
```graphql
query {
  item(id: 0){
    id
    name
    counts
    location
    manager {
      id
      name
    }
  }
}
```

### アイテム追加
```graphql
mutation {
  createItem(input: {name: "thing1", location: "anywhere",counts: 3, manager: 3}){
    id
    name
    location
    manager {
      id
      name
    }
  }
}
```

### アイテム削除
```graphql
mutation {
  deleteItem(id: 3) {
    id
    name
    location
    counts
    manager {
      name
      id
    }
  }
}
```
