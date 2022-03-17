package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var TeamBoard = make(map[string]*Team)

type Team struct {
	name           string
	MP, W, D, L, P int
}

func updateTeam(teamName string, gameUpdate int) {
	_, ok := TeamBoard[teamName]
	if !ok {
		*TeamBoard[teamName] = Team{name: teamName, MP: 0, W: 0, D: 0, L: 0, P: 0}
	}
	(*TeamBoard[teamName]).MP += 1
	switch gameUpdate {
	case 1:
		(*TeamBoard[teamName]).W += 1
		(*TeamBoard[teamName]).P += 3
	case -1:
		(*TeamBoard[teamName]).L += 1
	case 0:
		(*TeamBoard[teamName]).D += 1
		(*TeamBoard[teamName]).P += 1
	}
}

func Tally(reader io.Reader, writer io.Writer) error {

	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	bufString := buf.String()
	if err != nil {
		return err
	}
	line := ""
	for a := range bufString {
		switch bufString[a] {
		case '\n':
			game := strings.Split(line, ";")

			switch game[2] {
			case "win":
				updateTeam(game[0], 1)
				updateTeam(game[1], -1)
			case "loss":
				updateTeam(game[0], -1)
				updateTeam(game[1], 1)
			case "draw":
				updateTeam(game[0], 0)
				updateTeam(game[1], 0)
			default:
				return errors.New("Invalid result")
			}
			line = ""
		default:
			line += fmt.Sprint(bufString[a])
		}
	}
	return nil
}
