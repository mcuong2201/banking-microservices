package app

func EvaluateRisk(transactionAmount float64) string {
	if transactionAmount > 10000 {
		return "HIGH_RISK"
	}
	return "LOW_RISK"
}
