package main

func lexNumeric(source string, ic cursor) (*token, cursor, bool) {
	cur := ic // define cur to being the current cursor
	// cursor has a pointer, and a location object
	periodFound := false // if we have found a period that means that it is a float
	expMarkerFound := false // if we have found an exponential marker then we want to return
	// we are going to iterate through the current pointer up to the source len
	for ; cur.pointer < uint(len(source)); cur.pointer++ {
		// iterate from the current pointer until the end of the source

		c := source[cur.pointer]
		cur.loc.col++
		// update the location of the cursor to keep track of the values
		// c is defined to be to being the character at the srouce string
		// we are then going to up the cur.loc.col which means we are moving rihgt

		isDigit := c >= '0' && c <= '9' // check if its a digit or if it is etc
		isPeriod := c == '.'
		isExpMarker := c == 'e'

		// Must start with a digit or period
		if cur.pointer == ic.pointer { // if we have not moved from the current ointer?
			if !isDigit && !isPeriod {
				return nil, ic, false
			}

			periodFound = isPeriod
			continue
		}

		if isPeriod {
			if periodFound {
				return nil, ic, false
			}

			periodFound = true
			continue
		}

		if isExpMarker {
			if expMarkerFound {
				return nil, ic, false
			}

			// No periods allowed after expMarker
			periodFound = true
			expMarkerFound = true

			// expMarker must be followed by digits
			if cur.pointer == uint(len(source)-1) {
				return nil, ic, false
			}

			cNext := source[cur.pointer+1]
			if cNext == '-' || cNext == '+' {
				cur.pointer++
				cur.loc.col++
			}

			continue
		}

		if !isDigit {
			break
		}
	}

	// No characters accumulated
	if cur.pointer == ic.pointer {
		return nil, ic, false
	}

	return &token{
		value: source[ic.pointer:cur.pointer],
		loc:   ic.loc,
		kind:  numericKind,
	}, cur, true
}