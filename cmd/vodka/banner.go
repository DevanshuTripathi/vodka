package main

import (
	"fmt"
	"strings"
	"time"
)

func printPouringAnimation(msg string) {
	frames := []string{
		// Frame 0: Empty Rocks Glass
		`                            
                            
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + ` \                  / ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 1: Stream pouring
		`             ` + Cyan + `|  |` + Reset + `         
             ` + Cyan + `|  |` + Reset + `         
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + ` \       ` + Cyan + `|  |` + Gray + `       / ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 2: Splash at the bottom
		`             ` + Cyan + `|  |` + Reset + `         
             ` + Cyan + `|  |` + Reset + `         
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|       ` + Gray + `.` + Cyan + `|  |` + Gray + `.` + Gray + `       |` + Reset + `
        ` + Gray + ` \     ` + Gray + `o ` + Cyan + `\__/` + Gray + ` o` + Gray + `     / ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 3: Filling, turbulent sloshing
		`             ` + Cyan + `|  |` + Reset + `         
             ` + Cyan + `|  |` + Reset + `         
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|       ` + Gray + `.` + Cyan + `|  |` + Gray + `.` + Gray + `       |` + Reset + `
        ` + Gray + `|     ` + Gray + `o ` + Cyan + `~====~` + Gray + ` o` + Gray + `     |` + Reset + `
        ` + Gray + `|    ` + Cyan + `~==========~` + Gray + `    |` + Reset + `
        ` + Gray + ` \ ` + Cyan + `================` + Gray + ` / ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 4: More filling, heavy waves
		`             ` + Cyan + `|  |` + Reset + `         
             ` + Cyan + `|  |` + Reset + `         
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|      ` + Gray + `. ` + Cyan + `|  |` + Gray + ` .` + Gray + `      |` + Reset + `
        ` + Gray + `|    ` + Gray + `o ` + Cyan + `~======~` + Gray + ` o` + Gray + `    |` + Reset + `
        ` + Gray + `|   ` + Cyan + `~==============~` + Gray + `   |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 5: Stream tapers off, breaking into drops
		`             ` + Cyan + `:  :` + Reset + `         
             ` + Cyan + `.  .` + Reset + `         
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Cyan + `|  |` + Gray + `        |` + Reset + `
        ` + Gray + `|      ` + Gray + `. ` + Cyan + `|  |` + Gray + ` .` + Gray + `      |` + Reset + `
        ` + Gray + `|   ` + Gray + `o ` + Cyan + `~========~` + Gray + ` o` + Gray + `   |` + Reset + `
        ` + Gray + `| ` + Cyan + `~================~` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 6: Stream ends, ice cube drops from above
		`                 ` + Gray + `[--]` + Reset + `       
                            
        ` + Gray + `|        ` + Gray + `o  o` + Gray + `        |` + Reset + `
        ` + Gray + `|       ` + Gray + `.    .` + Gray + `       |` + Reset + `
        ` + Gray + `| ` + Cyan + `~=====` + Gray + `~~~~~~` + Cyan + `=====~` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 7: Ice falling rapidly
		`                            
                 ` + Gray + `[--]` + Reset + `       
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `| ` + Cyan + `~================~` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 8: Ice hits liquid - HUGE splash
		`                            
                            
        ` + Gray + `|         ` + Gray + `. .` + Gray + `        |` + Reset + `
        ` + Gray + `|        ` + Gray + `o | o` + Gray + `       |` + Reset + `
        ` + Gray + `| ` + Cyan + `~====` + Gray + `\_[--]_/` + Cyan + `====~` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `======` + Gray + `~~~~~~` + Cyan + `======` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 9: Ice settles, waves roll outward
		`                            
                            
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `| ` + Cyan + `~~~~~~~` + Gray + `[--]` + Cyan + `~~~~~~~` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `========` + Gray + `\/` + Cyan + `========` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,

		// Frame 10: Final state, liquid goes flat and calm
		`                            
                            
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `|                    |` + Reset + `
        ` + Gray + `| ` + Cyan + `_______` + Gray + `[--]` + Cyan + `_______` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `========` + Gray + `\/` + Cyan + `========` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + `| ` + Cyan + `==================` + Gray + ` |` + Reset + `
        ` + Gray + ` \` + Cyan + `==================` + Gray + `/ ` + Reset + `
        ` + Gray + `  '----------------'  ` + Reset,
	}

	// Tailored timing to make the physics feel real (fast pour, fast drop, slow settle)
	delays := []int{200, 150, 150, 150, 150, 150, 200, 100, 300, 300, 600}

	// Print exactly 10 newlines to make space for the 11-line animation
	fmt.Print(strings.Repeat("\n", 10))

	for i, frame := range frames {
		// Move cursor UP 10 lines and snap to the beginning of the line
		fmt.Print("\033[10A\r")

		fmt.Print(frame)
		time.Sleep(time.Duration(delays[i]) * time.Millisecond)
	}

	// The final output message, now with an ice cube emoji!
	fmt.Printf("\n\n  🧊 %s%s%s\n\n", Cyan, msg, Reset)
}
