package observer

import (
	"context"
)

type Promise[T any] struct {
	ctx        context.Context
	done       chan struct{} // Сигнализирует о завершении
	data       T
	err        error
	degradable bool
}

func Async[T any](ctx context.Context, fn func() (T, error)) *Promise[T] {
	p := &Promise[T]{
		ctx:        ctx,
		done:       make(chan struct{}),
	}

	// Запуск асинхронного выполнения задачи
	go p.execute(fn)
	return p
}

// execute выполняет функцию и сохраняет результат
func (p *Promise[T]) execute(fn func() (T, error)) {
	type channel struct {
		data T
		err  error
	}

	// Канал для получения результата асинхронной операции
	ch := make(chan channel, 1)

	// Запуск задачи в отдельной горутине
	go func() {
		data, err := fn()
		ch <- channel{data, err}
	}()

	select {
	case <-p.ctx.Done():
		p.err = p.ctx.Err() // Сохраняем ошибку контекста
	case res := <-ch:
		p.data = res.data
		p.err = res.err
	}

	close(p.done) // Сигнализируем о завершении
}

func Await[T any](p *Promise[T]) (T, error) {
	if p == nil {
		var empty T
		return empty, nil
	}

	<-p.done

	return p.data, p.err
}

func (p *Promise[T]) IsDegradable() bool {
	return p.degradable
}

func (p *Promise[T]) Degradable() *Promise[T] {
	p.degradable = true
	return p
}
