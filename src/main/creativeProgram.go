//creative_prgram.go
//Jessica Forrett and Alex Brisbois
//CSCI 324 Professor King

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//stores facts about the planet
type Planet struct {
	name           string  //the name of the planet
	eccentricity   float64 //the eccentricity (or flattening) of the planet's orbit
	axis           float64 //the semimajor axis of the planet's orbit (in au)
	orbital_period float64 //how long the planet's orbit is (in days)
	pyear          int     //The year when the planet is at perihelion
	pmonth         int     //The month when the planet is at perihelion
	pday           int     //The day when the planet is at perihelion
}

//implementing primitive data type methods
type Float struct {
	value float64
}

//returns the Float value as a string
func (f Float) FtoS() string {
	return strconv.FormatFloat(f.value, 'f', -1, 64)
}

//calculates the planet's current distance from the sun and prints it out
func sunDistance(year int, month int, day int, plnt Planet, c chan bool) {
	var days float64 = float64(year-plnt.pyear)*365.25 + float64(month-plnt.pmonth)*30.4375 + float64(day-plnt.pday)
	r := Float{plnt.axis * (1 - plnt.eccentricity*plnt.eccentricity) / (1 + plnt.eccentricity*math.Cos(((days*360)/plnt.orbital_period)/180*math.Pi))}
	fmt.Println("The current distance from " + plnt.name + " to the sun is " + r.FtoS() + " AU")
	c <- true
}

//error checking
func check(e error) {
	if e != nil {
		fmt.Println("An error has occurred")
		panic(e)
	}
}

//returns the orbital period of the planet
func orbit(p Planet, c chan float64) {
	orbit_per := p.orbital_period

	c <- orbit_per
}

