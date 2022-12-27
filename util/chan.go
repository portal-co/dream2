package util

import "github.com/docker/libchan"

type Sender[T any] struct {
	V libchan.Sender
}

type Receiver[T any] struct {
	V libchan.Receiver
}

type Pipe[T any] struct {
	Sender[T]
	Receiver[T]
}

func SendX[T any](v T, s Sender[T]) error {
	return s.V.Send(v)
}

func CloseX[T any](s Sender[T]) error {
	return s.V.Close()
}

func RecvX[T any](r Receiver[T]) (T, error) {
	var v T
	err := r.V.Receive(&v)
	if err != nil {
		return v, err
	}
	return v, nil
}
