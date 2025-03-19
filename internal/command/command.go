package command

import (
	"bufio"
	"os"
	"strings"
)

type Handler struct {
	reader *bufio.Reader
	Ch     chan Command
}

type Command struct {
	Command string
	Value   string
}

func NewHandler() *Handler {
	return &Handler{
		reader: bufio.NewReader(os.Stdin),
		Ch:     make(chan Command),
	}
}

func (h *Handler) Run() {
	for {
		text, _ := h.reader.ReadString('\n')
		s := strings.Split(text, " ")
		if len(s) < 2 {
			continue
		}

		h.Ch <- Command{
			Command: s[0],
			Value:   strings.TrimSpace(s[1]),
		}
	}
}
