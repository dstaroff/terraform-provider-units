/*
 * Copyright (c) 2024. Dmitry Starov
 * SPDX-License-Identifier: MPL-2.0
 */

package converter

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	base1000 float64 = 1000.0
	base1024 float64 = 1024.0
)

var (
	zero = types.NumberValue(big.NewFloat(0))

	Kibi = big.NewFloat(base1024)
	Mebi = new(big.Float).Mul(Kibi, Kibi)
	Gibi = new(big.Float).Mul(Mebi, Kibi)
	Tebi = new(big.Float).Mul(Gibi, Kibi)
	Pebi = new(big.Float).Mul(Tebi, Kibi)

	Kilo = big.NewFloat(base1000)
	Mega = new(big.Float).Mul(Kilo, Kilo)
	Giga = new(big.Float).Mul(Mega, Kilo)
	Tera = new(big.Float).Mul(Giga, Kilo)
	Peta = new(big.Float).Mul(Tera, Kilo)
)

type dataSizeConverter func(types.Number) types.Number

func bytesTo(coefficient *big.Float) dataSizeConverter {
	return func(number types.Number) types.Number {
		if number.Equal(zero) {
			return types.NumberValue(number.ValueBigFloat())
		}

		return types.NumberValue(
			new(big.Float).Quo(
				number.ValueBigFloat(),
				coefficient,
			),
		)
	}
}

func toBytes(coefficient *big.Float) dataSizeConverter {
	return func(number types.Number) types.Number {
		if number.Equal(zero) {
			return types.NumberValue(number.ValueBigFloat())
		}

		return types.NumberValue(
			new(big.Float).Mul(
				number.ValueBigFloat(),
				coefficient,
			),
		)
	}
}

var (
	KibibytesFromBytes = bytesTo(Kibi)
	MebibytesFromBytes = bytesTo(Mebi)
	GibibytesFromBytes = bytesTo(Gibi)
	TebibytesFromBytes = bytesTo(Tebi)
	PebibytesFromBytes = bytesTo(Pebi)

	KibibytesToBytes = toBytes(Kibi)
	MebibytesToBytes = toBytes(Mebi)
	GibibytesToBytes = toBytes(Gibi)
	TebibytesToBytes = toBytes(Tebi)
	PebibytesToBytes = toBytes(Pebi)

	KilobytesFromBytes = bytesTo(Kilo)
	MegabytesFromBytes = bytesTo(Mega)
	GigabytesFromBytes = bytesTo(Giga)
	TerabytesFromBytes = bytesTo(Tera)
	PetabytesFromBytes = bytesTo(Peta)

	KilobytesToBytes = toBytes(Kilo)
	MegabytesToBytes = toBytes(Mega)
	GigabytesToBytes = toBytes(Giga)
	TerabytesToBytes = toBytes(Tera)
	PetabytesToBytes = toBytes(Peta)
)
