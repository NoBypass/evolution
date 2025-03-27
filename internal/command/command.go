package command

import (
	"bufio"
	"evolution/internal/environment"
	"fmt"
	"os"
	"strconv"
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

func (h *Handler) Run(env *environment.Environment) {
	for {
		text, _ := h.reader.ReadString('\n')
		s := strings.Split(text, " ")
		if len(s) < 2 {
			continue
		}

		command, value := s[0], s[1]
		switch command {
		case "mspt":
			mspt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Printf("Invalid mspt value %s\n", value)
			}

			env.MSPT = mspt
		}
	}
}
