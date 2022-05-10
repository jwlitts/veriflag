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

func FlagFromArgList(flag string, args []string) (string, bool) {
	for ii, arg := range args {
		if arg == flag && len(args) > ii {
			return args[ii+1], true
		}
	}
	return "", false
}

func PosArgFromArgList(position int, args []string) (string, bool) {
	if len(args) <= position {
		return "", false
	} else {
		return args[position], true
	}
}

func FloatFlag(flag string, defval float64, condition func(x float64) bool, description string) func() (float64, error) {
	return func() (float64, error) {
		matchflag := "-" + flag
		flagval, found := FlagFromArgList(matchflag, os.Args)
		if !found {
			return defval, nil
		}
		parsed, err := strconv.ParseFloat(flagval, 64)
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

func FloatArg(position int, defval float64, condition func(x float64) bool, description string) func() (float64, error) {
	return func() (float64, error) {
		val, found := PosArgFromArgList(position, os.Args)
		if !found {
			return defval, nil
		}
		parsed, err := strconv.ParseFloat(val, 64)
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

func IntFlag(flag string, defval int64, condition func(x int64) bool, description string) func() (int64, error) {
	return func() (int64, error) {
		matchflag := "-" + flag
		flagval, found := FlagFromArgList(matchflag, os.Args)
		if !found {
			return defval, nil
		}
		parsed, err := strconv.ParseInt(flagval, 10, 64)
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

func IntArg(position int, defval int64, condition func(x int64) bool, description string) func() (int64, error) {
	return func() (int64, error) {
		val, found := PosArgFromArgList(position, os.Args)
		if !found {
			return defval, nil
		}
		parsed, err := strconv.ParseInt(val, 10, 64)
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

func StringFlag(flag string, defval string, condition func(x string) bool, description string) func() (string, error) {
	return func() (string, error) {
		matchflag := "-" + flag
		parsed, found := FlagFromArgList(matchflag, os.Args)
		if !found {
			return defval, nil
		}
		if condition(parsed) {
			return parsed, nil
		} else {
			return parsed, fmt.Errorf("invalid input value %s for flag %s", parsed, flag)
		}
	}
}

func StringArg(position int, defval string, condition func(x string) bool, description string) func() (string, error) {
	return func() (string, error) {
		parsed, found := PosArgFromArgList(position, os.Args)
		if !found {
			return defval, nil
		}
		if condition(parsed) {
			return parsed, nil
		} else {
			return parsed, fmt.Errorf("invalid input value %s for arg at position %d", parsed, position)
		}
	}
}
