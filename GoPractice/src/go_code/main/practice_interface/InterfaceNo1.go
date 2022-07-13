package main

import "bufio"

/*
ReadWriter结构定义
 */
type ReadWriter struct {
    // Read指针
    *bufio.Reader

    // Writer指针
    *bufio.Writer
}

func NewReadWriter(r *bufio.Reader, w *bufio.Writer) *ReadWriter {
    return nil
}