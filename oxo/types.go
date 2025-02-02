// Data types go here as the need arises to define them.

package oxo

type GameStateclass uint8 //Holds 16 permutations of 4 GameStates

type GameState uint8 //Holds 4 GameStates OWIN,XWIN,INPLAY and LEGAL

type Gridpos uint8 //Holds the position in a 3x3 grid from topleft 0 to bottom right 8

// These are the 4 GameStates, a Board can have multiple GameStates
const (
	IsInplay GameState = 1 << iota
	IsOwin
	IsXwin
	IsLegal
)

// These are all 16 permutations of 4 GameStates so there are NO overlaps.  Any board is in only one of these.
const (
	Classlxop0000 GameStateclass = iota //ILLEGAL
	ClasslxoP0001                       //ILLEGAL
	ClasslxOp0010                       //ILLEGAL
	ClasslxOP0011                       //ILLEGAL
	Classlxop0100                       //ILLEGAL
	ClasslxoP0101                       //ILLEGAL
	ClasslxOp0110                       //ILLEGAL
	ClasslxOP0111                       //ILLEGAL
	Classlxop1000                       //DRAW  Legal board Neither O or X win and game is full
	ClasslxoP1001                       //PLAY  Legal board Neither O or X win and game is in play
	ClasslxOp1010                       //OWIN  Legal board OWins and game is full.  OWin, not a DRAW
	ClasslxOP1011                       //OWIN  Legal board OWins and game is in play
	Classlxop1100                       //XWIN  Legal board XWins and game is full.  XWin, not a DRAW
	ClasslxoP1101                       //XWIN  Legal board XWins and game is in play
	ClasslxOp1110                       //ERROR Legal board but has both X AND O winning lines and is full. Error because one should be detected first.
	ClasslxOP1111                       //ERROR Legal board but has both X and O winning line and is in play.  Error because one should be detected first.
)
const (
	X = byte(88)
	O = byte(79)
	S = byte(32)
)
const (
	Tl Gridpos = iota //Topleft = 0
	Tc                //Topcentre
	Tr                //Toperight
	Ml                //Midleft
	Mc                //Midcentre
	Mr                //Midright
	Bl                //Bottomleft
	Bc                //Bottomcentre
	Br                //Bottomright =8
)

// board represents the state of a turn of the game on a 3 x 3 grid

type Player struct {
	Name   string         `json:"name"`
	Tactic func(Grid) int `json:"-"` //Tactic function for this player.
	Rank   int            `json:"rank"`
}

type Grid [9]byte

// a turn has a board and status string which will be one of these strings XWIN,OWIN,DRAW,PLAY
type Turn struct {
	Board  Grid   `json:"grid"`
	Status string `json:"status"`
}

// a complete game has a minimum of 5 turns and a maximum of 9 turns
// should I set a capacity for the slice?
type Game struct {
	X      Player `json:"xplayer"`
	O      Player `json:"oplayer"`
	Result string `json:"result"`
	Turns  []Turn `json:"turn"`
}

type Group struct {
	Games       []Game `json:"games"`
	NumOwins    int    `json:"numowins"`
	NumXwins    int    `json:"numxwins"`
	NumDraws    int    `json:"draws"`
	NumIllegals int    `json:"illegals"`
	NumTurns    int    `json:"numturns"`
}

func (r *Group) UpdateNums(g Game) {
	switch {
	case g.Result == "XWIN":
		r.NumXwins++
	case g.Result == "OWIN":
		r.NumOwins++
	case g.Result == "DRAW":
		r.NumDraws++
	case g.Result == "ILLEGAL":
		r.NumIllegals++
	}
	r.NumTurns = len(g.Turns)
	return
}

type Lookup map[string]string

type Tacticfunc func(Grid) int

func (board Grid) FindSpaces() ([]byte, bool) {
	var spc []byte
	for k, v := range board {
		if v == S {
			spc = append(spc, byte(k))
		}
	}
	return spc, len(spc) != 0
}
