package cards

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func TestFourDigits(t *testing.T) {
	Convey("Should be able to retrieve the last four digits of my card", t, func() {
		Convey("When I have four or more digits", func() {
			Convey("Using LastFour method", func() {
				card := Card{Number: "4012888888881881", Cvv: "111"}
				lastFour, err := card.LastFour()

				So(err, ShouldBeNil)
				So(lastFour, ShouldEqual, "1881")
			})

			Convey("Using LastFourDigits method", func() {
				card := Card{Number: "4012888888881881", Cvv: "111"}
				lastFour, err := card.LastFourDigits()

				So(err, ShouldBeNil)
				So(lastFour, ShouldEqual, "1881")
			})
		})

		Convey("When I don't have four or more digits", func() {
			card := Card{Number: "123", Cvv: "111"}
			lastFour, err := card.LastFour()

			So(err, ShouldNotBeNil)
			So(lastFour, ShouldEqual, "")
		})
	})
}

func TestWipe(t *testing.T) {
	Convey("Should be able to wipe our credit card", t, func() {
		card := Card{Number: "4012888888881881", Cvv: "111", Month: "02", Year: "2015"}
		card.Wipe()

		So(card.Number, ShouldEqual, "0000000000000000")
		So(card.Cvv, ShouldEqual, "0000")
		So(card.Month, ShouldEqual, "01")
		So(card.Year, ShouldEqual, "1970")
	})
}

