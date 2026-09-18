package sprint

func Season(month string) string {
	switch month {
	case "jan":
		return "winter"
	case "feb":
		return "winter"
	case "dec":
		return "winter"
	case "mar":
		return "spring"
	case "apr":
		return "spring"
	case "may":
		return "spring"
	case "jun":
		return "summer"
	case "jul":
		return "summer"
	case "aug":
		return "summer"
	case "sep":
		return "autumn"
	case "oct":
		return "autumn"
	case "nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}
