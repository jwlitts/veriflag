package veriflag

import (
	"fmt"
	"os"
	"strconv"
)

func BoolFlag(flag string, defval bool, description string) func() (bool, error) {
	return func() (bool, error) {
		if len(os.Args) == 1 {
			return defval, nil
		}
		matchflag := "-" + flag
		for _, arg := range os.Args {
			if arg == matchflag {
				return true, nil
			}
		}
		return false, nil
	}
}

func BoolArg(position int, defval bool, description string) func() (bool, error) {
	return func() (bool, error) {
		if len(os.Args) <= position {
			return defval, nil
		} else {
			parsed, err := strconv.ParseBool(os.Args[position])
			return parsed, err
		}
	}
}

func FloatFlag(flag string, defval float64, condition func(x float64) bool, description string) func() (float64, error) {
	return func() (float64, error) {
		if len(os.Args) == 1 {
			return defval, nil
		}
		matchflag := "-" + flag
		for ii, arg := range os.Args {
			if arg == matchflag && len(os.Args) > ii {
				parsed, err := strconv.ParseFloat(os.Args[ii+1], 64)
				if err != nil {
					return defval, err
				}
				if condition(parsed) {
					return parsed, nil
				} else {
					return parsed, fmt.Errorf("invalid input value %f for flag %s", parsed, flag)
				}
			}
		}
		return defval, nil
	}
}

func FloatArg(position int, defval float64, condition func(x float64) bool, description string) func() (float64, error) {
	return func() (float64, error) {
		if len(os.Args) <= position {
			return defval, nil
		} else {
			parsed, err := strconv.ParseFloat(os.Args[position], 64)
			if err != nil {
				return defval, err
			}
			if condition(parsed) {
				return parsed, nil
			} else {
				return parsed, fmt.Errorf("invalid input value %f for arg at position %d", parsed, position)
			}
		}
	}
}

func IntFlag(flag string, defval int64, condition func(x int64) bool, description string) func() (int64, error) {
	return func() (int64, error) {
		if len(os.Args) == 1 {
			return defval, nil
		}
		matchflag := "-" + flag
		for ii, arg := range os.Args {
			if arg == matchflag && len(os.Args) > ii {
				parsed, err := strconv.ParseInt(os.Args[ii+1], 10, 64)
				if err != nil {
					return defval, err
				}
				if condition(parsed) {
					return parsed, nil
				} else {
					return parsed, fmt.Errorf("invalid input value %d for flag %s", parsed, flag)
				}
			}
		}
		return defval, nil
	}
}

func IntArg(position int, defval int64, condition func(x int64) bool, description string) func() (int64, error) {
	return func() (int64, error) {
		if len(os.Args) <= position {
			return defval, nil
		} else {
			parsed, err := strconv.ParseInt(os.Args[position], 10, 64)
			if err != nil {
				return defval, err
			}
			if condition(parsed) {
				return parsed, nil
			} else {
				return parsed, fmt.Errorf("invalid input value %d for arg at position %d", parsed, position)
			}
		}
	}
}

func StringFlag(flag string, defval string, condition func(x string) bool, description string) func() (string, error) {
	return func() (string, error) {
		if len(os.Args) == 1 {

			return defval, nil
		}
		matchflag := "-" + flag
		for ii, arg := range os.Args {
			if arg == matchflag && len(os.Args) > ii {
				parsed := os.Args[ii+1]
				if condition(parsed) {
					return parsed, nil
				} else {
					return parsed, fmt.Errorf("invalid input value %s for flag %s", parsed, flag)
				}
			}
		}
		return defval, nil
	}
}

func StringArg(position int, defval string, condition func(x string) bool, description string) func() (string, error) {
	return func() (string, error) {
		if len(os.Args) <= position {
			return defval, nil
		} else {
			parsed := os.Args[position]
			if condition(parsed) {
				return parsed, nil
			} else {
				return parsed, fmt.Errorf("invalid input value %s for arg at position %d", parsed, position)
			}
		}
	}
}
