## chat example


```go

    // client1 
    go run examples/chat/client/client.go

    // client2 conn to client1
    go run examples/chat/client/client.go -p 1234 -t localhost:6666

```

## Instruction
1v1 command-line chat example,Only the sender and the receiver of the messages have the “key”(created by pass and salt) to read them,This ensures that no one besides you and the person you're talking to can decipher the messages.

