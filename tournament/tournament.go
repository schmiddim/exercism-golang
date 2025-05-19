package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name string
	mp   int
	w    int
	d    int
	l    int
	p    int
}

type Tournament struct {
	teams []*Team
}

func NewTournament() *Tournament {
	return &Tournament{}
}

func (t *Tournament) getTeamByName(name string) *Team {
	for _, t := range t.teams {
		if t.name == name {
			return t
		}
	}
	return nil
}

func (t *Tournament) ParseLine(line string) error {
	if line == "" {
		return nil
	}
	items := strings.Split(line, ";")

	if len(items) != 3 && line[0] != '#' {
		return errors.New("invalid line")
	}
	teamA := t.getTeamByName(items[0])
	teamB := t.getTeamByName(items[1])

	if teamA == nil {

		t.teams = append(t.teams, &Team{
			name: items[0],
		})
	}
	if teamB == nil {

		t.teams = append(t.teams, &Team{
			name: items[1],
		})

	}
	teamA = t.getTeamByName(items[0])
	teamB = t.getTeamByName(items[1])
	outcome := items[2]

	switch outcome {
	case "win":
		teamA.w += 1
		teamB.l += 1
		teamA.p += 3
		teamA.mp += 1
		teamB.mp += 1
	case "loss":
		teamB.w += 1
		teamA.l += 1
		teamB.p += 3
		teamA.mp += 1
		teamB.mp += 1
	case "draw":
		teamB.d += 1
		teamA.d += 1
		teamA.p += 1
		teamB.p += 1
		teamA.mp += 1
		teamB.mp += 1
	default:
		return errors.New("invalid outcome")
	}
	fmt.Println(teamA, teamB, outcome)
	return nil
}
func (t *Tournament) sortByPoints() {
	sort.Slice(t.teams, func(i, j int) bool {
		ti, tj := t.teams[i], t.teams[j]

		if ti.p != tj.p {
			return ti.p > tj.p // mehr Punkte zuerst
		}
		return ti.name < tj.name // Name als Tie-Breaker
	})
}
func Tally(reader io.Reader, writer io.Writer) error {
	var lines []string
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	t := NewTournament()
	for _, line := range lines {
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		err := t.ParseLine(line)
		if err != nil {
			return err
		}

	}
	var b strings.Builder // = io.Writer + direkt String()

	b.WriteString(header)
	t.sortByPoints()
	for _, t := range t.teams {
		writeLine(&b, *t)
	}
	writer.Write([]byte(b.String()))
	return nil
}

const header = "Team                           | MP |  W |  D |  L |  P\n"

func writeLine(w io.Writer, t Team) {
	fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
		t.name, t.mp, t.w, t.d, t.l, t.p)
}
