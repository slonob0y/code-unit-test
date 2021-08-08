package services

// jika nilai <= 40, maka nilai C
// jika nilai > 40 dan <= 70, maka nilai B
// jika nilai > 70, maka nilai C

func grade(nilai int) string {

	if nilai <= 40 {
		return "C"
	} else if nilai > 40 && nilai <= 70 {
		return "B"
	} else {
		return "A"
	}
}
