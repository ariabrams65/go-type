package messages

import (
    "encoding/json"
    "errors"
    "net"
)

type Message interface {
    Type() string
}

type JoinMessage struct {
    Username string 
    Roomname string
}

func (m JoinMessage) Type() string {
    return "join"
}

type PositionMessage struct {
    Username string
    Index int
}

func (m PositionMessage) Type() string {
    return "position"
}

type TextMessage struct {
    Text string
}

func (m TextMessage) Type() string {
    return "text"
}

type ErrorMessage struct {
    Error string
}

func (m ErrorMessage) Type() string {
    return "error"
}

type frame struct {
    Type string
    Data json.RawMessage
}

func EncodeMessage(m Message, conn net.Conn) error {
    data, err := json.Marshal(m) 
    if err != nil {
        return err
    }

    json.NewEncoder(conn).Encode(frame{
        Type: m.Type(),
        Data: data,
    })
    return nil
}

func DecodeMessage(conn net.Conn) (Message, error) {
    var f frame
    err := json.NewDecoder(conn).Decode(&f)
    if err != nil {
        return nil, err
    }

    switch f.Type {
    case "join":
        var m JoinMessage 
        err := json.Unmarshal(f.Data, &m)
        return m, err

    case "position":
        var m PositionMessage
        err := json.Unmarshal(f.Data, &m)
        return m, err

    case "text":
        var m TextMessage
        err := json.Unmarshal(f.Data, &m)
        return m, err

    case "error":
        var m ErrorMessage
        err := json.Unmarshal(f.Data, &m)
        return m, err
    }
    return nil, errors.New("Unknown type") 
}

