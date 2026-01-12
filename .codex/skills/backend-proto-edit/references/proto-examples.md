# Proto 编辑示例

基于 `api/admin/v1/user.proto` 风格。

## 1. Service 定义

```protobuf
//用户表
service User {
  //用户表-创建一条数据
  rpc CreateUser(CreateUserReq) returns (CreateUserReply) {
    option (google.api.http) = {
      post: "/admin/v1/user/create"
      body: "*"
    };
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
      parameters: {
        headers: {
          name: "Authorization"
          description: "Bearer Token"
          type: STRING
          required: true
        }
      }
    };
  }
}
```

## 2. Info 消息

```protobuf
//用户表信息
message UserInfo {
  string id = 1; // id
  string phone = 2; // 手机
  string nickname = 3; // 昵称
  int32 gender = 4; // 性别（0未知 1男 2女）
  int32 status = 9; // 状态
  string createdAt = 10; // 创建时间
  string updatedAt = 11; // 更新时间
  UserMembershipInfo userMembershipInfo = 12; // 关联信息
}
```

## 3. Create 请求

```protobuf
//请求-用户表-创建一条数据
message CreateUserReq {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      required: [
        "phone",
        "status"
      ]
    }
  };
  string phone = 1 [(buf.validate.field).string = {min_len: 1}]; // 手机
  string nickname = 2 [
    (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED,
    (buf.validate.field).string = {min_len: 1}
  ]; // 昵称
  int32 gender = 3 [(buf.validate.field).ignore = IGNORE_IF_UNPOPULATED]; // 性别
  int32 status = 6; // 状态
}

//响应-用户表-创建一条数据
message CreateUserReply {
  string id = 1; // id
}
```

## 4. Update 请求

```protobuf
//请求-用户表-更新一条数据
message UpdateUserReq {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      required: ["id"]
    }
  };
  string id = 1 [(buf.validate.field).string = {
    min_len: 1
    max_len: 128
  }]; // id
  string nickname = 2 [
    (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED,
    (buf.validate.field).string = {min_len: 1}
  ]; // 昵称
  int32 status = 6; // 状态
}

//响应-用户表-更新一条数据
message UpdateUserReply {}
```

## 5. List 请求（带过滤）

```protobuf
//请求-用户表-列表数据查询
message GetUserListReq {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      required: ["page", "pageSize"]
    }
  };
  int32 page = 1 [(buf.validate.field).int32 = {gte: 1}]; //页码
  int32 pageSize = 2 [(buf.validate.field).int32 = {gte: 1, lte: 1000}]; //页数
  string nickname = 3; // 昵称
  string phone = 4; // 手机
  int32 status = 5; // 状态
  repeated string createdAt = 6; // 创建时间
}

//响应-用户表-列表数据查询
message GetUserListReply {
  int32 total = 1; //总数
  repeated UserInfo list = 2; // 列表数据
}
```
