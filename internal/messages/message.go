package messages

import (
    "encoding/json"
    "errors"
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

type frame struct {
    Type string
    Data json.RawMessage
}

func EncodeMessage(m Message, encoder *json.Encoder) error {
    data, err := json.Marshal(m) 
    if err != nil {
        return err
    }

    encoder.Encode(frame{
        Type: m.Type(),
        Data: data,
    })
    return nil
}

func DecodeMessage(decoder *json.Decoder) (Message, error) {
    var f frame
    err := decoder.Decode(&f)
    if err != nil {
        return nil, err
    }

    switch f.Type {
    case "join":
        var m JoinMessage 
        err := json.Unmarshal(f.Data, &m)
        if err != nil {
            return nil, err
        }
        return m, nil

    case "position":
        var m PositionMessage
        err := json.Unmarshal(f.Data, &m)
        if err != nil {
            return nil, err
        }
        return m, nil

    case "text":
        var m TextMessage
        err := json.Unmarshal(f.Data, &m)
        if err != nil {
            return nil, err
        }
        return m, nil
    }
    return nil, errors.New("Unknown type") 
}

