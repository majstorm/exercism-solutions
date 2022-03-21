package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

var TeamBoard map[string]*Team

type Team struct {
	MP, W, D, L, P int
}

func updateTeam(teamName string, gameUpdate int) {
	_, ok := TeamBoard[teamName]
	if !ok {
		TeamBoard[teamName] = &Team{MP: 0, W: 0, D: 0, L: 0, P: 0}
	}

	TeamBoard[teamName].MP += 1
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
	TeamBoard = make(map[string]*Team)
	buf := bufio.NewScanner(reader)

	for buf.Scan() {
		line := buf.Text()
		if strings.TrimSpace(line) == "" || strings.Contains(line, "#") {
			continue
		}
		game := strings.Split(line, ";")
		if len(game) != 3 {
			return errors.New("Incorrect number of columns")
		}
		home, away, res := game[0], game[1], game[2]

		switch res {
		case "win":
			updateTeam(home, 1)
			updateTeam(away, -1)
		case "loss":
			updateTeam(home, -1)
			updateTeam(away, 1)
		case "draw":
			updateTeam(home, 0)
			updateTeam(away, 0)
		default:
			return errors.New("Invalid result")
		}
	}

	type kv struct {
		Key   string
		Value *Team
	}
	var ss []kv
	for k, v := range TeamBoard {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value.P == ss[j].Value.P {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value.P > ss[j].Value.P
	})

	fmt.Fprintln(writer, `Team                           | MP |  W |  D |  L |  P`)
	for _, kv := range ss {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", kv.Key, kv.Value.MP, kv.Value.W, kv.Value.D, kv.Value.L, kv.Value.P)
	}

	return nil
}
