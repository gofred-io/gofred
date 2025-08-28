package utils

import "runtime/debug"

func SafeGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
			}
		}()

		f()
	}()
}