func main() {

	//declare all of the planets
	mercury := Planet{"Mercury", 0.206, 0.3871, 88, 2017, 6, 19}
	venus := Planet{"Venus", 0.007, 0.7233, 225, 2017, 2, 20}
	earth := Planet{"Earth", 0.017, 1.0, 365.26, 2017, 1, 4}
	mars := Planet{"Mars", 0.093, 1.5273, 693.99, 2016, 10, 29}
	jupiter := Planet{"Jupiter", 0.048, 5.2028, 4346.59, 2023, 1, 20}
	saturn := Planet{"Saturn", 0.056, 9.5388, 7815.17, 2003, 7, 9}
	uranus := Planet{"Uranus", 0.047, 19.1914, 30681.84, 1966, 7, 16}
	neptune := Planet{"Neptune", 0.009, 30.0611, 60194.85, 2046, 2, 18}
	pluto := Planet{"Pluto", 0.248, 39.48, 90584.48, 1989, 9, 5}

	planet_list := [9]Planet{
		mercury,
		venus,
		earth,
		mars,
		jupiter,
		saturn,
		uranus,
		neptune,
		pluto,
	}

	//variables that the user will input
	ans := ""
	dayS := ""
	monthS := ""
	yearS := ""
	month := 0
	day := 0
	year := 0
	endMonth := 0
	planet := ""
	var plnt Planet

	fmt.Println("")
	fmt.Println("")
	fmt.Println("                  Welcome to the Planetarium!")
	fmt.Println("")
	fmt.Println("          _..._")
	fmt.Println("        .'     '.      _")
	fmt.Println("       /    .-\"\"-\\   _/ \\")
	fmt.Println("     .-|   /:.   |  |   |")
	fmt.Println("     |  \\  |:.   /.-'-./")
	fmt.Println("     | .-'-;:__.'    =/")
	fmt.Println("     .'=  *=|NASA _.='")
	fmt.Println("    /   _.  |    ;")
	fmt.Println("   ;-.-'|    \\   |")
	fmt.Println("  /   | \\    _\\  _\\")
	fmt.Println("  \\__/'._;.  ==' ==\\")
	fmt.Println("           \\    \\   |")
	fmt.Println("           /    /   /")
	fmt.Println("           /-._/-._/")
	fmt.Println("           \\   `\\  \\")
	fmt.Println("            `-._/._/")
	fmt.Println("")
	fmt.Println("")

	//User says if they want to continue with the program
	for ans != "yes" && ans != "no" {
		fmt.Println("Would you like to see the current distance of a planet from the sun? (yes/no)")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		ans = scan.Text()
	}

	//exit if not
	if ans == "no" {
		fmt.Println("Okay, bye!")
		return
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                     ;   :   ;")
	fmt.Println("                  .   \\_,!,_/   ,")
	fmt.Println("                   `.,'     `.,'")
	fmt.Println("                    /         \\")
	fmt.Println("               ~ -- :         : -- ~")
	fmt.Println("                    \\         /")
	fmt.Println("                   ,'`._   _.'`.")
	fmt.Println("                  '   / `!` \\   `")
	fmt.Println("                     ;   :   ;")
	fmt.Println("")
	fmt.Println("")

	//User enters in the year
	fmt.Println("")
	for year == 0 {
		fmt.Println("Okay! Please enter in the current year")
		Yscan := bufio.NewScanner(os.Stdin)
		Yscan.Scan()
		yearS = Yscan.Text()
		year, _ = strconv.Atoi(yearS)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("         _.._")
	fmt.Println("       .' .-'`")
	fmt.Println("      /  /")
	fmt.Println("      |  |")
	fmt.Println("      \\  \\")
	fmt.Println("       '._'-._")
	fmt.Println("          ```")
	fmt.Println("")

	//User enters in the month
	fmt.Println("")
	for month < 1 || month > 12 {
		fmt.Println("Okay! Please enter in the current month (1-12)")
		Mscan := bufio.NewScanner(os.Stdin)
		Mscan.Scan()
		monthS = Mscan.Text()
		month, _ = strconv.Atoi(monthS)
	}

	//Find the last possible day of the month
	switch month {
	case 1:
		endMonth = 31
	case 2:
		if year%4 == 0 {
			endMonth = 29
		} else {
			endMonth = 28
		}
	case 3:
		endMonth = 31
	case 4:
		endMonth = 30
	case 5:
		endMonth = 31
	case 6:
		endMonth = 30
	case 7:
		endMonth = 31
	case 8:
		endMonth = 31
	case 9:
		endMonth = 30
	case 10:
		endMonth = 31
	case 11:
		endMonth = 30
	case 12:
		endMonth = 31
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("				        _____")
	fmt.Println("				    ,-:` \\;',`'-, ")
	fmt.Println("				  .'-;_,;  ':-;_,'.")
	fmt.Println("				 /;   '/    ,  _`.-\\")
	fmt.Println("				| '`. (`     /` ` \\`|")
	fmt.Println("				|:.  `\\`-.   \\_   / |")
	fmt.Println("				|     (   `,  .`\\ ;'|")
	fmt.Println(" 				\\     | .'     `-'/")
	fmt.Println("				  `.   ;/        .'")
	fmt.Println("				    `'-._____.")
	fmt.Println("")

	//User enters in the day
	fmt.Println("")
	for day < 1 || day > endMonth {
		fmt.Println("Okay! Please enter in the current day")
		Dscan := bufio.NewScanner(os.Stdin)
		Dscan.Scan()
		dayS = Dscan.Text()
		day, _ = strconv.Atoi(dayS)
	}

	//receive input for the planet
	done := false
	fmt.Println("")
	for true {
		fmt.Println("Which planet's distance would you like us to calculate for you (The planet must be capitalized)?")
		Pscan := bufio.NewScanner(os.Stdin)
		Pscan.Scan()
		planet = Pscan.Text()
		for i := 0; i < 9; i++ {
			if planet == planet_list[i].name {
				plnt = planet_list[i]
				e := planet_list[i].eccentricity
				a := planet_list[i].axis

				//plug e and a into formula
				perihelion_dist := a * (1 - e)
				aphelion_dist := a * (1 + e)

				//outputs the distance when the planet is closest to the sun
				fmt.Println("")
				fmt.Print(planet + "'s perihelion distance is ")
				fmt.Println(perihelion_dist)

				//outputs the distance when the planet is farthest from the sun
				fmt.Print(planet + "'s aphelion distance is ")
				fmt.Println(aphelion_dist)
				fmt.Println("")
				done = true
				break
			}

		}
		if done {
			break
		}
	}

	//finds the distance from the planet to the sun
	cs := make(chan bool, 1)
	sunDistance(year, month, day, plnt, cs)

	//see if the user wants to see the distance of all planets from the sun
	ans = ""
	fmt.Println("")
	fmt.Println("Would you like to see the all the planets' orbital periods and current distance from the sun,? (yes/no)")
	for ans != "yes" && ans != "no" {
		fmt.Scanln(&ans)
	}

	//exit if not
	if ans == "no" {
		fmt.Println("Okay, bye!")
		return
	}

	//find all of the planets' orbitary periods
	c := make(chan float64, 3)
	go orbit(planet_list[0], c)
	go orbit(planet_list[1], c)
	ven_orbit, merc_orbit := <-c, <-c

	go orbit(planet_list[2], c)
	go orbit(planet_list[3], c)
	mar_orbit, ear_orbit := <-c, <-c

	go orbit(planet_list[4], c)
	go orbit(planet_list[5], c)
	sat_orbit, jup_orbit := <-c, <-c

	go orbit(planet_list[6], c)
	go orbit(planet_list[7], c)
	go orbit(planet_list[8], c)

	plu_orbit, ur_orbit, nep_orbit := <-c, <-c, <-c

	//outputs the orbital periods of each planet
	fmt.Println("")
	fmt.Print("Mercury's Orbital Period: ")
	fmt.Print(merc_orbit)
	fmt.Println(" days")
	fmt.Print("Venus's Orbital Period: ")
	fmt.Print(ven_orbit)
	fmt.Println(" days")
	fmt.Print("Earth's Orbital Period: ")
	fmt.Print(ear_orbit)
	fmt.Println(" days")
	fmt.Print("Mars's Orbital Period: ")
	fmt.Print(mar_orbit)
	fmt.Println(" days")
	fmt.Print("Jupiter's Orbital Period: ")
	fmt.Print(jup_orbit)
	fmt.Println(" days")
	fmt.Print("Saturn's Orbital Period: ")
	fmt.Print(sat_orbit)
	fmt.Println(" days")
	fmt.Print("Uranus's Orbital Period: ")
	fmt.Print(ur_orbit)
	fmt.Println(" days")
	fmt.Print("Neptune's Orbital Period: ")
	fmt.Print(nep_orbit)
	fmt.Println(" days")
	fmt.Print("Pluto's Orbital Period: ")
	fmt.Print(plu_orbit)
	fmt.Println(" days")
	fmt.Println("")

	ch := make(chan bool, 9)
	//Go through the planets and calculate the distance from each one
	for i := 0; i < 9; i++ {
		go sunDistance(year, month, day, planet_list[i], ch)
	}

	//values are recieved in a randomized order
	//even when the user enters the same information twice, the order of the planets is different
	me_done, ve_done, ea_done, ma_done, ju_done, sa_done, ur_done, ne_done, pl_done := <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch, <-ch

	//will print this out and exit the program once all the distances are calculated
	if me_done && ve_done && ea_done && ma_done && ju_done && sa_done && ur_done && ne_done && pl_done {
		fmt.Println("Have a nice day!")
	}

}
