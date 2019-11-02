Basic credit card validation using the [Luhn algorithm](http://en.wikipedia.org/wiki/Luhn_algorithm)

Currently identifies the following credit card companies:
* Aura
* American Express
* Bankcard
* Cabal
* China Unionpay
* Dankort
* Diners Club Carte Blanche
* Diners Club enRoute
* Diners Club International
* Discover
* Elo
* InterPayment
* InstaPayment
* JCB
* Hipercard
* Maestro
* MasterCard
* Visa
* Visa Electron

## Installation

```bash
go get github.com/sgumirov/go-cards-validation
```

## Attribution

This library based on original one : https://github.com/hubcash/cards . Changes include but not limited to:
* security fixes

See commit history for a full list.

## Quick Start

```go
// Initialize a new card:
card := creditcard.Card{Number: "4242424242424242", Cvv: "11111", Month: "02", Year: "2016"}

// Retrieve the card's brand (which credit card company this card belongs to)
err := card.Brand() // card.Brand({Code: "visa", Name: "Visa"})

// Display last four digits
lastFour, err := card.LastFour() // 4242

// Validate the card's number (without capturing)
err := card.Validate() // will return an error due to not allowing test cards

err := card.Validate(true) // this will work though
```
