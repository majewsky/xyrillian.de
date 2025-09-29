// SPDX-FileCopyrightText: 2025 Stefan Majewsky <majewsky@gmx.net>
// SPDX-License-Identifier: GPL-3.0-only

package util

import "log"

func MustSucceed(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MustReturn[T any](value T, err error) T {
	MustSucceed(err)
	return value
}
