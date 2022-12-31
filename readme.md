
this is a a project based on a puzzle game i was given. the point is to test if the game has any configuration which is unsolvable.

<img src="doc/puzzle.jpg" width="400">

Currently, it can solve a set of stops (ie the 3 pieces you move around to make a puzzle) and also roll through and try all stops.

```bash
$ go run main.go solve --help
solve for default pieces

Usage:
  pg-puzzel solve [flags]

Flags:
  -a, --all            try every stop combination
  -h, --help           help for solve
  -s, --stops string   board stops to solve, '[0-4],[0-4] [0-4],[0-4] [0-4],[0-4]' (default "0,0 0,4 4,2")
```

I used this as an excuse to learn features of golang, lots can be improved in the code and the code style.
