package main

func predictPartyVictory(senate string) string {
	remainR, remainD := 0, 0
	for i := range senate {
		if senate[i] == 'R' {
			remainR++
		} else {
			remainD++
		}
	}
	for remainR > 0 && remainD > 0 {
		tmp := []byte{}
		skipR, skipD := 0, 0
		for i := range senate {
			if senate[i] == 'R' {
				if skipR > 0 {
					skipR--
					remainR--
					continue
				}
				skipD++
			} else {
				if skipD > 0 {
					skipD--
					remainD--
					continue
				}
				skipR++
			}
			tmp = append(tmp, senate[i])
		}
		newSenate := []byte{}
		for i := range tmp {
			if tmp[i] == 'R' && skipR > 0 {
				skipR--
				remainR--
			} else if tmp[i] == 'D' && skipD > 0 {
				skipD--
				remainD--
			} else {
				newSenate = append(newSenate, tmp[i])
			}
		}
		senate = string(newSenate)
	}
	if remainD == 0 {
		return "Radiant"
	}
	return "Dire"
}