func TestValidation(t *testing.T) {
	month := strconv.Itoa(int(time.Now().UTC().Month()))
	year := strconv.Itoa(time.Now().UTC().Year())

	Convey("Expiration should matter", t, func() {
		Convey("When the expiration year is less than the current year", func() {
			card := Card{Number: "4012888888881881", Cvv: "111", Month: "02", Year: "2001"}
			err := card.Validate(true)

			So(err, ShouldNotBeNil)
		})

		Convey("When the month is non-sensical", func() {
			nextYear, _, _ := (time.Now().UTC()).AddDate(1, 0, 0).Date()

			examples := []int{0, 13}
			for _, month := range examples {
				card := Card{Number: "4012888888881881", Cvv: "111", Month: strconv.Itoa(int(month)), Year: strconv.Itoa(nextYear)}
				err := card.Validate(true)

				So(err, ShouldNotBeNil)
			}
		})

		Convey("When the expiration month and year is less than the current date", func() {
			year1, month1, _ := (time.Now().UTC()).AddDate(0, -1, 0).Date()

			card := Card{Number: "4012888888881881", Cvv: "111", Month: strconv.Itoa(int(month1)), Year: strconv.Itoa(year1)}
			err := card.Validate(true)

			So(err, ShouldNotBeNil)
		})

		Convey("and should work when we use only two numbers", func() {
			card := Card{Number: "4012888888881881", Cvv: "111", Month: month, Year: year[2:]}
			err := card.Validate(true)

			So(err, ShouldBeNil)
		})
	})

	Convey("CVV validation", t, func() {
		Convey("Should work with three characters", func() {
			card := Card{Number: "4012888888881881", Cvv: "111", Month: month, Year: year}
			err := card.Validate(true)

			So(err, ShouldBeNil)
		})

		Convey("Should work with four characters", func() {
			card := Card{Number: "4012888888881881", Cvv: "1111", Month: month, Year: year}
			err := card.Validate(true)

			So(err, ShouldBeNil)
		})

		Convey("Should give us an error if CVV is invalid", func() {
			Convey("Too many numbers", func() {
				card := Card{Number: "4012888888881881", Cvv: "11111", Month: month, Year: year}
				err := card.Validate(true)

				So(err, ShouldNotBeNil)
			})

			Convey("Too little numbers", func() {
				card := Card{Number: "4012888888881881", Cvv: "11", Month: month, Year: year}
				err := card.Validate(true)

				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Credit card number length should matter", t, func() {
		Convey("Should give us an error if the length is less than 13 characters", func() {
			card := Card{Number: "1234", Cvv: "1111", Month: month, Year: year}
			err := card.Validate(true)

			So(err, ShouldNotBeNil)
		})

		Convey("Should not give us an error if the nuber is greater than or equal to 13 characters", func() {
			card := Card{Number: "4012888888881881", Cvv: "1111", Month: month, Year: year}
			err := card.Validate(true)

			So(err, ShouldBeNil)
		})
	})

	Convey("Test cards", t, func() {
		numbers := []string{"4242424242424242", "4012888888881881", "4000056655665556", "5555555555554444", "5200828282828210", "5105105105105100", "378282246310005", "371449635398431", "6011111111111117", "6011000990139424", "30569309025904", "38520000023237", "3530111333300000", "3566002020360505"}

		Convey("should pass through", func() {
			for _, num := range numbers {
				Convey(num, func() {
					card := Card{Number: num, Cvv: "1111", Month: month, Year: year}
					err := card.Validate(true)

					So(err, ShouldBeNil)
				})
			}
		})

		Convey("should not pass through", func() {
			for _, num := range numbers {
				Convey(num, func() {
					card := Card{Number: num, Cvv: "1111", Month: month, Year: year}
					err := card.Validate()

					So(err, ShouldNotBeNil)
				})
			}
		})
	})

	Convey("Should be able to validate a number with the Luhn algorithm", t, func() {
		Convey("With a valid card", func() {
			Convey("Test Card", func() {
				card := Card{Number: "4242424242424242", Cvv: "1111", Month: month, Year: year}
				err := card.Validate(true)

				So(err, ShouldBeNil)
			})

			Convey("Real Card", func() {
				card := Card{Number: "4556974850403706", Cvv: "1111", Month: month, Year: year}
				err := card.Validate()

				So(err, ShouldBeNil)
			})
		})

		Convey("With an invalid card", func() {
			card := Card{Number: "42483272242424242", Cvv: "1111", Month: month, Year: year}
			err := card.Validate(true)

			So(err, ShouldNotBeNil)
		})

		Convey("Not enough numbers", func() {
			card := Card{Number: "424832", Cvv: "1111", Month: month, Year: year}
			err := card.ValidateNumber()

			So(err, ShouldBeFalse)
		})
	})
}

func TestMethod(t *testing.T) {
	month := strconv.Itoa(int(time.Now().UTC().Month()))
	year := strconv.Itoa(time.Now().UTC().Year())

	Convey("Card method should validate even when there's less than 13 characters", t, func() {
		Convey("Should work for American Express", func() {
			card := Card{Number: "3782822463", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "amex")
			So(card.Company.Name, ShouldEqual, "American Express")
		})

		Convey("Should work for Elo", func() {
			card := Card{Number: "50673090006418", Cvv: "111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "elo")
			So(card.Company.Name, ShouldEqual, "Elo")
		})

		Convey("Should work for Bankcard", func() {
			Convey("5610", func() {
				card := Card{Number: "56101111111", Cvv: "1111", Month: month, Year: year}
				err := card.Brand()

				So(err, ShouldBeNil)
				So(card.Company.Code, ShouldEqual, "bankcard")
				So(card.Company.Name, ShouldEqual, "Bankcard")
			})

			Convey("560221 - 560225", func() {
				card := Card{Number: "56022111111", Cvv: "1111", Month: month, Year: year}
				err := card.Brand()

				So(err, ShouldBeNil)
				So(card.Company.Code, ShouldEqual, "bankcard")
				So(card.Company.Name, ShouldEqual, "Bankcard")
			})
		})

		Convey("Should work for Diners club", func() {
			Convey("Carte blanche", func() {
				card := Card{Number: "300221111111111", Cvv: "1111", Month: month, Year: year}
				err := card.Brand()

				So(err, ShouldBeNil)
				So(card.Company.Code, ShouldEqual, "diners club carte blanche")
				So(card.Company.Name, ShouldEqual, "Diners")
			})

			Convey("Club enRoute", func() {
				card := Card{Number: "20142111111111", Cvv: "1111", Month: month, Year: year}
				err := card.Brand()

				So(err, ShouldBeNil)
				So(card.Company.Code, ShouldEqual, "diners club enroute")
				So(card.Company.Name, ShouldEqual, "Diners")
			})

			Convey("Club international", func() {
				card := Card{Number: "3002111111111", Cvv: "1111", Month: month, Year: year}
				err := card.Brand()

				So(err, ShouldBeNil)
				So(card.Company.Code, ShouldEqual, "diners club international")
				So(card.Company.Name, ShouldEqual, "Diners")
			})
		})

		Convey("Should work for China UnionPay", func() {
			card := Card{Number: "62111111111", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "china unionpay")
			So(card.Company.Name, ShouldEqual, "China UnionPay")
		})

		Convey("Should work for Dankort", func() {
			card := Card{Number: "501955555", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "dankort")
			So(card.Company.Name, ShouldEqual, "Dankort")
		})

		Convey("Should work for Diners Club", func() {
			card := Card{Number: "30569309025", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "diners club international")
			So(card.Company.Name, ShouldEqual, "Diners")
		})

		Convey("Should work for Discover", func() {
			card := Card{Number: "60111111111", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "discover")
			So(card.Company.Name, ShouldEqual, "Discover")
		})

		Convey("Should work for InterPayment", func() {
			card := Card{Number: "6360111331111111", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "interpayment")
			So(card.Company.Name, ShouldEqual, "InterPayment")
		})

		Convey("Should work for InstaPayment", func() {
			card := Card{Number: "6370111331111111", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "instapayment")
			So(card.Company.Name, ShouldEqual, "InstaPayment")
		})

		Convey("Should work for JCB", func() {
			card := Card{Number: "353011133", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "jcb")
			So(card.Company.Name, ShouldEqual, "JCB")
		})

		Convey("Should work for Maestro", func() {
			card := Card{Number: "501855555", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "maestro")
			So(card.Company.Name, ShouldEqual, "Maestro")
		})

		Convey("Should work for Mastercard", func() {
			card := Card{Number: "55555555", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "mastercard")
			So(card.Company.Name, ShouldEqual, "Mastercard")
		})

		Convey("Should work for Visa Electron", func() {
			card := Card{Number: "4026424242", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "visa electron")
			So(card.Company.Name, ShouldEqual, "Visa Electron")
		})

		Convey("Should work for Visa", func() {
			card := Card{Number: "4242424242", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "visa")
			So(card.Company.Name, ShouldEqual, "Visa")
		})

		Convey("Should work for Cabal", func() {
			card := Card{Number: "604211212212", Cvv: "111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldBeNil)
			So(card.Company.Code, ShouldEqual, "cabal")
			So(card.Company.Name, ShouldEqual, "Cabal")
		})

		Convey("Should fail to recognize an unknown number", func() {
			card := Card{Number: "1112424242", Cvv: "1111", Month: month, Year: year}
			err := card.Brand()

			So(err, ShouldNotBeNil)
			So(card.Company.Code, ShouldEqual, "")
			So(card.Company.Name, ShouldEqual, "")
		})
	})
}
