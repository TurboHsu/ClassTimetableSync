package config

type ConfigStruct struct {
	UserInfo struct {
		ClientID string
	}
	Init struct {
		Timezone     string
		FirstWeekDay string
		Replacement  []struct {
			Limiter  string
			Original string
			Replace  string
		}
	}
}
