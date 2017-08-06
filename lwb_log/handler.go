/**
 * log处理类库（类库）
 * ========================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ------------------------------------------------------------------------
 * ========================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_encode 17155 2017-02-06 06:29:05Z $
 */

package lwb_log

import (
	"io"
)

//Handler writes logs to somewhere
type Handler interface {
	Write(p []byte) (n int, err error)
	Close() error
}

//StreamHandler writes logs to a specified io Writer, maybe stdout, stderr, etc...
type StreamHandler struct {
	w io.Writer
}

func NewStreamHandler(w io.Writer) (*StreamHandler, error) {
	h := new(StreamHandler)

	h.w = w

	return h, nil
}

func (h *StreamHandler) Write(b []byte) (n int, err error) {
	return h.w.Write(b)
}

func (h *StreamHandler) Close() error {
	return nil
}

//NullHandler does nothing, it discards anything.
type NullHandler struct {
}

func NewNullHandler() (*NullHandler, error) {
	return new(NullHandler), nil
}

func (h *NullHandler) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (h *NullHandler) Close() error {
	return nil
}
